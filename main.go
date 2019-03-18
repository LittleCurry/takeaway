package main

import (
	"time"
	"fmt"
	"github.com/gin-gonic/gin"

	"gitlab.com/SiivaVideoStudio/cloud_server/takeaway/handle"
	"gitlab.com/SiivaVideoStudio/cloud_server/takeaway/misc/driver"
	"gitlab.com/SiivaVideoStudio/cloud_server/takeaway/misc/globals"
)

func init() {
	driver.RedisInit("127.0.0.1:6379", 0)
	driver.OrmInit("localhost:27017/takeaway")
}

func main() {

	fmt.Println("takeaway :", time.Now().Format("2006-01-02 15:04:05"))

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default() //获得路由实例
	router.Use(globals.Cors())

	user := router.Group("/user")
	{
		user.GET("", handle.UserInfo)
		user.GET("/create", handle.AddUser)
		user.GET("/list", handle.UserList)
	}

	order := router.Group("/order")
	{
		order.GET("", handle.OrderInfo)
		order.GET("/create", handle.AddOrder)
		order.GET("/list", handle.OrderList)
	}

	router.Run(":7000")
}