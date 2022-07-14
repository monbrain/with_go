package network

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Http(method string, url string, header map[string]string, data url.Values) ([]byte, error) {
	method = strings.ToUpper(method)
	client := http.Client{Timeout: 10 * time.Second}

	var req *http.Request
	switch method {
	case "GET":
		reqUrl := url
		if data.Encode() != "" {
			reqUrl = url + data.Encode()
		}
		req, _ = http.NewRequest(method, reqUrl, nil)
	default: // case "POST" / "DELETE"
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req, _ = http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))
	}

	for k, v := range header {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		msg := "HTTP " + method + " Error"
		log.Fatal(msg, err, resp.StatusCode)
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func HttpBearer(method string, url string, header map[string]string, data url.Values, bearer string) ([]byte, error) {
	header["Authorization"] = "Bearer " + bearer
	return Http(method, url, header, data)
}

// func HttpPOST(url string, header map[string]string, data url.Values) ([]byte, error) {
// 	client := http.Client{Timeout: 10 * time.Second}

// 	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	//println(data.Encode())
// 	for k, v := range header {
// 		req.Header.Add(k, v)
// 	}

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal("[HttpPost] HTTP POST Error", err, resp.StatusCode)
// 		return nil, err
// 	}

// 	defer resp.Body.Close()
// 	return ioutil.ReadAll(resp.Body)
// }

// func HttpDelete(url string, header map[string]string, data url.Values) ([]byte, error) {
// 	client := http.Client{Timeout: 10 * time.Second}

// 	req, err := http.NewRequest(http.MethodDelete, url, strings.NewReader(data.Encode()))
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	for k, v := range header {
// 		req.Header.Add(k, v)
// 	}

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal("[HttpGet] HTTP DELETE Error", err, resp.StatusCode)
// 		return nil, err
// 	}

// 	defer resp.Body.Close()
// 	return ioutil.ReadAll(resp.Body)
// }
