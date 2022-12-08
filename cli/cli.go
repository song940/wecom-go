package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/song940/wecom/wecom"
)

func Run() {

	args, flags := ParseArgs()
	fmt.Println(args, flags)

	if len(args) == 0 {
		panic("You must specify a command to run")
	}

	cmd := args[0]

	WECOM_CORP_ID := os.Getenv("WECOM_CORP_ID")
	WECOM_CORP_SECRET := os.Getenv("WECOM_CORP_SECRET")

	client := wecom.NewClient(
		WECOM_CORP_ID,
		WECOM_CORP_SECRET,
	)

	if cmd == "get-token" {
		token, err := client.GetToken()
		if err != nil {
			panic(err)
		}
		fmt.Println(token.AccessToken)
	}

	if cmd == "sendtext" {

		if len(args) != 2 {
			panic("missing required message parameter")
		}

		msg := args[1]
		to := flags["to"].(string)
		agentId := flags["agentId"].(string)
		n, err := strconv.ParseInt(agentId, 10, 32)
		if err != nil {
			panic(err)
		}

		meta := &wecom.WeComMessage{
			AgentId: uint32(n),
			To:      to,
		}
		resp, err := client.SendText(meta, msg)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.MsgId)
	}

	if cmd == "recall" {
		msgId := flags["msgId"].(string)
		resp, err := client.RecallMessage(msgId)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.ErrorMsg)
	}

	if cmd == "upload" {
		filename := args[1]
		resp, err := client.Upload(filename)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.MediaType, resp.MediaId)
	}
}
