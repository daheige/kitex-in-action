package rpc

import (
	"context"
	"log"

	"go.uber.org/zap"

	"kitex-example/internal/infras/logger"
	"kitex-example/internal/pb"
)

// GreeterImpl implements the last service interface defined in the IDL.
type GreeterImpl struct{}

// Hello implements the GreeterImpl interface.
func (s *GreeterImpl) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("request msg: ", req.Msg)

	logger.Info(ctx, "hello world", zap.String("foo", "bar"))
	resp := &pb.HelloReply{
		Msg: req.Msg,
	}

	// TODO: Your code here...
	return resp, nil
}
