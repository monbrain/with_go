package network_test

import (
	"fmt"
	// "log"
	"net/url"
	"testing"
	"with_go/env"
	"with_go/network"
)

func TestHttp(t *testing.T) {
	method := "get"
	uri := "https://www.google.co.kr/"
	// uri := "https://docs.upbit.com/reference/"
	header := map[string]string{}
	data := url.Values{}
	// data := ""
	resp, err := network.Http(method, uri, header, data)
	if err != nil {
		t.Error("Wrong result")
	}
	fmt.Printf("%T, %v\n", resp, string(resp[:]))
}

func TestHttpBearer(t *testing.T) {
	method := "get"
	uri := "https://api.upbit.com/v1/accounts"
	header := map[string]string{}
	data := url.Values{}
	keys, err := env.GetYmlConfig("coin_apis", "upbit")
	bearer := network.UpbitJwt(keys["access_key"].(string), keys["secret_key"].(string))
	resp, err := network.HttpBearer(method, uri, header, data, bearer)
	if err != nil {
		t.Error("Wrong result")
	}
	fmt.Printf("%T, %v\n", resp, string(resp[:]))
}
