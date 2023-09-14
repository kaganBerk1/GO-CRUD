package main
import (
	"fmt"
	"context"
	"https://github.com/kaganBerk1/GO-CRUD/application"
	)

func main(){
	app:= application.New()
	err:= app.Start()
	if err != nil{
		return fmt.Println("fail to listen the server SADGE: ")
	}
}
