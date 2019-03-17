package vm

type CreateUserReq struct {
	Phone  string `json:"phone"`
	Passwd string `json:"passwd"`
}

type UserRes struct {
	Phone      string `json:"phone"`
	UserId     string `json:"user_id"`
	Passwd     string `json:"passwd"`
	NickName   string `json:"nick_name"`
	Gender     int    `json:"gender"`
	WxOpenid   string `json:"wx_openid"`
	WxUnionid  string `json:"wx_unionid"`
	Address    string `json:"address"`
	Del        int    `json:"del"`
	CreateTime string `json:"create_time"`
}
