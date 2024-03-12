package wecom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type WeComMediaResponse struct {
	WeComErrorResponse
	MediaType string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}

// Upload file
// https://developer.work.weixin.qq.com/document/path/90253
func (wecom *Client) Upload(filename string) (resp *WeComMediaResponse, err error) {
	var (
		buf = new(bytes.Buffer)
		w   = multipart.NewWriter(buf)
	)
	data, _ := os.ReadFile(filename)
	part, _ := w.CreateFormFile("media", filepath.Base(filename))
	part.Write(data)
	token, _ := wecom.GetToken()
	url := API + fmt.Sprintf("/media/upload?access_token=%s&type=%s", token.AccessToken, "image")
	req, _ := http.NewRequest(http.MethodPost, url, buf)
	req.Header.Set("Content-Type", w.FormDataContentType())
	res, err := wecom.client.Do(req)
	if err != nil {
		return
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	// log.Println(string(body))
	resp = &WeComMediaResponse{}
	err = json.Unmarshal(body, &resp)
	return
}
