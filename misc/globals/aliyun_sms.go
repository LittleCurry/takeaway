package globals

import (
	"io/ioutil"
	"encoding/base32"
	"encoding/base64"
	"time"
	"sort"
	"crypto/hmac"
	"crypto/sha1"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"net/http"
	"encoding/json"
)

type smsResponse struct {
	RawResponse []byte `json:"-"`
	RequestId   string `json:"RequestId"`
	Code        string `json:"Code"`
	Message     string `json:"Message"`
	BizId       string `json:"BizId"`
}

func (m *smsResponse) isSuccessful() bool {
	return m.Code == "OK"
}

type smsClient struct {
	Request    *aLiYunCommunicationRequest
	GatewayUrl string
	Client     *http.Client
}

func newClient(gatewayUrl string) *smsClient {
	smsClient := new(smsClient)
	smsClient.Request = &aLiYunCommunicationRequest{}
	smsClient.GatewayUrl = gatewayUrl
	smsClient.Client = &http.Client{}
	return smsClient
}

func (smsClient *smsClient) execute(accessKeyId, accessKeySecret, phoneNumbers, signName, templateCode, templateParam string) (*smsResponse, error) {
	err := smsClient.Request.setParamsValue(accessKeyId, phoneNumbers, signName, templateCode, templateParam)
	if err != nil {
		fmt.Println("err1:", err)
		return nil, err
	}
	endpoint, err := smsClient.Request.buildSmsRequestEndpoint(accessKeySecret, smsClient.GatewayUrl)
	if err != nil {
		fmt.Println("err2:", err)
		return nil, err
	}

	request, _ := http.NewRequest("GET", endpoint, nil)
	response, err := smsClient.Client.Do(request)
	if err != nil {
		fmt.Println("err3:", err)
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("err4:", err)
		return nil, err
	}
	defer response.Body.Close()

	result := new(smsResponse)
	err = json.Unmarshal(body, result)

	result.RawResponse = body
	return result, err
}

type aLiYunCommunicationRequest struct {
	//system parameters
	AccessKeyId      string
	Timestamp        string
	Format           string
	SignatureMethod  string
	SignatureVersion string
	SignatureNonce   string
	Signature        string

	//business parameters
	Action          string
	Version         string
	RegionId        string
	PhoneNumbers    string
	SignName        string
	TemplateCode    string
	TemplateParam   string
	SmsUpExtendCode string
	OutId           string
}

var encoding = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h897")

func (req *aLiYunCommunicationRequest) setParamsValue(accessKeyId, phoneNumbers, signName, templateCode, templateParam string) error {
	req.AccessKeyId = accessKeyId
	//location, _ := time.LoadLocation("GMT")
	//time.Now().In(location).Format(time.RFC1123)

	local, err := time.LoadLocation("GMT")
	if err != nil {
		fmt.Println("ALiYunCommunicationRequest_err:", err)
		local = time.Local
	}
	req.Timestamp = time.Now().In(local).Format("2006-01-02T15:04:05Z")
	//time.Now().In(local).Format(time.RFC1123)
	req.Format = "json"
	req.SignatureMethod = "HMAC-SHA1"
	req.SignatureVersion = "1.0"
	ra := GenerateSession()
	req.SignatureNonce = ra

	req.Action = "SendSms"
	req.Version = "2017-05-25"
	req.RegionId = "cn-hangzhou"
	req.PhoneNumbers = phoneNumbers
	req.SignName = signName
	req.TemplateCode = templateCode
	req.TemplateParam = templateParam
	req.SmsUpExtendCode = "90999"
	req.OutId = "abcdefg"
	return nil
}

func (req *aLiYunCommunicationRequest) smsParamsIsValid() error {
	if len(req.AccessKeyId) == 0 {
		return errors.New("AccessKeyId required")
	}

	if len(req.PhoneNumbers) == 0 {
		return errors.New("PhoneNumbers required")
	}

	if len(req.SignName) == 0 {
		return errors.New("SignName required")
	}

	if len(req.TemplateCode) == 0 {
		return errors.New("TemplateCode required")
	}

	if len(req.TemplateParam) == 0 {
		return errors.New("TemplateParam required")
	}

	return nil
}

func (req *aLiYunCommunicationRequest) buildSmsRequestEndpoint(accessKeySecret, gatewayUrl string) (string, error) {
	var err error
	if err = req.smsParamsIsValid(); err != nil {
		return "", err
	}
	// common params
	systemParams := make(map[string]string)
	systemParams["SignatureMethod"] = req.SignatureMethod
	systemParams["SignatureNonce"] = req.SignatureNonce
	systemParams["AccessKeyId"] = req.AccessKeyId
	systemParams["SignatureVersion"] = req.SignatureVersion
	systemParams["Timestamp"] = req.Timestamp
	systemParams["Format"] = req.Format

	// business params
	businessParams := make(map[string]string)
	businessParams["Action"] = req.Action
	businessParams["Version"] = req.Version
	businessParams["RegionId"] = req.RegionId
	businessParams["PhoneNumbers"] = req.PhoneNumbers
	businessParams["SignName"] = req.SignName
	businessParams["TemplateParam"] = req.TemplateParam
	businessParams["TemplateCode"] = req.TemplateCode
	businessParams["SmsUpExtendCode"] = req.SmsUpExtendCode
	businessParams["OutId"] = req.OutId
	// generate signature and sorted  query
	sortQueryString, signature := generateQueryStringAndSignature(businessParams, systemParams, accessKeySecret)
	return gatewayUrl + "?Signature=" + signature + sortQueryString, nil
}

func generateQueryStringAndSignature(businessParams map[string]string, systemParams map[string]string, accessKeySecret string) (string, string) {
	keys := make([]string, 0)
	allParams := make(map[string]string)
	for key, value := range businessParams {
		keys = append(keys, key)
		allParams[key] = value
	}

	for key, value := range systemParams {
		keys = append(keys, key)
		allParams[key] = value
	}

	sort.Strings(keys)

	sortQueryStringTmp := ""
	for _, key := range keys {
		rstkey := specialUrlEncode(key)
		rstval := specialUrlEncode(allParams[key])
		sortQueryStringTmp = sortQueryStringTmp + "&" + rstkey + "=" + rstval
	}

	sortQueryString := strings.Replace(sortQueryStringTmp, "&", "", 1)
	stringToSign := "GET" + "&" + specialUrlEncode("/") + "&" + specialUrlEncode(sortQueryString)

	sign := sign(accessKeySecret+"&", stringToSign)
	signature := specialUrlEncode(sign)
	return sortQueryStringTmp, signature
}

func specialUrlEncode(value string) string {
	rstValue := url.QueryEscape(value)
	rstValue = strings.Replace(rstValue, "+", "%20", -1)
	rstValue = strings.Replace(rstValue, "*", "%2A", -1)
	rstValue = strings.Replace(rstValue, "%7E", "~", -1)
	return rstValue
}

func sign(accessKeySecret, sortquerystring string) string {
	h := hmac.New(sha1.New, []byte(accessKeySecret))
	h.Write([]byte(sortquerystring))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func SmsCode(phone string, code string) error {

	var (
		gatewayUrl      = "http://dysmsapi.aliyuncs.com/"
		accessKeyId     = "LTAIiEQnQoRXF7CC"
		accessKeySecret = "twbBohqIebq6nz3PXLSxTwtCZbegbs"
		phoneNumbers    = phone
		signName        = "之炜物联"
		templateCode    = "SMS_148075206"
		templateParam   = `{"code":"` + code + `"}`
	)

	smsClient := newClient(gatewayUrl)
	result, err := smsClient.execute(accessKeyId, accessKeySecret, phoneNumbers, signName, templateCode, templateParam)
	if err != nil {
		fmt.Println("err:", err)
		return err
	}
	resultJson, err := json.Marshal(result)
	if err != nil {
		fmt.Println("err:", err)
		return err
	}
	if result.isSuccessful() {
		fmt.Println("A SMS is sent successfully:", resultJson)
		return nil
	} else {
		fmt.Println("Failed to send a SMS2:", string(resultJson))
		return errors.New(string(resultJson))

	}
	return nil

}

func SmsMsg(phone string, msg string, time string) error {

	var (
		gatewayUrl      = "http://dysmsapi.aliyuncs.com/"
		accessKeyId     = "LTAIiEQnQoRXF7CC"
		accessKeySecret = "twbBohqIebq6nz3PXLSxTwtCZbegbs"
		phoneNumbers    = phone
		signName        = "之炜物联"
		templateCode    = "SMS_151230335"
		templateParam   = `{"msg":"` + msg + `", "time":"` + time + `"}`
	)

	smsClient := newClient(gatewayUrl)
	result, err1 := smsClient.execute(accessKeyId, accessKeySecret, phoneNumbers, signName, templateCode, templateParam)
	if err1 != nil {
		fmt.Println("err1:", err1)
		return err1
	}
	resultJson, err2 := json.Marshal(result)
	if err2 != nil {
		fmt.Println("err2:", err2)
		return err2
	}
	if result.isSuccessful() {
		fmt.Println("A SMS is sent successfully:", resultJson)
		return nil
	} else {
		fmt.Println("Failed to send a SMS2:", string(resultJson))
		return errors.New(string(resultJson))
	}
}

func SmsMsgInfo(phone string, msg string, info string) error {

	var (
		gatewayUrl      = "http://dysmsapi.aliyuncs.com/"
		accessKeyId     = "LTAIiEQnQoRXF7CC"
		accessKeySecret = "twbBohqIebq6nz3PXLSxTwtCZbegbs"
		phoneNumbers    = phone
		signName        = "之炜物联"
		templateCode    = "SMS_158547266"
		//templateCode    = "SMS_159770489"
		templateParam   = `{"msg":"` + msg + `", "info":"` + info + `"}`
	)

	smsClient := newClient(gatewayUrl)
	result, err1 := smsClient.execute(accessKeyId, accessKeySecret, phoneNumbers, signName, templateCode, templateParam)
	if err1 != nil {
		fmt.Println("err1:", err1)
		return err1
	}
	resultJson, err2 := json.Marshal(result)
	if err2 != nil {
		fmt.Println("err2:", err2)
		return err2
	}
	if result.isSuccessful() {
		fmt.Println("A SMS is sent successfully:", resultJson)
		return nil
	} else {
		fmt.Println("Failed to send a SMS2:", string(resultJson))
		return errors.New(string(resultJson))
	}
}
