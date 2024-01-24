package wecom

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	API = "https://qyapi.weixin.qq.com/cgi-bin"
)

type WeComClientConfig struct {
	CorpId     string
	CorpSecret string
}

type WeComClient struct {
	client *http.Client
	config *WeComClientConfig
}

type WeComErrorResponse struct {
	ErrorCode uint32 `json:"errcode,omitempty"`
	ErrorMsg  string `json:"errmsg,omitempty"`
}

type WeComTokenResponse struct {
	WeComErrorResponse
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func NewClient(config *WeComClientConfig) (wecom *WeComClient) {
	client := &http.Client{}
	wecom = &WeComClient{
		client: client,
		config: config,
	}
	return
}

// GetToken
// docs: https://developer.work.weixin.qq.com/document/path/90664
func (wecom *WeComClient) GetToken() (resp *WeComTokenResponse, err error) {
	api := fmt.Sprintf("/gettoken?corpid=%s&corpsecret=%s", wecom.config.CorpId, wecom.config.CorpSecret)
	data, err := wecom.SendRequest(http.MethodGet, api, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	if resp.ErrorCode != 0 {
		err = errors.New(resp.ErrorMsg)
	}
	return
}

func (wecom *WeComClient) SendRequest(method string, path string, body any) (data []byte, err error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return
	}
	// log.Println("SendRequest", method, url, string(payload))
	req, _ := http.NewRequest(method, API+path, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	res, err := wecom.client.Do(req)
	if err != nil {
		return nil, err
	}
	errorCode := res.Header.Get("error-code")
	errorMsg := res.Header.Get("error-msg")
	// log.Println("response:", errorCode, errorMsg)
	if errorCode != "0" {
		err = errors.New(errorMsg)
		return
	}
	data, err = io.ReadAll(res.Body)
	return
}
