package model

import "time"

type User struct {
	// 使用xorm的Column属性定义表中各字段的属性
	Id       int    `xorm:"pk autoincr notnull unique index comment('主键')"`
	UserName string `xorm:"varchar(20)"`
	Age      int
	Gender   int       `xorm:"default(0)"`
	Email    string    `xorm:"varchar(50)"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}

/*
可以为结构体编写get和set方法但没必要
go设计了使用大小写来控制公有/私有的方式
func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetAge(age int) {
	u.age = age
}

func (u *User) SetGender(gender int) {
	u.gender = gender
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u User) Name() striTime
	return u.name
}

func (u User) Age() int {
	return u.age
}

func (u User) Gender() int {
	return u.gender
}

func (u User) Email() string {
	return u.email
}
*/
