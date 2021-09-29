package doudianutils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Marshal 序列化参数
// 参考自 https://op.jinritemai.com/docs/guide-docs/10/813
func Marshal(m map[string]interface{}) string {
	// 做一次序列化，并禁用Html Escape
	buffer := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	_ = encoder.Encode(m)

	marshal := string(bytes.TrimSpace(buffer.Bytes())) // Trim掉末尾的换行符
	return marshal
}

// Sign 计算签名
// 参考自 https://op.jinritemai.com/docs/guide-docs/10/813
func Sign(appKey, appSecret, method string, timestamp int64, paramJson string) string {
	// 按给定规则拼接参数
	paramPattern := "app_key" + appKey + "method" + method + "param_json" + paramJson + "timestamp" + strconv.FormatInt(timestamp, 10) + "v2"
	signPattern := appSecret + paramPattern + appSecret
	return Hmac(signPattern, appSecret)
}

// Hmac 计算hmac
// 复制自 https://op.jinritemai.com/docs/guide-docs/10/813
func Hmac(s string, appSecret string) string {
	h := hmac.New(sha256.New, []byte(appSecret))
	_, _ = h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// BuildFetchRequest 构建一个调用 Open Api 的请求对象
// 参考自 https://op.jinritemai.com/docs/guide-docs/10/813
func BuildFetchRequest(appKey, host, method string, timestamp int64, paramJson, accessToken, sign string) (*http.Request, error) {
	methodPath := strings.Replace(method, ".", "/", -1)

	params := url.Values{}
	params.Add("method", method)
	params.Add("app_key", appKey)
	params.Add("access_token", accessToken)
	params.Add("timestamp", strconv.FormatInt(timestamp, 10))
	params.Add("v", "2")
	params.Add("sign", sign)
	params.Add("sign_method", "hmac-sha256")

	u := host + "/" + methodPath + "?" + params.Encode()

	req, err := http.NewRequest(http.MethodPost, u, strings.NewReader(paramJson))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	return req, nil
}
