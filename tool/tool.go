package tool

import (
	"fmt"
	"os/exec"
	"time"
)

func NetWorkStatus() bool {
	cmd := exec.Command("ping", "192.168.1.1")
	fmt.Println("网络测试开始:", time.Now().Unix())
	err := cmd.Run()
	fmt.Println("网络测试结束", time.Now().Unix())
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		fmt.Println("地址标志,OK")
	}
	return true
}
