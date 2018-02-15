package core

import (
	"sync"
	"google.golang.org/grpc"
	"context"
	"github.com/dispatchlabs/dapos/grpc"
)

// DAPoSService
type DAPoSService struct {
	running bool
}

// NewDAPoSService
func NewDAPoSService() *DAPoSService {
	return &DAPoSService{
		running: false,
	}
}

// Init
func (daposService *DAPoSService) Init() {
}

// Name
func (daposService *DAPoSService) Name() string {
	return "DAPoSService"
}

// IsRunning
func (daposService *DAPoSService) IsRunning() bool {
	return daposService.running
}

// Register
func (daposService *DAPoSService) RegisterGrpc(grpcServer *grpc.Server) {
	proto.RegisterDAPoSServiceServer(grpcServer, daposService)
}

// Go
func (daposService *DAPoSService) Go(waitGroup *sync.WaitGroup) {
	daposService.running = true
}

// BroadcastTransaction
func (daposService *DAPoSService) BroadcastTransaction(context.Context, *proto.Transaction) (*proto.TransactionResponse, error) {
	return nil, nil
}

// ReceiveTransaction
func (daposService *DAPoSService) ReceiveTransaction(context.Context, *proto.Transaction) (*proto.TransactionResponse, error) {
	return nil, nil
}
