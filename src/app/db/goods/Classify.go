package goods

import (
	"errors"
	db2 "pro10/src/app/db"
)

type Classify struct {
	Id       int    `db:"id"`
	Name     string `db:"name" form:"name" binding:"required,min=1,max=10"`
	Supclass string `db:"supclass" form:"supclass" binding:"required" `
	Classnum int    `db:"classnum" form:"classnum" binding:"required"`
	Shopnum  int    `db:"shopnum"`
}

//添加商品分类
func (c *Classify) AddClassify() error {
	db := db2.DbPointer()
	dbst := "insert into goodclassify(name,supclass,classnum) values(?,?,?)"
	_, err := db.Exec(dbst, c.Name, c.Supclass, c.Classnum)
	return err
}

//删除商品分类
func (c *Classify) DeleteClassify() error {
	db := db2.DbPointer()
	dbst := "delete from goodclassify where id=?"
	result, err := db.Exec(dbst, c.Id)
	k, _ := result.RowsAffected()
	if k == 0 {
		err = errors.New("无该分类")
	}
	return err
}

//分页查找分类
func (c *Classify) CheckPage(page int, grade int) (error, []Classify) {
	db := db2.DbPointer()
	var class []Classify = make([]Classify, 5, 5)
	var i int = 0
	initrow := (page - 1) * 5
	dbst := "select * from goodclassify where classnum=? limit ?,?"
	rows, err := db.Query(dbst, grade, initrow, 5)
	if err != nil {
		return err, nil
	}
	for rows.Next() {
		err = rows.Scan(&class[i].Id, &class[i].Name, &class[i].Supclass, &class[i].Classnum, &class[i].Shopnum)
		if err != nil {
			return err, nil
		}
		i++
	}
	return nil, class
}
