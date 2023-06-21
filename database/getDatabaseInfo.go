package database

import (
	"fmt"

	"xorm.io/xorm"
)

func GetDatabaseInfo(engine *xorm.Engine) {
	// 读取数据库表信息
	result, _ := engine.DBMetas()
	for i := 0; i < len(result); i++ {
		fmt.Printf("数据表名: %s \n", result[i].Name)
		for j := 0; j < len(result[i].Columns()); j++ {
			fmt.Printf("第%d列: %s \n", j, result[i].Columns()[j].Name)
		}
		fmt.Print("\n")
	}
}
