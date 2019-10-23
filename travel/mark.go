package travel

import (
	"encoding/json"
	"github.com/wonderivan/logger"
	"gotravel/konwledge"
	"gotravel/prepare"
	"gotravel/tool"
)

var callCount = 0

var failedCount = 0

func Mark(destination map[string]prepare.DestinationReq) string {
	logger.Debug(destination)
	for k, v := range destination {
		_ = k
		runMarkDestination(v)
	}
	return "marking"
}

func runMarkDestination(v prepare.DestinationReq) {
	destinations := string(run(v))
	logger.Debug(destinations)
	var source map[string]interface{}
	if err := json.Unmarshal([]byte(destinations), &source); err == nil {
		if source["code"].(float64) == 200 {
			results := source["data"].(map[string]interface{})["results"].([]interface{})
			for k, v := range results {
				logger.Debug(k, v)
				markInfo := v.(map[string]interface{})
				positionURL := markInfo["positionURL"]
				markInfoJson, _ := json.Marshal(markInfo)
				mString := string(markInfoJson)
				mark := &konwledge.Mark{
					Body: mString,
					Page: positionURL.(string),
				}
				konwledge.InsertMark(mark)
				var destination prepare.DestinationReq
				destination.SetDoName(positionURL.(string))
				destination.SetMethod("Get")
				logger.Debug(positionURL)
				runDocuments(destination)
			}
		} else {
			logger.Error("runMarkDestination failed , err msg is ", err)
		}
		logger.Debug(source)
	} else {
		logger.Error(err)
	}
}

func runDocuments(v prepare.DestinationReq) []byte {
	bytes := run(v)
	return bytes
}

func run(v prepare.DestinationReq) []byte {
	//execute request get and download data
	link := v.DoName() + v.Uri()
	param := v.Param()
	method := v.Method()
	if method == "Get" {
		h := tool.NewHttpSend(tool.GetUrlBuild(link, param))
		response, err := h.Get()
		//node, err := html.Parse(bytes.NewReader(response))
		//logger.Info(node.Data)
		logger.Debug(string(response))
		if err != nil {
			logger.Error("请求错误:", err)
			failedCount = failedCount + 1
			logger.Debug("错误返回 request times = ", failedCount)
		} else {
			callCount = callCount + 1

			logger.Debug("正常返回 request times = ", callCount)
			return response
		}
	} else if method == "Post" {
		h := tool.NewHttpSend(link)
		h.SetBody(param)
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
