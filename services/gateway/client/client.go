package client

import (
	"github.com/yuorei/hackathon/go/image"
	"github.com/yuorei/hackathon/go/user"
)

type Client struct {
	userClient  user.UserServiceClient
	imageClient image.UploadImageClient
}

func NewClient() *Client {
	client := &Client{}
	client.NewUserClient()
	client.NewImageClient()
	return client
}
