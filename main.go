package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "pro10/src/app/tool/validator"
	"pro10/src/app/user"
)

func main() {
	r := gin.Default()

	//使用session-cookie
	store := cookie.NewStore([]byte("sescret"))
	////注入中间件
	r.Use(sessions.Sessions("mysession", store))

	user.NewRou(r)
	r.Run(":8010")
}
