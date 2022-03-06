package Db

import (
	"fmt"
	"pro10/src/app/tool"
)

//分销商
type Distribution struct {
	User
	tool.Area
	money int `db:"money"`
}

//添加分销商
func (d *Distribution) AddDistributor() error {
	uid, err := d.User.AddUser()
	if err != nil {
		return err
	}
	db := DbPointer()
	dbstr01 := "insert into distributor(uid,grade,province,city,counties,money) values(?,?,?,?,?,?)"
	if _, err := db.Exec(dbstr01, uid, d.Grade, d.Province, d.City, d.Counties, d.money); err != nil {
		Db.Exec("delete from users where id=?", uid)
		return err
	}
	return nil
}

//登录分销商
func (d *Distribution) LoginDistributor() error {
	var distribution *Distribution = new(Distribution)
	var err error
	distribution.User, err = d.User.LoginUser()
	str2 := "SELECT * FROM distributor WHERE uid=?"
	adm := Db.QueryRow(str2, distribution.User.Id)
	err = adm.Scan(&distribution.Id, &distribution.Grade,
		&distribution.Province, &distribution.City,
		&distribution.Counties, &distribution.money)
	if err != nil {
		fmt.Println(err)
		return err
	}

	*d = *distribution

	return nil
}
