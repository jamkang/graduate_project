package goods

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro10/src/app/db/goods"
	"strconv"
)

//添加商品分类
func AddClassify(c *gin.Context) {
	var class *goods.Classify = new(goods.Classify)
	if err := c.ShouldBind(class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	if err := class.AddClassify(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "添加商品分类成功"})
}

//删除商品分类
func DeleteClassify(c *gin.Context) {
	var class *goods.Classify = new(goods.Classify)
	class.Id, _ = strconv.Atoi(c.Query("id"))
	if err := class.DeleteClassify(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "商品分类删除成功"})
}
