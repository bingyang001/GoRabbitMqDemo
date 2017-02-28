package publishServer

import (
	//"goRabbitMq"
	"goRabbitMq/common"
)

//运行服务
func Run() {
	common.InitLog()
	//声明消息处理器
	//声明消息路由，注册路由
	//pubRouting := goRabbitMq.HttpRouting{}
	//pubRouting.RegisterPublishRouting(&goRabbitMq.MessageHandle{})
	//.....
	//	welcomRouting := goRabbitMq.HttpRouting{}
	//	welcomRouting.RegisterWelcomRouting()
	Start()
	//声明路由数组
	//	r := make([]goRabbitMq.HttpRouting, 2)
	//	r[0] = pubRouting
	//	r[1] = welcomRouting
	//	//注册所有路由
	//	registerRouting(r)
	common.Log.Info("runing....\r\n")
}

//停止
func Stop() {

}
