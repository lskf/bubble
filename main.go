package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
)

func main() {
	//创建数据库
	//sql:create database bubble
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close() //程序退出关闭数据库
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()
	r.Run(":9090")
}
