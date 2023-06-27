package model

import (
	"fmt"
	"time"
)

type Cat struct {
	*Animal
}

type Dog struct {
	*Animal
}

func (c *Cat) Say() string {
	return "That was Cat"
}

func (d *Dog) Say() string {
	return "That was Dog"
}

func Test1() {

	c1 := Cat{
		Animal: &Animal{
			Id:         1,
			AnimalName: "cat1",
			Created:    time.Now(),
			Updated:    time.Now(),
		},
	}
	// 当Cat结构体没有实现Say()方法时,下列两种方法调用的都是被组合结构体Animal的Eat()方法
	fmt.Println(c1.Animal.Say())
	fmt.Println(c1.Say())

	d1 := Dog{
		Animal: &Animal{},
	}
	// 当Dog结构体也实现了Say()方法的时候
	fmt.Println(d1.Say())        // 调用Dog实现的方法
	fmt.Println(d1.Animal.Say()) // 调用被组合结构体Animal的方法
}
