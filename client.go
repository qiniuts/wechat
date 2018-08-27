package wechat

import (
	"fmt"
	"time"

	"github.com/qiniu/x/rpc.v7"
)

var (
	ApiHost string = "https://qyapi.weixin.qq.com"
)

type WechatWork struct {
	agentId    int
	corpId     string
	corpSecret string
	token      AccessToken
	Client     rpc.Client
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	ExpiresAt   int64  `json:"expires_at"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}

func NewWechatWork(corpId string, secretKey string, agentId int) (w *WechatWork, err error) {
	w = &WechatWork{
		corpId:     corpId,
		corpSecret: secretKey,
		agentId:    agentId,
		Client:     rpc.DefaultClient,
	}
	w.token, err = w.GetAccessToken()
	return
}

//提前一分钟判断为token失效
func (t *AccessToken) expired() bool {
	return t.ExpiresAt <= time.Now().Unix()+60
}

func (w *WechatWork) GetAccessToken() (token AccessToken, err error) {

	token = w.token
	if !token.expired() {
		return
	}

	token, err = w.refreshToken()
	return
}

func (w *WechatWork) refreshToken() (token AccessToken, err error) {

	url1 := fmt.Sprintf("%s/cgi-bin/gettoken?corpid=%s&corpsecret=%s", ApiHost, w.corpId, w.corpSecret)
	err = w.Client.CallWithJson(nil, &token, "GET", url1, nil)
	if err != nil {
		return
	}

	token.ExpiresAt = time.Now().Unix() + token.ExpiresIn
	w.token = token
	return
}
