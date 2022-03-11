package goods

import "github.com/gin-gonic/gin"

func NewGoodsRouter(goods *gin.RouterGroup) {
	goods.POST("/AddClassify", AddClassify)
	goods.GET("/DeleteClassify", DeleteClassify)
}
