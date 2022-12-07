package wecom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
	token, _ := wecom.GetToken()
	url := API + fmt.Sprintf("/message/send?access_token=%s", token.AccessToken)

	var payload map[string]any
	ja, _ := json.Marshal(meta)
	jb, _ := json.Marshal(message)

	json.Unmarshal(ja, &payload)
	json.Unmarshal(jb, &payload)

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	res, err := wecom.Client.Do(req)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(res.Body)
	var resp *WeComMessageResponse
	json.Unmarshal(data, &resp)
	return resp, err
}

func (wecom *WeCom) SendTextMessage(meta *WeComMessage, content string) (*WeComMessageResponse, error) {
	message := map[string]any{
		"text": map[string]any{
			"content": content,
		},
	}
	return wecom.SendMessage(meta, message)
}
