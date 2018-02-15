package main

import (
	"log"
	"golang.org/x/net/context"
	"dispatchlabs/dapos_poc/proto"
	"dispatchlabs/disgo_commons/types"
	"dispatchlabs/disgo/server"
	"time"
)


func init() {

	//TODO: get the instance of the GRPC Service ... then register it
	log.Println("Starting Listning for Delegeate %s\n")

	srv := server.GetGRPCServerInstance()

	//TODO: initialize an actual struct that you will use for the service.
	proto.RegisterTxChannelServer(srv, &ConsensusThingy{} )
	//serviceThingy.TellMeYouAreHere(self)

}

type ConsensusThingy struct {
	//TODO: add fields
}

//Receive the transaction
func (l *ConsensusThingy) ReceiveTransaction(ctx context.Context, in *proto.TransactionWrapper) (*proto.TransactionResponse, error) {
	log.Print("ReceiveTransaction")
	toAddr, err := types.NewAddress()
	if(err != nil) {
		panic("problem creating to address")
	}
	fromAddr, err := types.NewAddress()
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
	//TODO: just here to remove the unused warning ... remove
	log.Printf(string(trans.Value))
	return &proto.TransactionResponse{
	}, nil
}

func (l *ConsensusThingy) BroadcastTransaction(ctx context.Context, in *proto.TransactionWrapper) (*proto.TransactionResponse, error) {
	log.Print("ReceiveTransaction")

	return &proto.TransactionResponse{
	}, nil
}
