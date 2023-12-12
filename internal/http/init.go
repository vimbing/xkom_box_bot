package http

import (
	"github.com/vimbing/fhttp/cookiejar"
	tls "github.com/vimbing/utls"
)

func Init(opts ...Options) (*Client, error) {
	client := Client{}

	if len(opts) > 0 {
		client.Hello = opts[0].Hello
	} else {
		client.Hello = tls.HelloIOS_12_1
	}

	err := client.InitClient()

	client.Client.Jar, _ = cookiejar.New(nil)

	return &client, err
}
