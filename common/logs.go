package common

import (
	"fmt"
	"os"

	"github.com/op/go-logging"
)

var format = logging.MustStringFormatter(
	`%{time:2006-01-02T15:04:05} %{level:.1s} %{id:04d} %{module} %{message}`,
)
var Log = logging.MustGetLogger("gorabbitmq")

func InitLog() {
	fs, err := os.OpenFile("log/rabbitlog.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 077)
	if err == nil {
		backend2 := logging.NewLogBackend(fs, "", 0)
		backend2Formatter := logging.NewBackendFormatter(backend2, format)
		backend1Leveled := logging.AddModuleLevel(backend2)
		backend1Leveled.SetLevel(logging.ERROR, "")
		backend1Leveled.SetLevel(logging.INFO, "")
		backend1Leveled.SetLevel(logging.DEBUG, "")
		logging.SetBackend(backend1Leveled, backend2Formatter)
		Log.Debug("init log success.\r\n")
	} else {
		fmt.Println("openfile error :", err.Error())
	}
}
