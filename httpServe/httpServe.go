package httpServe

import (
	"demo/config"
	"demo/route"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// 启动HTTP服务,并打开主页
func HttpServerStart() {
	// 注册路由路径和处理函数
	route.Route()
	// 启动HTTP服务
	url := strings.Join([]string{config.Url, config.HttpPort}, ":") // 处理请求路径
	err := http.ListenAndServe(url, nil)
	if err != nil {
		log.Fatal(err)
	}

}

// 处理get请求
func GET(route string) {
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "get" {
			w.WriteHeader(405)
			return
		}

		fmt.Println(r.Form)
		fmt.Println()

	})
}

func POST() {

}

func PUT() {

}

func DELETE() {

}
