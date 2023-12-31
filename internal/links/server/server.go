package handlers

import (
	"context"
	"links/config"
	"links/internal/links"
	"links/internal/links/usecase"
	pb "links/pkg/api/proto"
)

type ShortLinkServer struct {
	pb.UnimplementedShortLinkServer
	LinkService links.LinkService
}

// NewShortLinkServer Create new instance ShortLinkServer
func NewShortLinkServer(ctx context.Context, conf *config.Config) (*ShortLinkServer, error) {
	var (
		server ShortLinkServer
		err    error
	)
	server.LinkService, err = usecase.NewLinkService(ctx, conf)
	if err != nil {
		return nil, err
	}
	return &server, nil
}

// Get Handler for get full link from short link
func (s *ShortLinkServer) Get(ctx context.Context, request *pb.SlRequest) (*pb.SlResponse, error) {
	link := &pb.SlResponse{}
	newLink, err := s.LinkService.GetLink(ctx, request.GetStartLink())
	if err != nil {
		link.ErrorMessage = err.Err.Error()
		link.ErrorCode = err.Code
	}
	link.ResultLink = newLink
	return link, nil
}

// Post Handler for get short link from full link
func (s *ShortLinkServer) Post(ctx context.Context, request *pb.SlRequest) (*pb.SlResponse, error) {
	link := &pb.SlResponse{}
	newLink, err := s.LinkService.PostLink(ctx, request.GetStartLink())
	if err != nil {
		link.ErrorMessage = err.Err.Error()
		link.ErrorCode = err.Code
	}
	link.ResultLink = newLink
	return link, nil
}
