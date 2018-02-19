package core

import (
	"sync"
	"google.golang.org/grpc"
	"context"
	"github.com/dispatchlabs/dapos/proto"
	log "github.com/sirupsen/logrus"
	"github.com/dispatchlabs/disgo_commons/types"
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
	log.WithFields(log.Fields{
		"method": "DAPoSService.Init",
	}).Info("initializing...")
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
	proto.RegisterDAPoSGrpcServer(grpcServer, daposService)
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

// CreateTransaction
func (daposService *DAPoSService) CreateTransaction(transaction *types.Transaction, transactions []types.Transaction) (*types.Transaction, error) {
	return nil, nil
}
