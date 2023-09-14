package application
import (
	"net/http"
	"fmt"
	"context"
)
type App struct{
	router http.Handler
}

func New() *App{
	app:=&App{
		router:loadRoutes(),
	}
	return app
}

func (a *App) start(ctx context.Context) error {
	server:=&http.Server{
		Addr: ":3000",
		Handler: a.router,
	}
	
	err:= server.ListenAndServe()
	if err != nil{
		return fmt.Errorf("fail to listen the server SADGE: %w", err)
	}
	return nil
}