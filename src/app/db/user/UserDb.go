package user

import (
	"errors"
	"fmt"
	"pro10/src/app/db"
	"pro10/src/app/tool"
	time2 "time"
)

//登录信息
type UserBase struct {
	Name   string `form:"name" json:"name" db:"name" binding:"required,min=2,max=15"`
	Passwd string `form:"passwd" json:"passwd" db:"passwd" binding:"required,min=5,max=15"`
	Blong  int    `form:"blong" json:"blong" db:"blong" binding:"gte=0,lte=1"`
}

//用户
type User struct {
	UserBase
	Id   int    `db:"id""`
	Sex  int    `form:"sex" json:"sex" db:"sex" binding:"required,gte=0,lte=1"`
	Tele string `form:"tele" json:"tele" db:"tele" binding:"checkTele"`
	Time int64  `db:"time"`
}

//添加用户
func (u *User) AddUser() (int, error) {
	//处理密码
	u.Passwd = tool.Encryption(u.Passwd)
	//获取时间
	u.Time = time2.Now().Unix()
	Db := db.DbPointer()
	//输入到数据库
	dbstring01 := "insert into users(name,passwd,sex, belong, tele,time) values(?,?,?,?,?,?)"
	if _, err := Db.Exec(dbstring01, u.Name, u.Passwd, u.Sex, u.Blong, u.Tele, u.Time); err != nil {
		fmt.Println("第一次", u.Passwd)
		return -1, err
	}
	dbstring02 := "select id from users where name=?"
	uidRow := Db.QueryRow(dbstring02, u.Name)
	var uid int
	uidRow.Scan(&uid)
	fmt.Println("uid:", uid)
	return uid, nil
}

//用户登录
func (u *User) LoginUser() (User, error) {
	var user User
	db := db.DbPointer()
	str := "SELECT * from users where name=?"
	useDa := db.QueryRow(str, u.Name)
	if err := useDa.Scan(&user.Id, &user.Name, &user.Passwd, &user.Blong,
		&user.Sex, &user.Time, &user.Tele); err != nil {
		fmt.Println(u.Name)
		return user, err
	}
	if tool.CompasionCode(user.Passwd, u.Passwd) {
		err := errors.New("用户密码错误")
		return user, err
	}
	return user, nil
}
