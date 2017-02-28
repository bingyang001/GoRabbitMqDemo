package goRabbitMq

import (
	"net/http"

	"github.com/valyala/fasthttp"
)

//消息
type MessagePackage struct {
	AppId       string
	Code        string
	MsgUniqueId string
	Body        interface{}
}

//发布消息响应
type PublishResponse struct {
	Code int
	Msg  string
}

//发布消息处理
type MessageHandle struct {
	//Pub func(msgPackage *MessagePackage) PublishResponse
}

type HttpRouting struct {
	Path   string
	Handle func(w http.ResponseWriter, req *http.Request)
}

type PublishHandle struct {
	Handle func(ctx *fasthttp.RequestCtx)
}

//type Semaphore struct {
//	Max int
//	w   sync.WaitGroup{}
//	c   make(chan int, 1)
//}
