package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"fmt"
)

var mysql *xorm.Engine

func init() {
	var err error
	mysql, err = xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/thpmax?charset=utf8")
	if err != nil {
		fmt.Println("数据库连接失败:", err)
	}

	InitArticleTotalCount()
}