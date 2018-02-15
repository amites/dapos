package core

import (
"sync"
"google.golang.org/grpc"
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

}

// Go
func (daposService *DAPoSService) Go(waitGroup *sync.WaitGroup) {
	daposService.running = true
}
