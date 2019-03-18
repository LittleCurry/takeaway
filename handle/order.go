package handle

import (
	"github.com/jinzhu/copier"
	"time"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"

	"github.com/LittleCurry/takeaway/vm"
	"github.com/LittleCurry/takeaway/misc/err_msg"
	"github.com/LittleCurry/takeaway/misc/model"
	"github.com/LittleCurry/takeaway/misc/driver"
)

func AddOrder(c *gin.Context) {

	createReq := vm.CreateOrderReq{}
	err1 := c.Bind(&createReq)
	fmt.Println("createReq:", createReq)
	if err1 != nil {
		fmt.Println("err1:", err1)
		c.JSON(http.StatusOK, err_msg.CodeMsg{1, "参数格式错误"})
		return
	}
	if len(createReq.UserId) == 0  {
		c.JSON(http.StatusOK, err_msg.CodeMsg{1, "user_id不能为空"})
		return
	}

	order := model.Order{}
	copier.Copy(&order, createReq)

	order.CreateTime = time.Now().Format("2006-01-02 15:04:05")

	_, err2 := driver.MySQL().Insert(&order)

	if err2 != nil {
		fmt.Println("err2:", err2)
		c.JSON(http.StatusOK, err_msg.CodeMsg{1, "创建用户失败"})
		return
	}

	c.JSON(http.StatusOK, err_msg.CodeMsg{0, "创建成功"})
	return

}

func OrderInfo(c *gin.Context) {
	order := model.Order{}
	orderId := c.Query("order_id")
	_, err := driver.MySQL().Where("order_id=?", orderId).Get(&order)
	if err != nil {
		fmt.Println("err:", err)
	}


	if len(order.OrderId) == 0 {
		c.JSON(http.StatusOK, err_msg.CodeMsg{1, "未找到该订单"})
		return
	}
	orderRes := vm.OrderRes{}
	copier.Copy(&orderRes, order)
	c.JSON(http.StatusOK, orderRes)
	return
}

func OrderList(c *gin.Context) {
	orderList := []model.Order{}
	driver.MySQL().Where("`del` = 0").Find(&orderList)
	ordersRes := make([]vm.OrderRes, 0)
	copier.Copy(&ordersRes, orderList)
	c.JSON(http.StatusOK, ordersRes)
	return
}
