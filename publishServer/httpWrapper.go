// httpWrapper
package publishServer

import (
	"goRabbitMq"
	"goRabbitMq/common"
	"net/http"
)

//rabbit mq http server 入口
func registerRouting(routing []goRabbitMq.HttpRouting) {
	common.Log.Infof("register routing count %d", len(routing))
	for _, r := range routing {
		http.HandleFunc(r.Path, r.Handle)
		common.Log.Infof("register path %s", r.Path)
	}
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		common.Log.Fatal("ListenAndServe: ", err)
	}
	common.Log.Info("register routing end...")
}

func close() {

}
