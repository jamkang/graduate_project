package user

import (
	"github.com/gin-gonic/gin"
)

//用户模块路由

type UserRou struct {
}

func NewRou(r *gin.Engine) {
	r.POST("/user/registerAdmi", AddAdministor)
	r.POST("/user/registerDis", AddDistributor)
	r.POST("/user/loginAdmi", LoginAdnimistor)
	r.POST("/user/loginDis", LoginDistributor)
}
