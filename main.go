package main

import (
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/LittleCurry/takeaway/misc/driver"
	"github.com/LittleCurry/takeaway/misc/globals"
	"github.com/LittleCurry/takeaway/handle"
)

func init() {
	//driver.RedisInit("127.0.0.1:6379", 0)
	//driver.OrmInit("root:1qaz!QAZ@tcp(localhost:3306)/takeaway?charset=utf8")
	/* 连接redis */
	driver.RedisInit("47.92.69.207:6379", 0)
	/* 连接mysql */
	driver.OrmInit("root:1qaz!QAZ@tcp(47.92.69.207:3306)/takeaway?charset=utf8")
}

func main() {

	fmt.Println("takeaway :", time.Now().Format("2006-01-02 15:04:05"))

	driver.RedisClient.Set("abc", "ceshiceshi啊", 7*24*time.Hour)
	res, err1 := driver.RedisClient.Get("abc").Result()

	fmt.Println("res:", res)
	fmt.Println("err1:", err1)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default() //获得路由实例
	router.Use(globals.Cors())

	//router.GET("/test", handle.Test)

	user := router.Group("/user")
	{
		user.GET("", handle.UserInfo)
		user.POST("/create", handle.AddUser)
		user.GET("/list", handle.UserList)
	}

	order := router.Group("/order")
	{
		order.GET("", handle.OrderInfo)
		order.POST("/create", handle.AddOrder)
		order.GET("/list", handle.OrderList)
	}

	router.Run(":7000")
}