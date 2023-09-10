package resolver

import (
	"github.com/yuorei/hackathon/application"
	"github.com/yuorei/hackathon/client"
)

//go:generate go run github.com/99designs/gqlgen generate
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	app    *application.App
	client *client.Client
}

func NewResolver() *Resolver {
	return &Resolver{
		app:    application.NewApplication(),
		client: client.NewClient(),
	}
}
