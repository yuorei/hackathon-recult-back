package application

import "github.com/yuorei/hackathon/client"

type App struct {
	c *client.Client
}

func NewApplication() *App {

	return &App{
		c: client.NewClient(),
	}
}
