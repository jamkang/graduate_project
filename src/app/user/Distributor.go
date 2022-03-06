package user

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"pro10/src/app/db/user"
	"pro10/src/app/tool"
	"strconv"
)

func AddDistributor(c *gin.Context) {
	distributor := new(user.Distribution)
	if err := c.ShouldBind(&distributor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err:": err})
		return
	} else {
		if bo := tool.CityDeal(distributor.Area); !bo {
			c.JSON(http.StatusBadRequest, gin.H{"err:": "城市填写不合理"})
			return
		}
		if err := distributor.AddDistributor(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message:": "注册成功"})
}

func LoginDistributor(c *gin.Context) {
	var distributor *user.Distribution = new(user.Distribution)
	if err := c.ShouldBind(&distributor.User.UserBase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	if yanz := c.PostForm("yanz"); yanz != "123" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "验证码不匹配"})
		return
	}
	if err := distributor.LoginDistributor(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}

	//设置session
	session := sessions.Default(c)
	id := md5.Sum([]byte(strconv.Itoa(distributor.Id)))
	fmt.Println(distributor)
	session.Set(id, distributor)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
}
