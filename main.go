package main

import (
	"context"
	"fmt"

	"github.com/fabruun/go-authentication/api"
)

func main() {
	app := api.New()

	err := app.Start(context.Background())
	if err != nil {
		fmt.Println("failed to start app", err)
	}
	fmt.Println("Running server on 127.0.0.1:8000")
}
