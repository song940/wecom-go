package cli

import (
	"log"

	"github.com/song940/wecom/wecom"
)

func Run() {

	client := wecom.NewClient(
		"wx154021007ed664e5",
		"z-W9O0wpdoUlhyneavoCa210xKQRj7qNceDeR6eGk9o",
	)

	// token, _ := client.GetToken()
	// log.Println(token.AccessToken)

	resp, err := client.SendTextMessage(&wecom.WeComMessage{
		To:      "@all",
		MsgType: "text",
		AgentId: 1,
	}, "test message golang")
	log.Println(resp.MsgId, err)
}
