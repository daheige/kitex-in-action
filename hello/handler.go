package main

import (
	"context"

	"hello/kitex_gen/pb"
)

// GreeterImpl implements the last service interface defined in the IDL.
type GreeterImpl struct{}

// Hello implements the GreeterImpl interface.
func (s *GreeterImpl) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := &pb.HelloReply{
		Msg: "hello," + req.Msg,
		Id:  1,
	}

	return resp, nil
}
