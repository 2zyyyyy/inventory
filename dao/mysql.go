/*
@Description：gorm 操作 mysql
@Author : gilbert
@Date : 2022/5/2 7:04 PM
*/

package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 定义全局的 DB 对象
var DB *gorm.DB

func MysqlInit() (err error) {
	dsn := "root:7788@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	//fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//app.User, app.Password, app.Host, app.Port, app.Name)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("connect mysql failed, err: %v\n", err)
		return
	}
	fmt.Println("mysql 数据库连接成功！")
	return
}
