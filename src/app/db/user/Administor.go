package user

import (
	"fmt"
	"pro10/src/app/db"
)

//管理员
type Administor struct {
	User
	Power int `form:"power" json:"tele" db:"power" binding:"gte=0,lte=100"`
}

//添加管理员
func (a *Administor) AddAdministor() error {
	uid, err := a.User.AddUser()
	if err != nil {
		return err
	}
	dbstring03 := "insert into administor(uid,power) values(?,?)"
	if _, err := db.Db.Exec(dbstring03, uid, a.Power); err != nil {
		db.Db.Exec("delete from users where id=?", uid)
		return err
	}
	return nil
}

//登录
func (a *Administor) LoginAdministor() error {
	var administor *Administor = new(Administor)
	var err error
	administor.User, err = a.User.LoginUser()
	if err != nil {
		fmt.Println("2", err)
		return err
	}
	str2 := "SELECT `power` FROM administor WHERE uid=?"
	adm := db.Db.QueryRow(str2, administor.User.Id)
	err = adm.Scan(&administor.Power)
	if err != nil {
		fmt.Println("3", err)
		return err
	}
	*a = *administor
	return nil
}
