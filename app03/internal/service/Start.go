package service

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	pb "kratosL/app03/api"
)

func (s *Service) Start(ctx context.Context, req *empty.Empty) (resp *pb.HelloResp, err error) {
	resp = &pb.HelloResp{
		Content:              "xx",
	}
	fmt.Println("aaa")
	return
}
