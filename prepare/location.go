package prepare

func LoadSourceLocation() map[string]DestinationReq {
	return getMyLocation()
}

func getMyLocation() map[string]DestinationReq {
	var locationMap map[string]DestinationReq
	locationMap = make(map[string]DestinationReq)
	locationMap["localTest"] = DestinationReq{
		doName: "http://192.168.128.149",
		uri:    "/uat/get",
		method: "Get",
		param:  map[string]string{"param": "yes"},
	}

	return locationMap
}
