package model

import (
	"time"
)

type Animal struct {
	// 使用xorm的Column属性定义表中各字段的属性
	Id         int       `xorm:"pk autoincr notnull unique index comment('主键')"`
	AnimalName string    `xorm:"varchar(20)"`
	Created    time.Time `xorm:"created"`
	Updated    time.Time `xorm:"updated"`
}

type AnimalAction interface {
	Say() string
	Run()
	Howl()
}

func (a *Animal) Say() string {
	return "That was Animal"
}

func (a *Animal) Run() {

}

func (a *Animal) Howl() {

}
