package _http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/no-src/log"
	"github.com/no-src/usword/http/const"
	"github.com/no-src/usword/res/lang"
)

// HttpClient Http客户端实现
type HttpClient struct {
	client *http.Client
}

// NewHttpClient 新建一个HTTP客户端实例
func NewHttpClient() *HttpClient {
	httpClient := &HttpClient{}
	httpClient.client = &http.Client{}
	return httpClient
}

// Get 发起Get请求
func (httpClient *HttpClient) Get(url string, body interface{}, contentType string, header, cookies map[string]string) (responseData []byte, err error) {
	responseData, err = httpClient.HttpSend(_http.GET, _http.HttpVersionDefault, url, contentType, body, header, cookies)
	return responseData, err
}

// Post 发起Post请求
func (httpClient *HttpClient) Post(url string, body interface{}, contentType string, header, cookies map[string]string) (responseData []byte, err error) {
	responseData, err = httpClient.HttpSend(_http.POST, _http.HttpVersionDefault, url, contentType, body, header, cookies)
	return responseData, err
}

// PostJson 发起Post请求,提交JSON数据
func (httpClient *HttpClient) PostJson(url string, body interface{}) (responseData []byte, err error) {
	responseData, err = httpClient.Post(url, body, _http.ApplicationJson, nil, nil)
	return responseData, err
}

// PostForm 发起Post请求,提交表单数据
func (httpClient *HttpClient) PostForm(url string, body interface{}) (responseData []byte, err error) {
	responseData, err = httpClient.Post(url, body, _http.HttpVersionDefault, nil, nil)
	return responseData, err
}

// HttpSend 发起HTTP请求
// method 请求方法GET、POST等
// protocol http协议版本
// url 请求的地址
// contentType 设置Content-Type请求头
// reqData 请求数据
// headers 请求头
// cookies 请求的cookies
func (httpClient *HttpClient) HttpSend(method, protocol, url, contentType string, reqData interface{}, headers, cookies map[string]string) (responseData []byte, err error) {
	var resp *http.Response
	var req *http.Request
	var body io.Reader
	method = strings.ToUpper(method)
	// 处理协议
	if len(protocol) == 0 {
		protocol = _http.HttpVersionDefault
	}
	protocolName, protoMajor, protoMinor := analyzeProtocol(protocol)

	// 检查url地址
	lowerUrl := strings.ToLower(url)
	if strings.Index(lowerUrl, "http://") != 0 && strings.Index(lowerUrl, "https://") != 0 {
		url = protocolName + "://" + url
	}
	switch method {
	case _http.GET:
		switch reqData.(type) {
		case nil:
			break
		case map[string]string:
			url = buildQueryStringWithUrl(reqData.(map[string]string), url)
			break
		case string:
			url = joinQueryString(url, reqData.(string))
			break
		}

		req, err = http.NewRequest(method, url, body)
		break
	case _http.POST:
		var body io.Reader
		switch reqData.(type) {
		case nil:
			break
		case []byte:
			if contentType == _http.ApplicationXWwwFormUrlencoded {
				body = bytes.NewBuffer(buildFormString(convertJsonToMap(string(reqData.([]byte)))))
			} else {
				body = bytes.NewBuffer(reqData.([]byte))
			}
			break
		case string:
			if contentType == _http.ApplicationXWwwFormUrlencoded {
				body = bytes.NewBuffer(buildFormString(convertJsonToMap(reqData.(string))))
			} else {
				body = bytes.NewBuffer([]byte(reqData.(string)))
			}
			break
		case map[string]string:
			if contentType == _http.ApplicationXWwwFormUrlencoded {
				body = bytes.NewBuffer(buildFormString(reqData.(map[string]interface{})))
			} else {
				bytesData, _ := json.Marshal(reqData)
				body = bytes.NewBuffer(bytesData)
			}
			break
		default:
			bytesData, _ := json.Marshal(reqData)
			body = bytes.NewBuffer(bytesData)
			break
		}
		req, err = http.NewRequest(method, url, body)
		break
	default:
		return nil, _http.MethodNotSupported
	}

	if err != nil {
		return nil, err
	}

	// 设置HTTP协议版本
	req.Proto = protocol
	req.ProtoMajor = protoMajor
	req.ProtoMinor = protoMinor

	// 设置Cookie
	if len(cookies) > 0 {
		req.Header.Add(_http.HeaderCookie, buildCookieString(cookies))
	}

	//设置ContentType
	if len(contentType) > 0 {
		if headers == nil {
			headers = make(map[string]string)
		}
		headers[_http.HeaderContentType] = contentType
	}

	// 设置Header
	if len(headers) > 0 {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	resp, err = httpClient.client.Do(req)

	if err != nil {
		log.Error(err, "[%s]%s[%s]", method, lang.HTTP_Error_RequestFailed, url)
		return nil, err
	}

	defer resp.Body.Close()

	responseData, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err, "[%s]%s[%s]", method, lang.HTTP_Error_ReadBodyFailed, url)
		return responseData, err
	}
	return responseData, err
}
