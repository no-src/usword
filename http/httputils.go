package _http

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// buildCookieString 构造Cookie值
func buildCookieString(dict map[string]string) string {
	result := ""
	for k, v := range dict {
		result = fmt.Sprintf("%s%s=%s;", result, k, v)
	}
	return result
}

// buildQueryStringWithUrl 构造QueryString
func buildQueryStringWithUrl(dict map[string]string, url string) string {
	if len(dict) == 0 {
		return url
	}
	queryString := buildQueryString(dict)
	url = joinQueryString(url, queryString)
	return url
}

// joinQueryString 拼接URL地址
func joinQueryString(url, queryString string) string {
	if len(queryString) == 0 {
		return url
	}
	if strings.Index(url, "?") < 0 {
		queryString = "?" + queryString[1:]
	}
	url = url + queryString
	return url
}

// buildQueryString 构造QueryString
func buildQueryString(dict map[string]string) string {
	result := ""
	for k, v := range dict {
		result = fmt.Sprintf("%s&%s=%s", result, k, v)
	}
	return result
}

// analyzeProtocol 解析http协议版本中的版本号
func analyzeProtocol(protocol string) (protocolName string, protoMajor, protoMinor int) {
	versionInfo := strings.Split(protocol, "/")
	if len(versionInfo) < 2 {
		return protocolName, protoMajor, protoMinor
	}
	protocolName = versionInfo[0]
	version := strings.Split(versionInfo[1], ".")
	if len(version) < 2 {
		return protocolName, protoMajor, protoMinor
	}
	protoMajor, _ = strconv.Atoi(version[0])
	protoMinor, _ = strconv.Atoi(version[1])
	return protocolName, protoMajor, protoMinor
}

// buildFormString 构造Form表单数据
func buildFormString(data map[string]interface{}) (bytes []byte) {
	str := ""
	if len(data) == 0 {
		return bytes
	}
	for k, v := range data {
		vStr := ""
		vBytes, err := json.Marshal(v)
		if err == nil {
			vStr = string(vBytes)
		}
		str += k + "=" + vStr + "&"
	}
	str = str[:len(str)-1]
	bytes = []byte(str)
	return bytes
}

// convertObjectToMap 将对象转换为字典
func convertObjectToMap(obj interface{}) (dict map[string]string) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return dict
	}
	dict = make(map[string]string)
	err = json.Unmarshal(bytes, &dict)
	if err != nil {
		return nil
	}
	return dict
}

// convertJsonToMap 将json转换为字典
func convertJsonToMap(jsonStr string) (dict map[string]interface{}) {
	bytes := []byte(jsonStr)
	dict = make(map[string]interface{})
	err := json.Unmarshal(bytes, &dict)
	if err != nil {
		return nil
	}
	return dict
}
