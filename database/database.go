package database

import (
	"demo/config"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func Conn() (*xorm.Engine, error) {
	path := strings.Join([]string{config.UserName, ":", config.Password, "@", config.Protocol, "(", config.IpAddress, ":", config.Port, ")/", config.DBName, "?charset=utf8mb4"}, "")

	engine, err := xorm.NewEngine("mysql", path)
	if err != nil {
		return nil, err
	}

	/*
			先使用IsTableExist判断g表是否存在,如果表存在输入已存在,
			不存在则使用CreateTables创建表


		// 判断在数据库中是否有结构体对应的数据表
		isUserTableExist, _ := engine.IsTableExist(&model.User{})
		isAnimalTabeExist, _ := engine.IsTableExist(&model.Animal{})
		if isUserTableExist && isAnimalTabeExist {
			// 表已存在时输入相关信息
			fmt.Println("user表和animal表已经存在")
		} else {
			// 不存在时创建表
			err := engine.CreateTables(&model.User{}, &model.Animal{})
			if err != nil {
				return nil, err
			}
			fmt.Println("数据表创建成功")
		}
	*/

	/*
		// 显示生成的SQL语句
		engine.ShowSQL(true)
		// 设置连接池的空闲数大小
		engine.SetMaxIdleConns(30)
		// 设置最大打开连接数
		engine.SetMaxOpenConns(30)

		// 测试连通性
		if err = engine.Ping(); err != nil {
			fmt.Println("数据库连接失败")
		}
		err = engine.Sync(new(model.User))
		if err != nil {
			fmt.Println("表结构同步失败")
		}
	*/
	return engine, nil
}

// 插入数据
func Insert() {

}
