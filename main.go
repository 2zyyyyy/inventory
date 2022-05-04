package main

import (
	"inventory/dao"
	"inventory/models"
	"inventory/routers"
	"log"
)

/*
@Description：清单项目主入口
@Author : gilbert
@Date : 2022/5/2 7:01 PM
*/

func main() {
	err := dao.MysqlInit()
	if err != nil {
		panic(err)
	}
	// 自动同步
	err = dao.DB.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Println("自动迁移失败，err：", err)
		return
	}
	// 注册路由
	r := routers.RegisterRouters()
	err = r.Run(":8888")
	if err != nil {
		log.Println("router run failed, err:", err)
		return
	}
}
