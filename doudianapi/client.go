package doudianapi

import (
	"context"
	"doudiango/doudianutils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Client struct {
	AppKey             string
	AppSecret          string
	DoudianApiEndpoint string
	HttpDo             func(req *http.Request) (*http.Response, error)

	// HandleError 是用来处理错误的函数。推荐使用 github.com/saodd/alog.CE
	HandleError func(context.Context, error, ...map[string]interface{})

	DebugFlag   bool
	DebugPrintf func(format string, v ...interface{})
}

// NewClient 提供常用的配置。如果需要更多特性，请自己实例化。
func NewClient(appKey string, appSecret string, debug bool) *Client {
	return &Client{
		AppKey:             appKey,
		AppSecret:          appSecret,
		DoudianApiEndpoint: "https://openapi-fxg.jinritemai.com",
		HttpDo:             http.DefaultClient.Do,
		HandleError:        func(c context.Context, err error, v ...map[string]interface{}) { log.Println(err, v) },
		DebugFlag:          debug,
		DebugPrintf:        log.Printf,
	}
}

// Do 快捷方法，执行一个请求
func (client *Client) Do(ctx context.Context, params RequestParams, method string, session string, res ResponseData) (err error) {
	// 0. Debug输出
	var u string
	var reqBody string
	var respBody []byte
	if client.DebugFlag {
		defer func() {
			client.DebugOutput(u, reqBody, respBody, err)
		}()
	}
	timestamp := time.Now().Unix()

	// 1. 序列化参数
	reqBody = doudianutils.Marshal(params.ToSignMap())
	// 2. 计算签名
	signVal := doudianutils.Sign(client.AppKey, client.AppSecret, method, timestamp, reqBody)
	// 3. 构建请求
	req, err := doudianutils.BuildFetchRequest(client.AppKey, client.DoudianApiEndpoint, method, timestamp, reqBody, session, signVal)
	if err != nil {
		client.HandleError(ctx, err)
		return err
	}
	u = req.URL.String()
	// 4. 执行请求
	resp, err := client.HttpDo(req)
	if err != nil {
		client.HandleError(ctx, err, nil)
		return err
	}
	respBody, _ = ioutil.ReadAll(resp.Body)
	if err = resp.Body.Close(); err != nil {
		client.HandleError(ctx, err)
		return err
	}

	// 3. 解析错误
	if err = json.Unmarshal(respBody, res); err != nil {
		return err
	}
	if err := res.Err(); err != nil {
		return err
	}

	// 4. 解析结果
	return json.Unmarshal(respBody, res)
}

type RequestParams interface {
	// ToSignMap 提供需要签名的参数
	ToSignMap() map[string]interface{}

	// Valid 检查参数是否合法
	Valid() error
}
type ResponseData interface {
	error
	Err() error
}

func (client *Client) DebugOutput(u, reqBody string, respBody []byte, err error) {
	const template = `[TAOBAOAPI-DEBUG]
  ***Request:
    POST %s
    %s
  ***Response:
    %s
  ***Error:
    %s
  ***END***
`
	client.DebugPrintf(template, u, reqBody, string(respBody), err)
}

// SystemError 封装了抖店返回的错误信息。
type SystemError struct {
	ErrNo   int    `json:"err_no"`
	LogId   string `json:"log_id"`
	Message string `json:"message"`
}

func (e *SystemError) Error() string {
	return fmt.Sprintf("DoudianAPIError: %s | %d | %s", e.Message, e.ErrNo, e.LogId)
}

func (e *SystemError) Err() error {
	if e.ErrNo != 0 {
		return e
	}
	return nil
}
