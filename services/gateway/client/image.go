package client

import (
	"context"
	"io"
	"log"

	"github.com/99designs/gqlgen/graphql"
	pb "github.com/yuorei/hackathon/go/image"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	// TCPサーバーのアドレスを指定
	imageServerAddr = "localhost:50052"
	chunkSizeBytes  = 1 * 1024 * 1024   // チャンクサイズ（1MB）
)

func (c *Client) NewImageClient() {
	imageConn, err := grpc.Dial(
		imageServerAddr,
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
	c.imageClient = pb.NewUploadImageClient(imageConn)
}

func (c *Client) UploadImage(imageFile *graphql.Upload) (*pb.UploadResponse, error) {
	conn, err := grpc.Dial(imageServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// gRPCクライアントの初期化
	client := pb.NewUploadImageClient(conn)

	// クライアントストリーミングを使用して画像をサーバーに送信
	stream, err := client.Upload(context.Background())
	if err != nil {
		return nil, err
	}

	// Meta情報を作成して送信
	meta := &pb.UploadRequest_Meta{
		Meta: &pb.Meta{
			Name: imageFile.Filename,
			// Desc: "Image Description",
		},
	}
	firstRequest := &pb.UploadRequest{Value: meta}
	if err := stream.Send(firstRequest); err != nil {
		return nil, err
	}

	buffer := make([]byte, chunkSizeBytes)
	for {
		n, err := imageFile.File.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// データチャンクを作成して送信
		dataChunk := &pb.UploadRequest_Data{
			Data: buffer[:n],
		}
		dataRequest := &pb.UploadRequest{Value: dataChunk}
		if err := stream.Send(dataRequest); err != nil {
			return nil, err
		}
	}

	// ストリームを閉じて応答を受け取る
	response, err := stream.CloseAndRecv()
	if err != nil {
		return nil, err
	}

	log.Printf("Image uploaded successfully. Image URL: %s\n", response.ImageUrl)
	return response, nil
}
