package main

import (
	"/home/scot/Desktop/grpc/Server"
	protos "/home/scot/Desktop/grpc/protos/MessageHeader"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	log := log.New()
	gs := grpc.NewServer()
	c := Server.NewMessageHeader(log)
	protos.RegisterEdgeServiceServer(gs, c)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9092))
	if err != nil {
		//log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}

	// listen for requests
	gs.Serve(l)

}
