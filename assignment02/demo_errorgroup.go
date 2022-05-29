package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// https://pkg.go.dev/golang.org/x/sync/errgroup

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("world"))
	})

	serverOut := make(chan struct{})
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		serverOut <- struct{}{}
	})

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	// g1, if g1退出, ctx cancel ，g2, g3 都会退出
	g.Go(func() error {
		err := server.ListenAndServe()
		if err != nil {
			log.Println("g1 error", err.Error())
		}
		return err
	})

	// if g2退出  server.Shutdown ->g1退出
	// ctx cancel, g3退出
	g.Go(func() error {
		select {
		case <-ctx.Done():
			log.Println("errgroup exit.")
		case <-serverOut:
			log.Println("server will shut down.")
		}

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		log.Println("shutting down")
		return server.Shutdown(timeoutCtx)
	})

	// g3 watch SIGTERM
	// g3 接收到quit 信号，ctx cancel，g2,g1退出
	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			log.Println("g3 ctx execute cancel...")
			return ctx.Err()
		case sig := <-quit:
			return fmt.Errorf("g3 get os signal: %v", sig)
		}
	})
	// 所有的Go方法启动的协程结束后执行
	fmt.Printf("errgroup exiting: %+v\n", g.Wait())
}
