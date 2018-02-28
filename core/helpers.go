package core

import (
	"math/rand"
	"time"
	"github.com/dispatchlabs/disgo_commons/types"
	log "github.com/sirupsen/logrus"
	"github.com/dispatchlabs/disgo_commons/constants"
	"sync"
)

var once sync.Once

func logSeparator() {
	log.Info("~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~")
}
func prefixLinesWith(lines []string, prefix string) []string {
	var prefixedLines = []string{}

	for index, line := range lines {
		if index == 0 {
			prefixedLines = append(prefixedLines, line)
		} else {
			prefixedLines = append(prefixedLines, prefix+line)
		}
	}

	return prefixedLines
}

func GetRandomNumber(boundary int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(boundary)
}

func getDictKeysAsList() [][constants.AddressLength]byte {
	keys := make([][constants.AddressLength]byte, 0)
	for k, _ := range GetNodes() {
		keys = append(keys, k)
	}

	return keys
}

func getRandomNonDelegateNode(nodeToIgnore *Node) *Node {
	nodes := GetNodes()
	nodesNames := getDictKeysAsList()

	var theNode *Node
	for {
		randomNum := GetRandomNumber(len(nodes))
		theNode = GetNodes()[nodesNames[randomNum]]

		if nodeToIgnore != nil && nodeToIgnore.Wallet.Id == theNode.Wallet.Id {
			continue
		}

		if !theNode.IsDelegate {
			break
		}
	}

	return theNode
}


var nodes *map[[constants.AddressLength]byte]*Node


var bytes = make([]byte, 0, 0)

var GenesisBlock = &Block{
	Prev: nil,
	Next: nil,
	Transaction: types.Transaction{
		//0,
		//crypto.NewHash(),
		//0,
		//[constants.AddressLength]byte{},
		//[constants.AddressLength]byte{},
		//100,
		//time.Now(),
		//[][constants.AddressLength]byte{},
	},
}

func CreateNodeAndAddToList(address [constants.AddressLength]byte, newMember string, initialBalance int64, isDelegate bool) (*Node) {
	wallet := types.WalletAccount{
		newMember,
		address,
		newMember,
		initialBalance,
	}

	node := Node{
		GenesisBlock:    GenesisBlock,
		CurrentBlock:    nil,
		VoteChannel:     make(chan Vote),
		Wallet:          wallet,
		IsDelegate:      isDelegate,
		TxFromChainById: map[int64]*types.Transaction{},
		AllVotes:        make(map[int64]*Votes),
	}
	node.CurrentBlock = node.GenesisBlock

	GetNodes()[address] = &node
	return &node
}

func ElectDelegate(address [constants.AddressLength]byte) {
	GetNodes()[address].IsDelegate = true
	GetNodes()[address].StartVoteCounting()
}

func GetNodes() map[[constants.AddressLength]byte]*Node {
	once.Do(func() {
		nodes = &map[[constants.AddressLength]byte]*Node{}
	})
	return *nodes
}

func getNodeByAddress(address [constants.AddressLength]byte) *Node {
	return GetNodes()[address]
}
