package examples

import (
	"fmt"
	"time"

	chatgpt "github.com/xyhelper/chatgpt-go"
)

// chatgpt client
var cli *chatgpt.Client

func ExampleNewClient() {
	fmt.Printf("%T", cli)

	// Output: *chatgpt.Client
}

func init() {
	token := `copy-from-cookies`

	cli = chatgpt.NewClient(
		chatgpt.WithDebug(false),
		chatgpt.WithTimeout(60*time.Second),
		chatgpt.WithAccessToken(token),
	)
}
