package main

import (
	"github.com/IhorBondartsov/stupid-things/little-proj/data-service/config"
	"github.com/IhorBondartsov/stupid-things/little-proj/data-service/transport"
	"golang.org/x/net/context"
)

type DataService struct{
	Port int `json:"port"`
	Host string `json:"host"`
}

func NewDataService(cfg config.Config) *DataService{
	return &DataService{
		Port: 8002,
		Host: "0.0.0.0",
	}
}

func (d *DataService)HealthCheck(context.Context, *transport.HealthCheckRequest) (*transport.HealthCheckResponse, error){
	return &transport.HealthCheckResponse{Reply:"Ok"}, nil
}