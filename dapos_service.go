package dapos

import (
	"sync"
	"github.com/dispatchlabs/dapos/proto"
	"github.com/dispatchlabs/dapos/core"
	"github.com/dispatchlabs/disgo_commons/services"
	"github.com/dispatchlabs/disgo_commons/crypto"
)

// DAPoSService
type DAPoSService struct {
	running bool
}

// NewDAPoSService
func NewDAPoSService() *DAPoSService {
	daposService := &DAPoSService{
		running: false,
	}
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
	node = core.CreateNodeAndAddToList(wa,"DelegateNode", 0, true)
}

//Protocol functions
func (daposService *DAPoSService) WithGrpc() *DAPoSService {
	proto.RegisterDAPoSGrpcServer(services.GetGrpcServer(), daposService)
	return daposService
}

func (daposService *DAPoSService) WithHttp() *DAPoSService {
	//add the http routing
	return daposService
}

