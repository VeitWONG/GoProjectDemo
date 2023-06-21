package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

func LoginController(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("template/layout.gohtml",
		"template/login.gohtml")
	if err != nil {
		fmt.Println("Error " + err.Error())
	}
	template.Execute(response, nil)
}
