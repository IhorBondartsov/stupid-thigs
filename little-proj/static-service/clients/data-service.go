package clients

import (
	"github.com/IhorBondartsov/stupid-things/little-proj/data-service/transport"
	"google.golang.org/grpc"
	"log"
)

func DataServiceConnect(address string) transport.DataServiceClient{
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return transport.NewDataServiceClient(conn)
}