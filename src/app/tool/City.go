package tool

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

//地区
type Area struct {
	Grade    int    `form:"grade" json:"grade" db:"grade" binding:"required,gte=0,lte=2" `
	Province string `form:"province" json:"province" db:"province"`
	City     string `form:"city" json:"city" db:"city"`
	Counties string `form:"counties" json:"counties" db:"counties"`
}

//城市处理
func CityDeal(area Area) bool {
	city, bo := Province(area.Province)
	if bo {
		for i := 1; i <= area.Grade; i++ {
			switch i {
			case 1:
				city, bo = City(city, area.City)
			case 2:
				city, bo = City(city, area.Counties)
			default:
				bo = false
			}
		}
	}
	if bo {
		return true
	}
	return false
}

//城市处理--省
func Province(pro string) (string, bool) {
	prostring := "./src/app/tool/city/0.json"
	return checkCity(prostring, pro)
}

//城市查找，市和下一级
func City(pro string, city string) (string, bool) {
	strpro := "./src/app/tool/city/" + pro + ".json"
	return checkCity(strpro, city)
}

func checkCity(url string, city string) (string, bool) {
	filePte, err := os.Open(url)
	if err != nil {
		fmt.Println("打开文件失败", err)
	}
	defer filePte.Close()
	var data []map[string]interface{}
	decoder := json.NewDecoder(filePte)
	decoder.Decode(&data)
	for _, value := range data {
		if value["cn_name"].(string) == city {
			//fmt.Println(value["id"], value["cn_name"])
			return strconv.FormatFloat(value["id"].(float64), 'f', 0, 64), true
		}
	}
	return "", false
}
