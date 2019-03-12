package tool

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	ContentTypeJson = "application/json"
	ContentTypeForm = "application/x-www-form-urlencoded"
)

/**
http get request
*/
func Http_get(getUrl string) string {
	resp, err := http.Get(getUrl)
	if err != nil {
	}
	err = resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
	return string(body)
}

/**
http post request json
*/
func Http_post(postUrl string) string {
	resp, err := http.Post(postUrl, ContentTypeJson, strings.NewReader("body"))
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
	return string(body)
}
