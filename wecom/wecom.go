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

type WeCom struct {
	Client *http.Client

	CorpId     string
	CorpSecret string
}

type WeComErrorResponse struct {
	ErrorCode uint32 `json:"errcode,omitempty"`
	ErrorMsg  string `json:"errmsg,omitempty"`
}

type WeComTokenResponse struct {
	WeComErrorResponse
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

func NewClient(corpId string, corpSecret string) (wecom *WeCom) {
	client := &http.Client{}
	wecom = &WeCom{
		Client:     client,
		CorpId:     corpId,
		CorpSecret: corpSecret,
	}
	return
}

// GetToken
// docs: https://developer.work.weixin.qq.com/document/path/90664
func (wecom *WeCom) GetToken() (*WeComTokenResponse, error) {
	api := fmt.Sprintf("/gettoken?corpid=%s&corpsecret=%s", wecom.CorpId, wecom.CorpSecret)
	data, err := wecom.SendRequest(http.MethodGet, api, nil)
	if err != nil {
		return nil, err
	}
	var resp *WeComTokenResponse
	json.Unmarshal(data, &resp)
	//
	if resp.ErrorCode != 0 {
		err = errors.New(resp.ErrorMsg)
	}
	return resp, err
}

func (wecom *WeCom) SendRequest(method string, path string, body any) ([]byte, error) {
	url := API + path
	payload, _ := json.Marshal(body)
	// log.Println("SendRequest", method, url, string(payload))
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	// errorCode := res.Header.Get("error-code")
	// errorMsg := res.Header.Get("error-msg")
	// if errorCode != "0" {
	// 	return nil, errors.New(errorMsg)
	// }
	data, err := io.ReadAll(res.Body)
	return data, err
}
