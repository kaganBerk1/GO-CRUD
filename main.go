package main
import (
	"fmt"
	"context"
	"github.com/kaganBerk1/GO-CRUD/application"
	)

func main(){
	app:= application.New()
	err:= app.Start(context.TODO())
	if err != nil{
		return fmt.Println("fail to listen the server SADGE: ")
	}
}
