package tool

import (
	"fmt"
	"github.com/wonderivan/logger"
	"strconv"
	"testing"
	"time"
)

func Test_Get(t *testing.T) {
	h := NewHttpSend(GetUrlBuild("https://jobs.zhaopin.com/000054844250055.htm", map[string]string{}))
	_, err := h.Get()
	if err != nil {
		t.Error("请求错误:", err)
	} else {
		t.Log("正常返回")
	}
}

func Test_Get2(t *testing.T) {
	h := NewHttpSend(GetUrlBuild("http://192.168.128.149:80/uat/get", map[string]string{"param": "yes"}))
	response, err := h.Get()
	logger.Debug(string(response))
	if err != nil {
		t.Error("请求错误:", err)
	} else {
		t.Log("正常返回")
	}
}

func Test_Post(t *testing.T) {
	h := NewHttpSend("http://127.0.0.1/test.php")
	h.SetBody(map[string]string{"name": "xiaochuan"})
	_, err := h.Post()
	if err != nil {
		t.Error("请求错误:", err)
	} else {
		t.Log("正常返回")
	}
}

func Test_Json(t *testing.T) {
	h := NewHttpSend("http://127.0.0.1/test.php")
	h.SetSendType("JSON")
	h.SetBody(map[string]string{"name": "xiaochuan"})
	_, err := h.Post()
	if err != nil {
		t.Error("请求错误:", err)
	} else {
		t.Log("正常返回")
	}
}

func Benchmark_GET(b *testing.B) {
	formatInt := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	logger.Debug(formatInt)
	now := time.Now()
	year, mon, day := now.UTC().Date()
	hour, min, sec := now.UTC().Clock()
	zone, _ := now.UTC().Zone()
	fmt.Printf("UTC 时间是 %d-%d-%d %02d:%02d:%02d %s\n",
		year, mon, day, hour, min, sec, zone)
}
