package main

import (
	"context"
	"fmt"
	"github.com/davidhong101/go-study-lib/log"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
)

func main() {

	var eg errgroup.Group
	httpErrChan := make(chan int)

	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/panic", panicHandler)
	server := &http.Server{Addr: ":8080", Handler: mux}

	eg.Go(func() error {
		// http server

		err := server.ListenAndServe()
		if err != nil {
			log.ERROR("http.ListenAndServe: %v", err)
			httpErrChan <- 1
			return errors.Wrap(err, "http.ListenAndServe fail. service: 8080")
		}

		return nil
	})

	eg.Go(func() error {
		// os signal handler

		c := make(chan os.Signal)
		signal.Notify(c)

		for {
			select {
			case <-c:
				server.Shutdown(context.TODO())
			case <-httpErrChan:
				return nil
			}
		}

		return nil
	})

	if err := eg.Wait(); err != nil {
		log.ERROR("error happened. err: %v", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func panicHandler(w http.ResponseWriter, r *http.Request) {
	panic("panic http")
}
