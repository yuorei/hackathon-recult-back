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
	// TCPサーバーのアドレスを指定
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
		log.Fatal("Connection failed. err: ", err)
		return
	}

	// gRPCクライアントを生成
	c.userClient = user.NewUserServiceClient(userConn)
}

func (c *Client) CreateUser(input model.CreateUserInput) (*user.CreateUserResponse, error) {
	var gender *user.Gender
	if input.Gender != nil {
		switch input.Gender.String() {
		case model.GenderMan.String():
			gender = user.Gender_MAN.Enum()
		case model.GenderWoman.String():
			gender = user.Gender_WOMAN.Enum()
		case model.GenderOther.String():
			gender = user.Gender_GENDER_OTHER.Enum()
		}
	}

	var affiliation *user.Affiliation
	if input.Affiliation != nil {
		switch input.Affiliation.String() {
		case model.AffiliationStudent.String():
			affiliation = user.Affiliation_STUDENT.Enum()
		case model.AffiliationOther.String():
			affiliation = user.Affiliation_AFFILIATION_OTHER.Enum()
		}
	}

	// リクエストの生成
	request := &user.CreateUserRequest{
		Name:  input.Name,
		Email: input.Email,
		// TODO: パスワードのハッシュ化
		Password:    input.Password,
		Gender:      *gender,
		Affiliation: *affiliation,
	}

	response, err := c.userClient.CreateUser(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
