package wecom

import (
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
	ErrorCode uint8  `json:"errcode,omitempty"`
	ErrorMsg  string `json:"errmsg,omitempty"`
}

type WeComTokenResponse struct {
	WeComErrorResponse
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

// GetToken
// docs: https://developer.work.weixin.qq.com/document/path/90664
func (wecom *WeCom) GetToken() (*WeComTokenResponse, error) {
	url := API + fmt.Sprintf("/gettoken?corpid=%s&corpsecret=%s", wecom.CorpId, wecom.CorpSecret)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(res.Body)
	//
	var resp *WeComTokenResponse
	json.Unmarshal(data, &resp)
	//
	if resp.ErrorCode == 0 {
		err = errors.New(resp.ErrorMsg)
	}
	return resp, err
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
