package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/IhorBondartsov/stupid-things/little-proj/data-service/transport"
	"github.com/IhorBondartsov/stupid-things/little-proj/static-service/clients"
	"github.com/IhorBondartsov/stupid-things/little-proj/static-service/config"
)

type Service struct {
	Port int
	Host string

	DataServiceAddress string
}

func NewService(config config.Config) *Service {
	return &Service{
		Port: config.Port,
		Host: config.Host,
	}
}

func (s *Service) Run() error {
	handler := http.FileServer(http.Dir("static"))

	http.Handle("/", handler)

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "About Page")
	})

	http.HandleFunc("/ping/dataservice", s.PingDataService)

	return http.ListenAndServe("0.0.0.0:8001", middleware(handler))
}

func (s *Service) PingDataService(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	dataServCli := clients.DataServiceConnect(s.DataServiceAddress)
	resp, err := dataServCli.HealthCheck(ctx, &transport.HealthCheckRequest{Greeting: "Hey"})
	if err != nil {
		log.Print(err)
	}
	_, err = fmt.Fprint(w, resp.GetReply())
	if err != nil {
		log.Print(err)
	}
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.Print(time.Now(), r.URL)
	})
}
