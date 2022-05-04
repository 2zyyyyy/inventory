/*
@Description：路由注册
@Author : gilbert
@Date : 2022/5/4 8:31 PM
*/

package routers

import (
	"github.com/gin-gonic/gin"
	"inventory/controller"
)

func RegisterRouters() *gin.Engine {
	r := gin.Default()
	// 设置 gin 模板文件引用的静态文件路径
	r.Static("/static", "static")
	// 设置 gin 的模板文件路径
	r.LoadHTMLGlob("templates/*")

	// 访问待办清单首页
	r.GET("/", controller.IndexHandler)

	// v1 待办清单
	v1Group := r.Group("v1")
	{
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改
		v1Group.PUT("todo/:id", controller.UpdateTodo)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DelTodo)
	}
	return r
}
