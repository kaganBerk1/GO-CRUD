package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func New() *App {
	app := &App{
		router: loadRoutes(),
		rdb:    redis.NewClient(&redis.Options{}),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}
	ch := make(chan error, 1)
	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("fail to 	conenct to redis: %w", err)
	}
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("fail to 	starting server %w", err)
		}
		close(ch)
	}()
	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}
	}()

	fmt.Println("3000 listening...")
	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}

}
