// fastHttpServer
package publishServer

import (
	"fmt"
	"goRabbitMq"
	"goRabbitMq/common"

	"github.com/valyala/fasthttp"
)

func Start() {
	//handle routing
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/go/publish":
			//h := &goRabbitMq.PublishHandle{}
			//h.Handle(ctx)
			goRabbitMq.Handle2(ctx)
		case "/":
			FastHttpRegisterWelcomHandle(ctx)
		default:
			FastHttpRegister404Handle(ctx)
		}
	}
	common.Log.Debug("begin listen ....\r\n")
	if err := fasthttp.ListenAndServe(":8080", requestHandler); err != nil {
		common.Log.Fatalf("Error in ListenAndServe: %s", err)
	}
	//select {}
}

//default handle
func FastHttpRegisterWelcomHandle(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "welcom to go rabbitmq service v1.0!\n\n")
	ctx.SetContentType("text/plain; charset=utf8")
}

//404 handle
func FastHttpRegister404Handle(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "not find!\n\n")
	ctx.SetContentType("text/plain; charset=utf8")
	ctx.SetStatusCode(404)
}
