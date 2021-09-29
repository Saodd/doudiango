package doudianapi

import (
	"encoding/json"
	"github.com/saodd/alog"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

var (
	client  *Client
	secrets *SecretsDoudian
)

type ShopDoudian struct {
	AccessToken string `json:"access_token"`
	//Iids        []int  `json:"iids"`
}

type SecretsDoudian struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`

	Shop1 ShopDoudian `json:"shop_1"`
}

func getSecrets() {
	secrets = new(SecretsDoudian)
	resp, err := http.Get("http://localhost:26666/secret/doudian")
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if err := resp.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, secrets); err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	getSecrets()
	client = &Client{
		AppKey:             secrets.AppKey,
		AppSecret:          secrets.AppSecret,
		HttpDo:             http.DefaultClient.Do,
		DebugFlag:          true,
		DoudianApiEndpoint: "https://openapi-fxg.jinritemai.com",
		HandleError:        alog.CE,
		DebugPrintf:        log.Printf,
	}
	os.Exit(m.Run())
}
