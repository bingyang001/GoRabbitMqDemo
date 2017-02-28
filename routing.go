// GoRabbitMq project GoRabbitMq.go
package goRabbitMq

import (
	"goRabbitMq/common"
	"io/ioutil"
	"net/http"
	//"time"

	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"
)

//注册路由
func (rout *HttpRouting) RegisterPublishRouting(handle *MessageHandle) {
	rout.Path = "/go/publish"
	rout.Handle = func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		common.Log.Info("receive request , path /go/publish")
		by, er := ioutil.ReadAll(req.Body)
		if er != nil {
			common.Log.Errorf("read body error %s", er)
			return
		}
		o := &MessagePackage{}
		err := ffjson.Unmarshal(by, o)
		if err != nil {
			common.Log.Errorf("Deserialization body error %s", err)
			return
		}
		response := handle.Publish(o)
		responseBy, _ := ffjson.Marshal(response)
		w.Write(responseBy)
		common.Log.Infof("handle end...code=%d,msg=%s", response.Code, response.Msg)
	}
}

func (rout *HttpRouting) RegisterWelcomRouting() {
	rout.Path = "/"
	rout.Handle = func(w http.ResponseWriter, req *http.Request) {
		msg := []byte("welcom to go rabbitmq service v0.1.")
		w.Write(msg)
	}
}

func Handle2(ctx *fasthttp.RequestCtx) {
	//startTime := time.Now()
	by := ctx.Request.Body()
	//common.Log.Info("request body length")
	//构建message 包
	o := &MessagePackage{}
	//now := time.Now()
	err := ffjson.Unmarshal(by, o)
	if err != nil {
		common.Log.Errorf("Deserialization body error %s", err)
		return
	}
	//runTime := time.Now().Sub(now).Seconds() * 1000
	//common.Log.Infof("ffjson.Unmarshal request body ,run time %s", ToString("run ok ,milliseconds  ", runTime))
	h := &MessageHandle{}
	response := h.Publish(o)
	responseBy, _ := ffjson.Marshal(response)
	ctx.Write(responseBy)
	ctx.SetContentType("apllication/json")
	//totalTime := time.Now().Sub(startTime).Seconds() * 1000
	//common.Log.Infof("handle end...code=%d,msg=%s,total time=%s", response.Code, response.Msg, ToString("run ok ,milliseconds  ", totalTime))
	//common.Log.Infof("handle end...code=%d,msg=%s", response.Code, response.Msg)
}
