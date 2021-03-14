package router

import (
	"bubble/controller"
	"bubble/setting"
	"github.com/gin-gonic/gin"
)

func SetupRouter() (r *gin.Engine) {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r = gin.Default()
	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", controller.Create)
		v1Group.GET("/todo", controller.List)
		v1Group.PUT("/todo/:id", controller.Update)
		v1Group.DELETE("/todo/:id", controller.Delete)
	}
	return
}
