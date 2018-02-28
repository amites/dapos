package dapos

import (
	log "github.com/sirupsen/logrus"
	"sync"
	"github.com/dispatchlabs/dapos/core"
	"github.com/dispatchlabs/disgo_commons/types"
	"github.com/dispatchlabs/disgo_commons/crypto"
	"github.com/dispatchlabs/disgo_commons/constants"
)

var once sync.Once
var node *core.Node

//TODO: need to get the hardcoded values from .disgover config file
func GetNode() (*core.Node) {
	once.Do(func() {
		node = NewNode(crypto.GetWalletAddressBytes(), "test", 100, true)
	})
	return node
}

func NewNode(address [constants.AddressLength]byte, newMember string, initialBalance int64, isDelegate bool) (*core.Node) {
	wallet := types.WalletAccount{
		newMember,
		address,
		newMember,
		initialBalance,
	}

	node := core.Node{
		GenesisBlock:    core.GenesisBlock,
		CurrentBlock:    nil,
		VoteChannel:     make(chan core.Vote),
		Wallet:          wallet,
		IsDelegate:      isDelegate,
		TxFromChainById: map[int64]*types.Transaction{},
		AllVotes:        make(map[int64]*core.Votes),
	}
	node.CurrentBlock = node.GenesisBlock

	core.GetNodes()[address] = &node
	return &node

}


// BroadcastTransaction
//TODO: convert types.Transaction to proto.Transaction and send
func BroadcastTx(*types.Transaction) {
	log.Printf("BroadcastTransaction")
}


// CreateTransaction
func (daposService *DAPoSService) CreateTransaction(transaction *types.Transaction, transactions []types.Transaction) (*types.Transaction, error) {
	log.WithFields(log.Fields{
		"method": "DAPoSService.CreateTransaction",
	}).Info(transaction.MarshalJSON())
	return nil, nil
}

