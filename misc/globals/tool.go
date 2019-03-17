package globals

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net"
	"strings"
)

func GetLimitAndStart(c *gin.Context) (int, int) {
	start, err := strconv.Atoi(c.Query("start"))
	if err != nil {
		start = 0
	}
	if start < 0 {
		start = 0
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 20
	}
	if limit < 0 || limit > 1000 {
		limit = 20
	}
	return start, limit
}

func GetMacAdress() string {

	// 获取本机的MAC地址
	interfaces, _ := net.Interfaces()
	macAddress := ""
	for _, inter := range interfaces {
		if len(inter.HardwareAddr) > 0 {
			macAddress = inter.HardwareAddr.String()
		}
	}
	return strings.Replace(macAddress, ":", "", -1)

}
