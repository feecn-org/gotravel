package tool

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	ContentType     = "Content-Type"
	ContentTypeJson = "application/json"
	ContentTypeForm = "application/x-www-form-urlencoded"
	MethodPost      = "POST"
	MethodGet       = "Get"
	Cookie          = "cookie"
)

/**
http get request
*/
func Http_Get(getUrl string) string {
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
func Http_Post(postUrl string) string {
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

/**
http post form
*/
func Http_Post_Form(postUrl string) string {
	resp, err := http.PostForm(postUrl,
		url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
	return string(body)
}

/**
http request with header and cookie
*/
func httpDo(method string, urlStr string, header string, cookie string) string {
	client := &http.Client{}

	req, err := http.NewRequest(method, urlStr, strings.NewReader(header))
	if err != nil {
		// handle error
	}

	req.Header.Set(ContentType, ContentTypeForm)
	//req.Header.Set(Cookie, "name=anny")
	req.Header.Set(Cookie, cookie)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
	return string(body)
}
