package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	c := &Class{}
	c.Students = make(map[int]*Student, 50)
	for {
		fmt.Println("select the option number：")
		fmt.Print("1. add  2.list  3.remove  4.edit 5.exit\n")
		var do int8
		_, err := fmt.Scan(&do)
		if err != nil {
			fmt.Println("input error")
		}
		switch do {
		case 1:
			c.AddStudent()
		case 2:
			c.ShowStudentList()
		case 3:
			c.RemoveStudent()
		case 4:
			c.EditStudent()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("input error")
		}
	}
}

// Student student struct
type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Score int8   `json:"score"`
}

// Class class struct
type Class struct {
	Students map[int]*Student
}

// add student
func (c *Class) AddStudent() {
	var id int
	var name string
	var score int8
	fmt.Print("input id:")
	_, err := fmt.Scan(&id)
	fmt.Print("input name:")
	_, err = fmt.Scan(&name)
	fmt.Print("input score:")
	_, err = fmt.Scan(&score)
	if err != nil {
		fmt.Println("storing error")
	}
	_, isExist := c.Students[id]
	if isExist {
		fmt.Println("student existed")
		return
	}
	student := &Student{
		ID:    id,
		Name:  name,
		Score: score,
	}
	c.Students[id] = student
	fmt.Println("storing success")
}

// show student list
func (c *Class) ShowStudentList() {
	fmt.Printf("\t%s\t%s\t%s\n", "ID", "name", "score")
	sortIds := make([]int, 0, len(c.Students))
	for id := range c.Students {
		sortIds = append(sortIds, id)
	}
	sort.Ints(sortIds)
	for _, id := range sortIds {
		student := c.Students[id]
		fmt.Printf("\t%d\t%s\t%d\n", student.ID, student.Name, student.Score)
	}
}

// remove student from class
func (c *Class) RemoveStudent() {
	fmt.Println("input the id which you want to remove")
	var id int
	_, err := fmt.Scan(&id)
	if err != nil {
		fmt.Println("input error")
	}
	_, isExist := c.Students[id]
	if !isExist {
		fmt.Printf("%d not existed", id)
		return
	}
	delete(c.Students, id)
	fmt.Println("remove success")
}
func (c *Class) EditStudent() {
	fmt.Println("input the id which you want to edit")
	var id int
	_, err := fmt.Scan(&id)
	if err != nil {
		fmt.Println("input error")
	}
	_, isExist := c.Students[id]
	if !isExist {
		fmt.Printf("%d not existed", id)
		return
	}
	var name string
	var score int8
	fmt.Print("input name:")
	_, err = fmt.Scan(&name)
	fmt.Print("input score:")
	_, err = fmt.Scan(&score)
	if err != nil {
		fmt.Println("storing error")
	}
	student := &Student{
		ID:    id,
		Name:  name,
		Score: score,
	}
	c.Students[id] = student
	fmt.Println("edit success")
}
