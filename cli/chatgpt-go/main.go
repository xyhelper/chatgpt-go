package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/xyhelper/chatgpt-go"
)

func main() {
	token := uuid.New().String()
	cli := chatgpt.NewClient(
		// chatgpt.WithDebug(true),
		chatgpt.WithTimeout(120*time.Second),
		chatgpt.WithAccessToken(token),
		chatgpt.WithBaseURI("https://freechat.lidong.xin"),
	)
	conversationID := ""
	parentMessage := ""

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入您的消息：")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入失败：", err)
			continue
		}
		stream, err := cli.GetChatStream(message, conversationID, parentMessage)
		if err != nil {
			log.Fatalf("get chat stream failed: %v\n", err)
		}

		var answer string
		for text := range stream.Stream {
			// log.Printf("stream text: %s\n", text.Content)
			print(strings.Replace(text.Content, answer, "", 1))

			answer = text.Content
			conversationID = text.ConversationID
			parentMessage = text.MessageID
		}

		if stream.Err != nil {
			log.Fatalf("stream closed with error: %v\n", stream.Err)
		}
		// 输出换行
		println()
	}
}

// GetStream will return text stream
