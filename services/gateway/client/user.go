package client

import (
	"log"

	"github.com/yuorei/hackathon/hackathon-recult-proto/go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (c *Client) NewUserClient() {
	// サーバーのアドレスを指定
	userAddress := "localhost:8080"
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
