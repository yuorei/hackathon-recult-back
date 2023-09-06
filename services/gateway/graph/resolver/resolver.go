package resolver

import "github.com/yuorei/hackathon/services/gateway/application"

//go:generate go run github.com/99designs/gqlgen generate
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	app *application.App
}

func NewResolver(app *application.App) *Resolver {
	return &Resolver{app: app}
}
