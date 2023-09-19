package client

import (
	"log"

	"github.com/yuorei/hackathon/go/image"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (c *Client) NewImageClient() {
	// TCPサーバーのアドレスを指定
	imageAddress := "localhost:50052"
	// サーバーに接続する
	imageConn, err := grpc.Dial(
		imageAddress,
		// コネクションでSSL/TLSを使用しない
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// コネクションが確立されるまで待機する(同期処理をする)
		grpc.WithBlock(),
	)

	if err != nil {
		log.Fatal("Connection failed. err: ", err)
		return
	}

	// gRPCクライアントを生成
	c.imageClient = image.NewUploadImageClient(imageConn)
}
