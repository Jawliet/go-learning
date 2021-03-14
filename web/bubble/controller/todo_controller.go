package controller

import (
	"bubble/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Create(c *gin.Context) {
	var todo model.Todo
	_ = c.BindJSON(&todo)

	if err := model.Create(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
func List(ctx *gin.Context) {
	if list, err := model.List(); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, list)
	}
}
func Update(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	todo, err := model.One(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	_ = ctx.BindJSON(&todo)
	if err = model.Update(todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

func Delete(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	intId, _ := strconv.Atoi(id)
	if err := model.Delete(intId); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
