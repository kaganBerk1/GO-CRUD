package main

import (
	app "GO-CRUD/GO-CRUD/appilcation"
	"context"
	"fmt"
)

func main() {
	app := app.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("fail to listen the server SADGE: ")
	}
}
