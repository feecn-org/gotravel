package prepare

func LoadSourceLocation() map[string]DestinationReq {
	return getMyLocation()
}

func getMyLocation() map[string]DestinationReq {
	var locationMap map[string]DestinationReq
	locationMap = make(map[string]DestinationReq)

	//http://127.0.0.1:8080/uat/get?param=yes&and=jack

	//locationMap["localTest"] = DestinationReq{
	//	doName: "http://127.0.0.1:8080",
	//	uri:    "/uat/get",
	//	method: "Get",
	//	param:  map[string]string{"param": "yes","and":"jack"},
	//}

	//locationMap["localTest"] = DestinationReq{
	//	doName: "https://jobs.zhaopin.com/CC300184216J00095339615.htm",
	//	method: "Get",
	//	param:  map[string]string{},
	//}

	locationMap["zhaopin"] = DestinationReq{
		doName: "https://fe-api.zhaopin.com",
		uri:    "/c/i/sou",
		method: "Get",
		param: map[string]string{"pageSize": "90", "cityId": "530", "workExperience": "-1", "education": "-1", "companyType": "-1", "employmentType": "-1",
			"jobWelfareTag": "-1", "kw": "java", "kt": "3", "_v": "0.45623841", "x-zp-page-request-id": "7ebedd4821b54a5a8f0b8471356e5dfb-1571280451080-760373", "x-zp-client-id": "f074fd10-6474-450f-8f4c-3826a9114f39"},
	}

	return locationMap
}
