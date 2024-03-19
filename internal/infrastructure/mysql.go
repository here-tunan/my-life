package infrastructure

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-my-life/env"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var Mysql *xorm.Engine

func init() {
	dataSourName := fmt.Sprintf("%s:%s@tcp(%s:%d)/life?charset=utf8mb4",
		env.Prop.Mysql.Username, env.Prop.Mysql.Password, env.Prop.Mysql.Host, env.Prop.Mysql.Port)

	newEngine, _ := xorm.NewEngine("mysql", dataSourName)

	// 设置驼峰转换
	newEngine.SetMapper(names.GonicMapper{})

	Mysql = newEngine
}
