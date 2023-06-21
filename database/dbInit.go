package database

import "fmt"

func DBInit() {
	engine, _ := Conn()
	defer engine.Close()

	_, err := engine.Exec("CREATE DATABASE go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully")
	}
}
