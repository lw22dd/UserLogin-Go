package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 添加这一行
)

var DB *sql.DB

func DataBaseInit() {
	fmt.Println("开始链接数据库")

	//数据库
	//用户名:密码@tcp(ip:端口)/数据库的名字
	dsn := "root:123456@tcp(127.0.0.1:3306)/vue"
	//连接数据集
	db, err := sql.Open("mysql", dsn) //open不会检验用户名和密码
	if err != nil {
		fmt.Printf("dsn:%s invalid,err:%v\n", dsn, err)
		return
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		fmt.Printf("open %s faild,err:%v\n", dsn, err)
		return
	}
	fmt.Println("连接数据库成功~")
	DB = db

	// 查询所有表
	tables, err := db.Query("SHOW TABLES")
	if err != nil {
		fmt.Println("查询表失败:", err)
		return
	}
	defer tables.Close()

	fmt.Println("数据库中的表：")
	for tables.Next() {
		var tableName string
		tables.Scan(&tableName)
		fmt.Println("- ", tableName)
	}

}
