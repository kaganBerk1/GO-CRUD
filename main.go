package main
import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	)

func main(){
	fmt.Println("Hello, World!")

	router:= chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/hello",basicHandler)

	server:=&http.Server{
		Addr: ":3000",
		Handler: router,
	}
	err:= server.ListenAndServe()
	if err != nil{
		fmt.Println("fail to listen the server SADGE", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hellow World From Server"))
}