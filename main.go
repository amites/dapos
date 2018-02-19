package main

import (
	"net"
	"os"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/dispatchlabs/dapos/core"
	"github.com/dispatchlabs/dapos/proto"

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
	log.WithFields(log.Fields{
		"method": "main",
	}).Info("listening on 1975...")
	reflection.Register(grpcServer)
	if error := grpcServer.Serve(listener); error != nil {
		log.Fatalf("failed to serve: %v", error)
	}

	conn, err := grpc.Dial("localhost:1975", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial server: %v", err)
	}
	client := proto.NewDAPoSGrpcClient(conn)
	client.BroadcastTransaction(context.Background(), &proto.Transaction{})

}
