package core

import (
	"sync"
	"golang.org/x/net/context"
	"github.com/dispatchlabs/dapos/proto"
	log "github.com/sirupsen/logrus"
	"time"
	"github.com/dispatchlabs/disgo_commons/types"
	"github.com/dispatchlabs/disgo_commons/services"
	"github.com/dispatchlabs/disgo_commons/crypto"
)

var node *Node

// DAPoSService
type DAPoSService struct {
	running bool
}

func GetNode() (*Node) {
	return node
}
// NewDAPoSService
func NewDAPoSService() *DAPoSService {
	daposService := &DAPoSService{
		running: false,
	}
	return daposService
}

func (daposService *DAPoSService) WithGrpc() *DAPoSService {
	proto.RegisterDAPoSGrpcServer(services.GetGrpcServer(), daposService)
	return daposService
}

func (daposService *DAPoSService) WithHttp() *DAPoSService {
	//add the http routing
	return daposService
}

// Name
func (daposService *DAPoSService) Name() string {
	return "DAPoSService"
}

// IsRunning
func (daposService *DAPoSService) IsRunning() bool {
	return daposService.running
}

// Go
func (daposService *DAPoSService) Go(waitGroup *sync.WaitGroup) {
	daposService.running = true
	wa, err := crypto.NewWalletAddress()
	if(err != nil) {
		panic(err)
	}
	//TODO: Temporary code
	node = CreateNodeAndAddToList(wa,"DelegateNode", 0, true)
}

// BroadcastTransaction
//TODO: need to leverage disgover KDHT for this
func (daposService *DAPoSService) BroadcastTransaction(context.Context, *proto.Transaction) (*proto.TransactionResponse, error) {
	log.Printf("BroadcastTransaction")
	return nil, nil
}

//TODO: convert types.Transaction to proto.Transaction and send
func BroadcastTx(*types.Transaction) {
	log.Printf("BroadcastTransaction")
}

// ReceiveTransaction
func (daposService *DAPoSService) ReceiveTransaction(ctx context.Context, in *proto.Transaction) (*proto.TransactionResponse, error) {
	log.Printf("ReceiveTransaction")

	transaction := &types.Transaction{}
	transaction.Type = int(in.Type)
	transaction.Hash = crypto.ToHash(in.Hash)
	copy(transaction.To[:], in.To)
	copy(transaction.From[:], in.From)
	transaction.Value = in.Value
	transaction.Time = time.Unix(in.Time, 0)
	node.ProcessTx(transaction)

	//TODO: just here to remove the unused warning ... don't think we need to have any response
	log.Printf(string(transaction.Value))
	return &proto.TransactionResponse{
	}, nil
}

//TODO: Temporary code for testing the call to consensus.  This allows a client caller to have nodes to use that will be recognized here
func (daposService *DAPoSService) RegisterTestNode(ctx context.Context, inNode *proto.TestNode) (*proto.TransactionResponse, error) {
	log.Printf("RegisterTestNode %s", inNode.Name)
	walletAddress := [20]byte{}
	copy(walletAddress[:], inNode.Address)
	CreateNodeAndAddToList(walletAddress, inNode.Name, inNode.Balance, false)
	return &proto.TransactionResponse{}, nil
}


// CreateTransaction
func (daposService *DAPoSService) CreateTransaction(transaction *types.Transaction, transactions []types.Transaction) (*types.Transaction, error) {
	return nil, nil
}
