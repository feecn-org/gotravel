package prepare

import "gotravel/tool"

const (
	SourceLocation = "www.hao123.com"
)

func loadSourceLocation() {
	println(tool.Http_Get(SourceLocation))
}

func getMyLocation() map[string]string {
	var locationMap map[string]string
	locationMap = make(map[string]string)
	locationMap["location1"] = "a"
	locationMap["location2"] = "b"
	//locationMap[ =""
	return locationMap
}
