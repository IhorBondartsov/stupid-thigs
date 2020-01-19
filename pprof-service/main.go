package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

// Running debug server  on http://localhost:8088/debug/pprof/

func main() {
	done := make(chan error, 2)
	stop := make(chan struct{})
	go func() {
		done <- serveDebug(stop)
	}()
	go func() {
		done <- serveApp(stop)
	}()

	var stopped bool
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Println(err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}

}

func serveDebug(stop <-chan struct{}) error {
	return serve("0.0.0.0:8088", http.DefaultServeMux, stop)
}

func serveApp(stop <-chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hell World!!!")
	})
	return serve("0.0.0.0:8080", mux, stop)
}

func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}
