package travel

import (
	"github.com/wonderivan/logger"
	"gotravel/prepare"
	"gotravel/tool"
)

func Mark(destin map[string]prepare.DestinationReq) string {
	logger.Debug(destin)
	for k, v := range destin {
		runMarkDestion(k, v)
	}
	return "marking"
}

func runMarkDestion(k string, v prepare.DestinationReq) {
	run(k, v)
}

func run(k string, v prepare.DestinationReq) []byte {
	//execute request get and download data
	logger.Debug(k, v)
	link := v.DoName() + v.Uri()
	param := v.Param()
	method := v.Method()
	if method == "Get" {
		h := tool.NewHttpSend(tool.GetUrlBuild(link, param))
		response, err := h.Get()
		logger.Debug(string(response))
		if err != nil {
			logger.Error("请求错误:", err)
		} else {
			logger.Debug("正常返回")
			return response
		}
	} else if method == "Post" {
		h := tool.NewHttpSend("http://127.0.0.1/test.php")
		h.SetBody(map[string]string{"name": "xiaochuan"})
		response, err := h.Post()
		if err != nil {
			logger.Error("请求错误:", err)
		} else {
			logger.Debug("正常返回")
			return response
		}
	}
	return nil
}
