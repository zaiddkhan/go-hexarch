package grpc

import (
	"go-arch/internal/adapters/framework/left/grpc/pb"
	"go-arch/internal/ports"
	"google.golang.org/grpc"
	"net"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (grpca Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	arithmeticServiceServer := grpca
	grpcServer := grpc.NewServer()
	pb.(grpcServer, arithmeticServiceServer)
}
