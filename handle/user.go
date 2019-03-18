package handle

import (
	"time"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"strconv"
	"github.com/jinzhu/copier"
	"github.com/LittleCurry/takeaway/vm"
	"github.com/LittleCurry/takeaway/misc/err_msg"
	"github.com/LittleCurry/takeaway/misc/model"
	"github.com/LittleCurry/takeaway/misc/driver"
	"github.com/LittleCurry/takeaway/misc/globals"
)

func AddUser(c *gin.Context) {

	createReq := vm.CreateUserReq{}
	err1 := c.Bind(&createReq)
	fmt.Println("createReq:", createReq)
	if err1 != nil {
		fmt.Println("err1:", err1)
		c.JSON(http.StatusOK, err_msg.CodeMsg{1, "参数格式错误"})
		return
	}
	if len(createReq.Phone) == 0 || len(createReq.Passwd) == 0 {
		c.JSON(http.StatusOK, err_msg.CodeMsg{1, "账号或密码不能为空"})
		return
	}
	onlyUser := model.User{}
	driver.MySQL().Where("phone=?", createReq.Phone).Get(&onlyUser)
	if len(onlyUser.UserId) > 0 {
		c.JSON(http.StatusOK, err_msg.CodeMsg{1, createReq.Phone + "账号已存在"})
		return
	}

	user := model.User{}
	copier.Copy(&user, createReq)
	user.Passwd = globals.MakeMd5FromString(createReq.Passwd)

	user.UserId = strconv.Itoa(int(time.Now().Unix()))
	user.CreateTime = time.Now().Format("2006-01-02 15:04:05")

	_, err2 := driver.MySQL().Insert(&user)

	if err2 != nil {
		fmt.Println("err2:", err2)
		c.JSON(http.StatusOK, err_msg.CodeMsg{1, "创建用户失败"})
		return
	}

	c.JSON(http.StatusOK, err_msg.CodeMsg{0, "创建成功"})
	return

}

func UserInfo(c *gin.Context) {
	user := model.User{}
	userId := c.Query("user_id")
	_, err := driver.MySQL().Where("user_id=?", userId).Get(&user)
	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println("user:", user)

	if len(user.UserId) == 0 {
		c.JSON(http.StatusOK, err_msg.CodeMsg{1, "未找到该用户"})
		return
	}
	usersRes := vm.UserRes{}
	copier.Copy(&usersRes, user)
	c.JSON(http.StatusOK, usersRes)
	return
}

func UserList(c *gin.Context) {
	userList := []model.User{}
	driver.MySQL().Where("`del` = 0").Find(&userList)
	usersRes := make([]vm.UserRes, 0)
	copier.Copy(&usersRes, userList)
	c.JSON(http.StatusOK, usersRes)
	return
}
