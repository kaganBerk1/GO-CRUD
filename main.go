package main

import (
	application "GO-CRUD/GO-CRUD/appilcation"
	"context"
	"fmt"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("fail to listen the server SADGE: ")
	}

}
