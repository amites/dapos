package main

import (
	"net"
	"os"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/dispatchlabs/dapos/core"
)

func main() {

	// Setup log.
	formatter := &log.TextFormatter{
		FullTimestamp: true,
		ForceColors:   false,
	}
	log.SetFormatter(formatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	// Create TCP listener/GRPC.
	listener, error := net.Listen("tcp", ":1975")
	if error != nil {
		log.Fatalf("failed to listen: %v", error)
	}
	grpcServer := grpc.NewServer()

	// Create DAPoSService
	daposService := core.DAPoSService{}
	daposService.Init()
	daposService.RegisterGrpc(grpcServer)
	go daposService.Go(nil)

	// Serve.
	reflection.Register(grpcServer)
	log.WithFields(log.Fields{
		"method": "main",
	}).Info("listening on 1975...")
	if error := grpcServer.Serve(listener); error != nil {
		log.Fatalf("failed to serve: %v", error)
	}
}
