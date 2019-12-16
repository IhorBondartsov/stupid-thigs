package main

import (
	"fmt"
	"github.com/IhorBondartsov/stupid-things/little-proj/static-service/config"
	"log"
	"os"
)

func main() {
	log.Print("Server is starting")

	dataServiceAddress := os.Getenv("DATA_SERVICE_URL")
	fmt.Println("DATA_SERVICE_URL:", dataServiceAddress)

	c := config.ReadFlags()

	s := NewService(c)
	s.DataServiceAddress = dataServiceAddress
	err := s.Run()
	if err!= nil{
	panic(err)
	}
}
