package client

import (
	"github.com/yuorei/hackathon/go/user"
)

type Client struct {
	userClient user.UserServiceClient
}

func NewClient() *Client {
	client := &Client{}
	client.NewUserClient()
	return client
}
