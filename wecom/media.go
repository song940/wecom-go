package wecom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
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
func (wecom *WeCom) Upload(filename string) (*WeComMediaResponse, error) {

	var (
		buf = new(bytes.Buffer)
		w   = multipart.NewWriter(buf)
	)

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	data, _ := io.ReadAll(f)
	part, _ := w.CreateFormFile("media", filepath.Base(f.Name()))
	part.Write(data)

	token, _ := wecom.GetToken()
	url := API + fmt.Sprintf("/media/upload?access_token=%s&type=%s", token.AccessToken, "image")
	req, _ := http.NewRequest(http.MethodPost, url, buf)
	req.Header.Set("Content-Type", w.FormDataContentType())
	res, err := wecom.Client.Do(req)
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(res.Body)

	log.Println(string(body))

	var resp *WeComMediaResponse
	json.Unmarshal(body, &resp)
	return resp, err
}
