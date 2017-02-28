package goRabbitMq

import (
	"bytes"
	"encoding/gob"
	"goRabbitMq/common"
	"strconv"

	"github.com/pquerna/ffjson/ffjson"
)

//如果条件为true，返回a，否则返回b
func If(condition bool, a, b interface{}) interface{} {
	if condition {
		return a
	} else {
		return b
	}
}

func ToBytes(v interface{}) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	if err != nil {
		common.Log.Errorf("message body encode error %s", err.Error())
	}
	by := buf.Bytes()
	return by
}

func JsonToBytes(v interface{}) []byte {
	by, err := ffjson.Marshal(v)
	if err != nil {
		common.Log.Errorf("message body encode error %s", err.Error())
	}
	return by
}

func ToString(args ...interface{}) string {
	result := ""
	for _, arg := range args {
		switch val := arg.(type) {
		case int:
			result += strconv.Itoa(val)
		case string:
			result += val
		case float64:
			result += strconv.FormatFloat(val, 'f', 6, 64)
		}
	}
	return result
}
