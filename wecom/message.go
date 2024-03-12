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
	ToUser                 string `json:"touser,omitempty"`
	ToParty                string `json:"toparty,omitempty"`
	ToTag                  string `json:"totag,omitempty"`
	MsgType                string `json:"msgtype"`
	AgentId                uint32 `json:"agentid"`
	EnableDuplicateCheck   uint32 `json:"enable_duplicate_check,omitempty"`
	DuplicateCheckInterval uint32 `json:"duplicate_check_interval,omitempty"`
}

// SendMessage
// https://developer.work.weixin.qq.com/document/path/90236
func (wecom *Client) SendMessage(meta *WeComMessage, message map[string]any) (resp *WeComMessageResponse, err error) {
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
	if err != nil {
		return
	}
	resp = &WeComMessageResponse{}
	err = json.Unmarshal(data, &resp)
	return
}

func (wecom *Client) RecallMessage(msgId string) (*WeComErrorResponse, error) {
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

func (wecom *Client) SendText(meta *WeComMessage, content string) (*WeComMessageResponse, error) {
	meta.MsgType = "text"
	message := map[string]any{
		"text": map[string]any{
			"content": content,
		},
	}
	return wecom.SendMessage(meta, message)
}

func (wecom *Client) SendImage(meta *WeComMessage) (*WeComMessageResponse, error) {
	meta.MsgType = "image"
	message := map[string]any{
		"image": map[string]any{
			"media_id": "",
		},
	}
	return wecom.SendMessage(meta, message)
}

func (wecom *Client) SendVideo(meta *WeComMessage) (*WeComMessageResponse, error) {
	meta.MsgType = "video"
	message := map[string]any{
		"video": map[string]any{
			"media_id":    "",
			"title":       "",
			"description": "",
		},
	}
	return wecom.SendMessage(meta, message)
}

func (wecom *Client) SendFile(meta *WeComMessage) (*WeComMessageResponse, error) {
	meta.MsgType = "file"
	message := map[string]any{
		"file": map[string]any{
			"media_id": "",
		},
	}
	return wecom.SendMessage(meta, message)
}

type TextCard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	ButtonText  string `json:"btntxt"`
}

func (wecom *Client) SendTextCard(meta *WeComMessage, textcard *TextCard) (*WeComMessageResponse, error) {
	meta.MsgType = "textcard"
	message := map[string]any{
		"textcard": textcard,
	}
	return wecom.SendMessage(meta, message)
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"picurl"`
	AppID       string `json:"appid,omitempty"`
	PagePath    string `json:"pagepath,omitempty"`
}

func (wecom *Client) SendNews(meta *WeComMessage, articles []Article) (*WeComMessageResponse, error) {
	meta.MsgType = "news"
	message := map[string]any{
		"news": map[string]any{
			"articles": articles,
		},
	}
	return wecom.SendMessage(meta, message)
}

func (wecom *Client) SendMarkdown(meta *WeComMessage, content string) (*WeComMessageResponse, error) {
	meta.MsgType = "markdown"
	message := map[string]any{
		"markdown": map[string]any{
			"content": content,
		},
	}
	return wecom.SendMessage(meta, message)
}
