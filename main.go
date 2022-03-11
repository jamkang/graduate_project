package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"pro10/src/app/goods"
	_ "pro10/src/app/tool/validator"
	"pro10/src/app/user"
)

func main() {
	r := gin.Default()

	//使用session-cookie
	store := cookie.NewStore([]byte("sescret"))
	////注入中间件
	r.Use(sessions.Sessions("mysession", store))

	//路由分组
	user.NewUserRou(r.Group("/user"))
	goods.NewGoodsRouter(r.Group("/good"))

	r.Run(":8010")
}
