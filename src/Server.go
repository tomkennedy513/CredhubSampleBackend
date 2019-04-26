package src

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "test/proto"
)
type Server struct {
}

func (s *Server) Get(ctx context.Context, req *GetByNameRequest) (*GetResponse, error) {
	credentialByName := GetValue(req.Name)
	data := credentialByName.Data["data"].(map[string]interface{})
	return &GetResponse{
		Name: req.Name,
		Type: "value",
		Data: []byte(data[req.Name].(string)),
	}, nil
}
func (s *Server) Set(ctx context.Context, req *SetByNameRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (s *Server) Delete(ctx context.Context, req *DeleteByNameRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

