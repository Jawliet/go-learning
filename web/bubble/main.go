package main

import (
	"bubble/dao"
	"bubble/model"
	"bubble/router"
	"bubble/setting"
	"fmt"
)

func main() {
	//if len(os.Args)<2 {
	//	fmt.Println("Usage: ./bubble/config/config.ini")
	//	return
	//}
	//if err:=setting.Init(os.Args[1]);err!=nil {
	//	fmt.Printf("locad config from file failed, err:%v\n",err)
	//	return
	//}
	_ = setting.Init("/Users/Jawliet/GoLandProjects/go-learning/web/bubble/conf/config.ini")
	if err := dao.InitMySql(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	_ = dao.DB.AutoMigrate(&model.Todo{})
	r := router.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
