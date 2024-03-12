package wecom

import (
	"encoding/json"
	"net/http"
)

type WeComGroup struct {
}

type WeComGroupResponse struct {
	WeComErrorResponse
	ChatId string `json:"chatid"`
}

// https://developer.work.weixin.qq.com/document/path/90245
func (wecom *Client) CreateGroup(group *WeComGroup) (*WeComGroupResponse, error) {
	token, _ := wecom.GetToken()
	api := "/appchat/create?access_token=" + token.AccessToken
	data, err := wecom.SendRequest(http.MethodPost, api, group)
	var resp *WeComGroupResponse
	json.Unmarshal(data, &resp)
	return resp, err
}

// https://developer.work.weixin.qq.com/document/path/90246
func (wecom *Client) UpdateGroupInfo() {
	// https://qyapi.weixin.qq.com/cgi-bin/appchat/update?access_token=ACCESS_TOKEN
}

// https://developer.work.weixin.qq.com/document/path/90247
func (wecom *Client) GetGroupInfo() {
	// https://qyapi.weixin.qq.com/cgi-bin/appchat/get?access_token=ACCESS_TOKEN&chatid=CHATID
}

// https://developer.work.weixin.qq.com/document/path/90248
func (wecom *Client) GroupSend() {
	// https://qyapi.weixin.qq.com/cgi-bin/appchat/send?access_token=ACCESS_TOKEN
}
