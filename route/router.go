package route

import (
	"demo/controller"
	"net/http"
)

// 路由和控制器注册
func Route() {
	http.HandleFunc("/", controller.HomeController)
	http.HandleFunc("/login", controller.LoginController)
}
