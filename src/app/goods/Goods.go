package goods

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pro10/src/app/db/goods"
	"pro10/src/app/tool"
	"strconv"
)

func AddGoods(c *gin.Context) {
	var good = new(goods.Goods)
	if err := c.ShouldBind(good); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	img, _ := c.FormFile("img")
	imgurl, err := tool.WriteImg(img, strconv.Itoa(int(good.Money)))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	good.Url = imgurl
	err = good.AddGood()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		fmt.Println("3", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": err, "message": "添加商品成功"})
}

func SeekGoods(c *gin.Context) {
	var g = new(goods.Goods)
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusFound, gin.H{"err": err, "message": errors.New("页码出现错误")})
		return
	}
	var goods = make([]goods.Goods, 5)
	goods, err = g.SeekGood(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err, "message": errors.New("数据出错")})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "查询成功", "data": goods})
}

func DeleteGood(c *gin.Context) {
	var g = new(goods.Goods)
	var err error
	g.Id, err = strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusFound, gin.H{"err": err, "message": errors.New("服务器未知问题")})
		fmt.Println("1", err)
		return
	}
	err = g.DeleteGood()
	if err != nil {
		c.JSON(http.StatusFound, gin.H{"err": err, "message": errors.New("服务器未知问题")})
		fmt.Println("2", err)
		return
	}
	c.JSON(http.StatusFound, gin.H{"err": err, "message": errors.New("商品删除成功")})
}
