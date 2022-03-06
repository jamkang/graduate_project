package user

import (
	"github.com/gin-gonic/gin"
)

//用户模块路由

type UserRou struct {
}

func NewUserRou(user *gin.RouterGroup) {
	user.POST("registerAdmi", AddAdministor)
	user.POST("/registerDis", AddDistributor)
	user.POST("/loginAdmi", LoginAdnimistor)
	user.POST("/loginDis", LoginDistributor)
}
