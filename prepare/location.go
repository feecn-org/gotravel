package prepare

import "gotravel/tool"

const (
	SourceLocation = "www.hao123.com"
)

func LoadSourceLocation() map[string]string {
	return getMyLocation()
}

func getMyLocation() map[string]string {
	var locationMap map[string]string
	locationMap = make(map[string]string)
	locationMap["baidu"] = tool.SEARCH_BAIDU
	locationMap["location2"] = "b"
	//locationMap[ =""
	return locationMap
}
