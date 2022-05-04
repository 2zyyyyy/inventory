/*
@Description：控制器
@Author : gilbert
@Date : 2022/5/2 11:12 PM
*/

package controller

import (
	"github.com/gin-gonic/gin"
	"inventory/models"
	"log"
	"net/http"
)

/** 访问url -> controller -> service -> dao -> models 【请求来了 -> 控制器 -> 业务逻辑 -> 模型层的增删改查】
 * @Description 访问首页
 * @Param
 * @return
 **/

// IndexHandler 首页
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// CreateTodo 创建 todo
func CreateTodo(c *gin.Context) {
	// 从请求中取数据
	var todo models.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		log.Println("bind json failed, err:", err)
		return
	}
	// 落库
	_ = models.CreateTodo(&todo)
	// 响应
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "保存数据失败",
			"error":   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// GetTodoList 查看所有待办清单
func GetTodoList(c *gin.Context) {
	// 查询所有数据
	todoList, err := models.GetTodoList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "查询数据失败",
			"error":   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

// UpdateTodo 更新清单数据
func UpdateTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "get id failed.",
		})
	} else {
		todo, err := models.GetTodoById(id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "查询数据失败",
				"error":   err.Error(),
			})
		} else {
			// 查询到了则更新数据
			_ = c.BindJSON(&todo)
			err = models.UpdateTodo(todo)
			if err != nil {
				// 更新失败
				c.JSON(http.StatusOK, gin.H{
					"message": "更新数据失败",
					"err":     err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		}
	}
}

// DelTodo 删除清单数据
func DelTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"message": "当前 id 查询数据失败",
			"err":     "error",
		})
	} else {
		todo := models.Todo{}
		// 删除数据
		err := models.DelTodoById(id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "当前 id 删除数据失败",
				"err":     err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, todo)
		}
	}
}
