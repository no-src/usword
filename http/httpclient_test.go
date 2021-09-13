package _http

import (
	"github.com/no-src/log"
	"testing"
)

func TestGet(t *testing.T) {
	url := "https://golang.org/"
	httpClient := NewHttpClient()
	responseData, err := httpClient.Get(url, nil, "", nil, nil)
	if err != nil {
		log.Error(err, "Get Request Failed:%s", url)
		t.Fail()
	}
	content := string(responseData)
	log.Debug("Get Request:%s,Result:%s", url, content)
}

func TestPostJson(t *testing.T) {
	url := "https://golang.org/"
	httpClient := NewHttpClient()
	responseData, err := httpClient.PostJson(url, nil)
	if err != nil {
		log.Error(err, "Post Request Failed：%s", url)
		t.Fail()
	}
	content := string(responseData)
	log.Debug("Post Request:%s,Result:%s", url, content)
}

func init() {
	log.InitDefaultLogger(log.NewConsoleLogger(log.DebugLevel))
}
