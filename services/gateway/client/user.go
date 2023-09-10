package client

import (
	"context"
	"log"

	"github.com/yuorei/hackathon/go/user"
	"github.com/yuorei/hackathon/graph/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (c *Client) NewUserClient() {
	// サーバーのアドレスを指定
	userAddress := "localhost:50051"
	// サーバーに接続する
	userConn, err := grpc.Dial(
		userAddress,
		// コネクションでSSL/TLSを使用しない
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// コネクションが確立されるまで待機する(同期処理をする)
		grpc.WithBlock(),
	)

	if err != nil {
		log.Fatal("Connection failed.")
		return
	}

	// gRPCクライアントを生成
	c.userClient = user.NewUserServiceClient(userConn)
}

func (c *Client) CreateUser(input model.CreateUserInput) (*user.CreateUserResponse, error) {
	// リクエストの生成
	request := &user.CreateUserRequest{
		Name:        "test",
		Email:       input.Email,
		Password:    "password",
		Gender:      0,
		Affiliation: 0,
	}

	response, err := c.userClient.CreateUser(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
