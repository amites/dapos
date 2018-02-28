package dapos

import (
	"time"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/dispatchlabs/dapos/core"
	"github.com/dispatchlabs/dapos/proto"
	"github.com/dispatchlabs/disgo_commons/types"
	"github.com/dispatchlabs/disgo_commons/crypto"
	"fmt"
	"github.com/dispatchlabs/disgover"
)

// ReceiveTransaction
func (daposService *DAPoSService) ReceiveTransaction(ctx context.Context, inTx *proto.Transaction) (*proto.TransactionResponse, error) {
	log.Printf("ReceiveTransaction")

	transaction := convertToDomain(inTx)
	node.ProcessTx(transaction)

	//TODO: just here to remove the unused warning ... don't think we need to have any response
	return &proto.TransactionResponse{
	}, nil
}

//TODO: need to leverage disgover KDHT for this
func (daposService *DAPoSService) BroadcastTransaction(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	log.Printf("BroadcastTransaction")
	contacts := disgover.GetDisgover().GetContactList()
	for _, contact := range *contacts {
		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", contact.Endpoint.Host, contact.Endpoint.Port), grpc.WithInsecure())
		if err != nil {
			log.Fatalf("cannot dial server: %v", err)
		}
		client := proto.NewDAPoSGrpcClient(conn)

		response, _ := client.BroadcastTransaction(ctx, tx)
		if response == nil {
			log.Error("Could not Broadcast transaction to: ")
		}
	}
	return nil, nil

}


//TODO: Temporary code for testing the call to consensus.  This allows a client caller to have nodes to use that will be recognized here
func (daposService *DAPoSService) RegisterTestNode(ctx context.Context, inNode *proto.TestNode) (*proto.TransactionResponse, error) {
	log.Printf("RegisterTestNode %s", inNode.Name)
	walletAddress := [20]byte{}
	copy(walletAddress[:], inNode.Address)
	core.CreateNodeAndAddToList(walletAddress, inNode.Name, inNode.Balance, false)
	return &proto.TransactionResponse{}, nil
}

func convertToDomain(ptx *proto.Transaction) *types.Transaction {
	transaction := types.Transaction{
		Hash:	crypto.NewHash(ptx.Hash),
		Id: 	ptx.DelegateId,
		From: 	crypto.ToWalletAddress(ptx.From),
		To:		crypto.ToWalletAddress(ptx.To),
		Time:	time.Unix(ptx.Time, 0),
		Value:	ptx.Value,
		Type:	1,
	}
	return &transaction
}

func convertToProto(tx *types.Transaction) *proto.Transaction {
	transaction := proto.Transaction{
		Hash:	crypto.ToBytes(tx.Hash),
		DelegateId: 	tx.Id,
		From: 	crypto.ToWalletAddressBytes(tx.From),
		To:		crypto.ToWalletAddressBytes(tx.To),
		Time:	int64(tx.Time.UnixNano()),
		Value:	tx.Value,
		Type:	1,
	}
	return &transaction
}
