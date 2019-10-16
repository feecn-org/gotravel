package travel

import (
	"github.com/wonderivan/logger"
)

func Mark(destin map[string]string) string {
	logger.Debug(destin)
	for k, v := range destin {
		logger.Debug(k, v)
	}
	return "marking"
}

func runMarkDestion(k string, v string) {
	go run(k, v)
}

func run(k string, v string) {
	//execute request get and download data
	//tool.Http_Get()

}
