package cli

import (
	"log"
	"os"

	"github.com/song940/wecom/wecom"
)

func Run() {
	WECOM_CORP_ID := os.Getenv("WECOM_CORP_ID")
	WECOM_CORP_SECRET := os.Getenv("WECOM_CORP_ID")

	client := wecom.NewClient(
		WECOM_CORP_ID,
		WECOM_CORP_SECRET,
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
