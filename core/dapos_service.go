package core

import (
	"sync"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"github.com/dispatchlabs/dapos/proto"
	log "github.com/sirupsen/logrus"
	"time"
	"github.com/dispatchlabs/disgo_commons/types"
	service "github.com/dispatchlabs/disgo_commons/service"
)

var node *Node
// DAPoSService
type DAPoSService struct {
	running bool
}

// NewDAPoSService
func NewDAPoSService() *DAPoSService {

	daposService := &DAPoSService{
		running: false,
	}
	daposService.Init()
	daposService.RegisterGrpc(service.GetGrpcServer())

	return daposService
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
	wa, err := types.NewWalletAddress()
	if(err != nil) {
		panic(err)
	}
	node = CreateNodeAndAddToList(*wa, "Bob", 100)
}

// BroadcastTransaction
func (daposService *DAPoSService) BroadcastTransaction(context.Context, *proto.Transaction) (*proto.TransactionResponse, error) {
	log.Printf("BroadcastTransaction")
	return nil, nil
}

func BroadcastTx(*types.Transaction) {
	log.Printf("BroadcastTransaction")
	//TODO: convert types.Transaction to proto.Transaction and send
}

// ReceiveTransaction
func (daposService *DAPoSService) ReceiveTransaction(ctx context.Context, in *proto.Transaction) (*proto.TransactionResponse, error) {
	log.Printf("ReceiveTransaction")
	toAddr, err := types.NewWalletAddress()
	if(err != nil) {
		panic("problem creating to address")
	}
	fromAddr, err := types.NewWalletAddress()
	if(err != nil) {
		panic("problem creating to address")
	}
	trans := types.Transaction {
		//Hash: in.Hash,
		Type: int(in.Type),
		To: *toAddr,
		From: *fromAddr,
		Value: in.Value,
		Time: time.Unix(in.Time, 0),
	}
	node.ProcessTx(&trans)

	//TODO: just here to remove the unused warning ... remove
	log.Printf(string(trans.Value))
	return &proto.TransactionResponse{
	}, nil
}

// CreateTransaction
func (daposService *DAPoSService) CreateTransaction(transaction *types.Transaction, transactions []types.Transaction) (*types.Transaction, error) {
	return nil, nil
}
