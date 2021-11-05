package main

import (
	"fmt"
	"os"

	call2faSDK "github.com/rikkicom/call2fa-go-sdk"
)

func main() {
	// If you like, enable debug of HTTP requests, 0 to disable
	_ = os.Setenv("GOREQUEST_DEBUG", "1")

	// Configure the client
	cfg := call2faSDK.Config{
		Login:    "****",
		Password: "****",
	}

	// Create the Call2FA client
	c, err := call2faSDK.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// Configure variables
	code := "w2353"
	lang := "uk" // or "ru"
	phoneNumber := "+380631010121"

	// Do the request to start the call
	response, err := c.DictateCodeCall(phoneNumber, code, lang)
	if err != nil {
		panic(err)
	}

	fmt.Println("Call ID:", response.CallID)
}
