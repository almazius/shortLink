package handlers

import (
	"context"
	"links/internal/links/usecase"
	pb "links/pkg/api/proto"
)

type ShortLinkServer struct {
	pb.UnimplementedShortLinkServer
}

func (s *ShortLinkServer) Get(ctx context.Context, request *pb.SlRequest) (*pb.SlResponse, error) {
	link := &pb.SlResponse{}
	newLink, err := usecase.GetLink(request.GetStartLink())
	if err != nil {
		link.ErrorMessage = err.Err.Error()
		link.ErrorCode = err.Code
	}
	link.ResultLink = newLink
	return link, nil
}

func (s *ShortLinkServer) Post(ctx context.Context, request *pb.SlRequest) (*pb.SlResponse, error) {
	link := &pb.SlResponse{}
	newLink, err := usecase.PostLink(request.GetStartLink())
	if err != nil {
		link.ErrorMessage = err.Err.Error()
		link.ErrorCode = err.Code
	}
	link.ResultLink = newLink
	return link, nil
}
