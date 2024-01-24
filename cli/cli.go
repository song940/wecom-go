package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/song940/wecom/wecom"
)

func Run() {

	args, flags := ParseArgs()
	if len(args) == 0 {
		panic("You must specify a command to run")
	}

	cmd := args[0]

	config := &wecom.WeComClientConfig{
		CorpId:     os.Getenv("WECOM_CORP_ID"),
		CorpSecret: os.Getenv("WECOM_CORP_SECRET"),
	}
	client := wecom.NewClient(config)

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
			ToUser:  to,
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
