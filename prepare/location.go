package prepare

const (
	SOURCE_LOCATION = "www.hao123.com"
)

func loadSourceLocation() {
	//http.Post(SOURCE_LOCATION);
}

func getMyLocation() map[string]string {
	var locationMap map[string]string
	locationMap = make(map[string]string)
	locationMap["location1"] = "a"
	locationMap["location2"] = "b"
	//locationMap[ =""
	return locationMap
}
