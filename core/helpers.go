package core

import (
	"math/rand"
	"sync"
	"time"
	"github.com/dispatchlabs/disgo_commons/types"
	"github.com/dispatchlabs/disgo_commons/crypto"
	log "github.com/sirupsen/logrus"
)

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

func getDictKeysAsList() []types.WalletAddress {
	keys := make([]types.WalletAddress, 0)
	for k, _ := range getNodes() {
		keys = append(keys, k)
	}

	return keys
}

func getRandomNonDelegateNode(nodeToIgnore *Node) *Node {
	nodes := getNodes()
	nodesNames := getDictKeysAsList()

	var theNode *Node
	for {
		randomNum := GetRandomNumber(len(nodes))
		theNode = getNodes()[nodesNames[randomNum]]

		if nodeToIgnore != nil && nodeToIgnore.Wallet.Id == theNode.Wallet.Id {
			continue
		}

		if !theNode.IsDelegate {
			break
		}
	}

	return theNode
}


var nodes *map[types.WalletAddress]*Node

var once sync.Once

var bytes = make([]byte, 0, 0)

var GenesisBlock = &Block{
	Prev: nil,
	Next: nil,

	Transaction: types.Transaction{
		0,
		crypto.CreateHash(),
		0,
		types.WalletAddress{},
		types.WalletAddress{},
		100,
		time.Now(),
		[]types.WalletAddress{},
	},
}

func CreateNodeAndAddToList(address types.WalletAddress, newMember string, initialBalance int64, isDelegate bool) (*Node) {
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

	getNodes()[address] = &node
	return &node
}

func ElectDelegate(address types.WalletAddress) {
	getNodes()[address].IsDelegate = true
	getNodes()[address].StartVoteCounting()
}

func getNodes() map[types.WalletAddress]*Node {
	once.Do(func() {
		nodes = &map[types.WalletAddress]*Node{}
	})
	return *nodes
}

func getNodeByAddress(address types.WalletAddress) *Node {
	return getNodes()[address]
}
