# chatgpt-go

ChatGPT SDK 是用 Golang 编写的，可以用于快速集成 ChatGPT 的聊天机器人功能。

## 快速开始

1. 安装 ChatGPT-Go SDK。

```shell
go get -u github.com/xyhelper/chatgpt-go
```

2. 在独立的对话中进行聊天。

```go
package main

import (
	"log"
	"time"

	chatgpt "github.com/xyhelper/chatgpt-go"
)

func main() {
	token := `random token`

	cli := chatgpt.NewClient(
		chatgpt.WithDebug(true),                            // 开启调试模式
		chatgpt.WithTimeout(120*time.Second),               // 设置超时时间为120秒
		chatgpt.WithAccessToken(token),                     // 设置token
		chatgpt.WithBaseURI("https://freechat.lidong.xin"), // 设置base uri
	)

	// chat in independent conversation
	message := "你好"
	text, err := cli.GetChatText(message)
	if err != nil {
		log.Fatalf("get chat text failed: %v", err)
	}

	log.Printf("q: %s, a: %s\n", message, text.Content)
}

```

3. 在连续的对话中进行聊天。

```go
package main

import (
	"log"
	"time"

	chatgpt "github.com/xyhelper/chatgpt-go"
)

func main() {
	// new chatgpt client
	token := `random token`

	cli := chatgpt.NewClient(
		chatgpt.WithDebug(true),
		chatgpt.WithTimeout(60*time.Second),
		chatgpt.WithAccessToken(token),
		chatgpt.WithBaseURI("https://freechat.lidong.xin"),
	)

	// chat in continuous conversation

	// first message
	message := "对我说你好"
	text, err := cli.GetChatText(message)
	if err != nil {
		log.Fatalf("get chat text failed: %v", err)
	}

	log.Printf("q: %s, a: %s\n", message, text.Content)

	// continue conversation with new message
	conversationID := text.ConversationID
	parentMessage := text.MessageID
	newMessage := "再说一次"

	newText, err := cli.GetChatText(newMessage, conversationID, parentMessage)
	if err != nil {
		log.Fatalf("get chat text failed: %v", err)
	}

	log.Printf("q: %s, a: %s\n", newMessage, newText.Content)
}
```

> 如果你想要在当前对话之外开始一个新的对话，你不需要重置客户端。只需在`GetChatText`方法中移除`conversationID`和`parentMessage`参数，即可获得一个新的文本回复，从而开始一个新的对话。

4. 使用流（stream）获取聊天内容。

```go
package main

import (
	"log"
	"time"

	chatgpt "github.com/xyhelper/chatgpt-go"
)

func main() {
	// new chatgpt client
	token := `random token`

	cli := chatgpt.NewClient(
		chatgpt.WithDebug(true),
		chatgpt.WithTimeout(120*time.Second),
		chatgpt.WithAccessToken(token),
		chatgpt.WithBaseURI("https://freechat.xyhelper.cn"),
	)

	message := "你好"
	stream, err := cli.GetChatStream(message)
	if err != nil {
		log.Fatalf("get chat stream failed: %v\n", err)
	}

	var answer string
	for text := range stream.Stream {
		log.Printf("stream text: %s\n", text.Content)

		answer = text.Content
	}

	if stream.Err != nil {
		log.Fatalf("stream closed with error: %v\n", stream.Err)
	}

	log.Printf("q: %s, a: %s\n", message, answer)
}

```

## DEMO

cli/chatgpt-go 是一个使用 ChatGPT-Go SDK 的示例程序。
可以使用以下命令运行它：

```shell
cd cli/chatgpt-go
go run main.go
```

也可以使用 go install 命令安装它：

```shell
go install github.com/xyhelper/chatgpt-go/cli/chatgpt-go@latest
```

然后运行：

```shell
chatgpt-go
```

## 作品演示

[https://xyhelper.cn](https;//xyhelper.cn)

## 友情链接

- [CoolAdmin](https://cool-js.com) - 一个项目,用 COOL 就够了。AI 编码、物联网、函数云开发等更加智能和全面的功能，让应用开发紧跟时代的步伐，cool 绝不落后！！！
