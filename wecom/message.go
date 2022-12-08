package wecom

import (
	"encoding/json"
	"net/http"
)

type WeComMessageResponse struct {
	WeComErrorResponse
	MsgId string `json:"msgid"`
}

type WeComMessage struct {
	To      string `json:"touser"`
	MsgType string `json:"msgtype"`
	AgentId uint32 `json:"agentid"`
}

// SendMessage
// https://developer.work.weixin.qq.com/document/path/90236
func (wecom *WeCom) SendMessage(meta *WeComMessage, message map[string]any) (*WeComMessageResponse, error) {
	token, err := wecom.GetToken()
	if err != nil {
		return nil, err
	}

	var payload map[string]any
	ja, _ := json.Marshal(meta)
	jb, _ := json.Marshal(message)

	json.Unmarshal(ja, &payload)
	json.Unmarshal(jb, &payload)

	api := "/message/send?access_token=" + token.AccessToken
	data, err := wecom.SendRequest(http.MethodPost, api, payload)
	var resp *WeComMessageResponse
	json.Unmarshal(data, &resp)
	return resp, err
}

func (wecom *WeCom) RecallMessage(msgId string) (*WeComErrorResponse, error) {
	token, _ := wecom.GetToken()
	payload := map[string]string{
		"msgid": msgId,
	}
	api := "/message/recall?access_token=" + token.AccessToken
	data, err := wecom.SendRequest(http.MethodPost, api, payload)
	if err != nil {
		return nil, err
	}
	var resp *WeComErrorResponse
	json.Unmarshal(data, &resp)
	return resp, err
}

func (wecom *WeCom) SendText(meta *WeComMessage, content string) (*WeComMessageResponse, error) {
	meta.MsgType = "text"
	message := map[string]any{
		"text": map[string]any{
			"content": content,
		},
	}
	return wecom.SendMessage(meta, message)
}

func (wecom *WeCom) SendImage(meta *WeComMessage) {

}

func (wecom *WeCom) SendVideo(meta *WeComMessage) {

}

func (wecom *WeCom) SendFile(meta *WeComMessage) {

}

func (wecom *WeCom) SendTextCard(meta *WeComMessage) {

}

func (wecom *WeCom) SendNews(meta *WeComMessage) {

}

func (wecom *WeCom) SendMarkdown(meta *WeComMessage) {

}
