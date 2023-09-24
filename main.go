package main

import (
	application "GO-CRUD/GO-CRUD/appilcation"
	"context"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	app := application.New()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	err := app.Start(ctx)
	if err != nil {
		fmt.Println("fail to listen the server SADGE: ")
	}
	defer cancel()

}
