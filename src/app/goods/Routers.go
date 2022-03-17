package goods

import "github.com/gin-gonic/gin"

func NewGoodsRouter(goods *gin.RouterGroup) {
	goods.POST("/AddClassify", AddClassify)
	goods.GET("/DeleteClassify", DeleteClassify)
	goods.POST("/AddGood", AddGoods)
	goods.GET("/SeekGoods", SeekGoods)
	goods.GET("/DeleteGood", DeleteGood)
}
