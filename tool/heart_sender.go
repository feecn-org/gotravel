package tool

//
//import (
//	"net"
//	"os"
//	"net/http"
//	"strings"
//	"encoding/json"
//	"github.com/spf13/viper"
//	log "travelsky.com/agent/logger"
//	"time"
//)
//
//const (
//	HEARTBEAT  = "heartBeat"
//	CONTENT_TYPE = "application/json"
//	SEND_URL  = "heart_send_url"
//	SEND_INTERVAL = "interval"
//	LOG_AGENT = "logagent"
//	AGENT_TYPE = "log"
//)
//
//type HeartInfo struct {
//	AgentName string
//	AgentType string
//	AgentIp string
//}
//
//func sendHeart(){
//	heartBeat := viper.GetStringMap(HEARTBEAT)
//	sendUrl := heartBeat[SEND_URL]
//	sendInterval := heartBeat[SEND_INTERVAL]
//	heartInfo := HeartInfo{
//		AgentName:LOG_AGENT,
//		AgentType:AGENT_TYPE,
//		AgentIp:GetHostIp()}
//	//timeconf := time.Second * sendInterval.  (float64)
//	//log.Info.Println(timeconf)
//	t := time.NewTicker(time.Duration(sendInterval.(float64)) * time.Second)
//	heartJson, errs := json.Marshal(heartInfo)
//	if errs != nil{
//		log.Error.Println(errs)
//	}
//	for {
//		select {
//		case <-t.C:
//			resp , errs := http.Post(sendUrl.(string),CONTENT_TYPE,strings.NewReader(string(heartJson)))
//			if resp == nil{
//				_=errs
//				//log.Error.Println(errs)
//			}
//		}
//	}
//}
//
////获取本机ip
//func GetHostIp() string{
//	addrs, err := net.InterfaceAddrs()
//	if err != nil {
//		os.Exit(1)
//	}
//	for _, address := range addrs {
//		// 检查ip地址判断是否回环地址
//		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
//			if ipnet.IP.To4() != nil {
//				return ipnet.IP.String()
//			}
//		}
//	}
//	return ""
//}
//
//
