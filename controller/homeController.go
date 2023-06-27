package controller

import (
	"demo/model"
	"fmt"
	"html/template"
	"net/http"
)

func HomeController(response http.ResponseWriter, request *http.Request) {
	teacher := model.Teacher{
		Name:    "测试",
		Subject: "测试",
	}
	students := []model.Student{
		{Id: 1001, Name: "Peter", Country: "China"},
		{Id: 1002, Name: "Jeniffer", Country: "Sweden"},
	}
	class := model.Class{
		Teacher:  teacher,
		Students: students,
	}

	template, err := template.ParseFiles("template/layout.gohtml",
		"template/nav.gohtml",
		"template/content.gohtml")

	if err != nil {
		fmt.Println("Error " + err.Error())
	}
	template.Execute(response, class)
}
