package model

import "bubble/dao"

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func Create(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func List() (list []*Todo, err error) {
	if err = dao.DB.Find(&list).Error; err != nil {
		return nil, err
	}
	return
}
func One(id string) (todo *Todo, err error) {
	if err = dao.DB.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return
}
func Update(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}
func Delete(id int) (err error) {
	err = dao.DB.Delete(&Todo{ID: id}).Error
	return
}
