package user

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"pro10/src/app/db/user"
	"strconv"
)

//用户注册
func AddAdministor(c *gin.Context) {
	var adminstor user.Administor
	if err := c.ShouldBind(&adminstor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "数据传输问题"})
		return
	}
	repasswd := c.PostForm("repasswd")
	yzcode := c.PostForm("yzcode")
	if repasswd != adminstor.Passwd {
		c.JSON(http.StatusBadRequest, gin.H{"err:": "密码不一致"})
	} else if yzcode != "1234" {
		c.JSON(http.StatusBadRequest, gin.H{"err:": "验证码错误"})
	} else {
		if err02 := adminstor.AddAdministor(); err02 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err:": err02})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

//管理员登录
func LoginAdnimistor(c *gin.Context) {
	var adminstor *user.Administor = new(user.Administor)
	if err := c.ShouldBind(&adminstor.User.UserBase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	if yanz := c.PostForm("yanz"); yanz != "123" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "验证码不匹配"})
		return
	}
	if err := adminstor.LoginAdministor(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	//设置session
	session := sessions.Default(c)
	id := md5.Sum([]byte(strconv.Itoa(adminstor.Id)))
	session.Set(id, adminstor)
	session.Save()
	fmt.Println(session.Get(id))
	c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
}
