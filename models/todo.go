/*
@Description：todo 实体类
@Author : gilbert
@Date : 2022/5/2 7:04 PM
*/

package models

import (
	"fmt"
	"inventory/dao"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// CreateTodo 创建 todo
func CreateTodo(todo *Todo) (err error) {
	if err = dao.DB.Create(&todo).Error; err != nil {
		return err
	}
	return
}

// GetTodoList 获取 todo 列表
func GetTodoList() (todoList []*Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

// GetTodoById 通过 id 获取 todo 数据
func GetTodoById(id string) (todo *Todo, err error) {
	if err = dao.DB.Find(&todo).Error; err != nil {
		return nil, err
	}
	return
}

// UpdateTodo 更新 todo 数据
func UpdateTodo(todo *Todo) (err error) {
	if err = dao.DB.Save(todo).Error; err != nil {
		fmt.Printf("dao.db.save failed, err:%v\n", err)
		return
	}
	return
}

// DelTodoById 根据 id 删除 todo 数据
func DelTodoById(id string) (err error) {
	if err = dao.DB.Where("id = ?", id).Delete(&Todo{}).Error; err != nil {
		fmt.Printf("dao.db.delete failed, err:%v\n", err)
		return
	}
	return
}
