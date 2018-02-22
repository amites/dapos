package core

import (
	"sync"
	"golang.org/x/net/context"
	"github.com/dispatchlabs/dapos/proto"
	log "github.com/sirupsen/logrus"
	"time"
	"github.com/dispatchlabs/disgo_commons/types"
	"github.com/dispatchlabs/disgo_commons/services"
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
	wa, err := types.NewWalletAddress()
	if(err != nil) {
		panic(err)
	}
	//TODO: Temporary code
	node = CreateNodeAndAddToList(*wa,"DelegateNode", 0, true)
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

	trans := types.Transaction {
		//Hash: in.Hash,
		Type: int(in.Type),
		To: 	*types.GetAddressFromBytes(in.To),
		From:	*types.GetAddressFromBytes(in.From),
		Value: in.Value,
		Time: time.Unix(in.Time, 0),
	}
	node.ProcessTx(&trans)

	//TODO: just here to remove the unused warning ... don't think we need to have any response
	log.Printf(string(trans.Value))
	return &proto.TransactionResponse{
	}, nil
}

//TODO: Temporary code for testing the call to consensus.  This allows a client caller to have nodes to use that will be recognized here
func (daposService *DAPoSService) RegisterTestNode(ctx context.Context, inNode *proto.TestNode) (*proto.TransactionResponse, error) {
	log.Printf("RegisterTestNode %s", inNode.Name)
	wa := types.GetAddressFromBytes(inNode.Address)
	CreateNodeAndAddToList(*wa, inNode.Name, inNode.Balance, false)
	return &proto.TransactionResponse{}, nil
}


// CreateTransaction
func (daposService *DAPoSService) CreateTransaction(transaction *types.Transaction, transactions []types.Transaction) (*types.Transaction, error) {
	return nil, nil
}
