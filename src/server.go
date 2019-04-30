package src

import (
	"context"
	"fmt"
	. "test/proto"
)

type Server struct {
}

func (s *Server) Get(ctx context.Context, req *GetByNameRequest) (*GetResponse, error) {
	credentialValue, exists := GetValue(req.Name)
	if !exists {
		return &GetResponse{}, fmt.Errorf("the request could not be completed because the credential does not exist or you do not have sufficient authorization")
	}

	return &GetResponse{
		Name: req.Name,
		Type: "value",
		Data: credentialValue,
	}, nil
}
func (s *Server) Set(ctx context.Context, req *SetByNameRequest) (*SetResponse, error) {
	SetValue(req.Name, req.Data)
	return &SetResponse{
		Name: req.Name,
		Type: "value",
		Data: req.Data,
	}, nil
}
func (s *Server) Delete(ctx context.Context, req *DeleteByNameRequest) (*DeleteResponse, error) {
	if DeleteValue(req.Name) {
		return &DeleteResponse{}, nil
	} else {
		return &DeleteResponse{}, fmt.Errorf("the request could not be completed because the credential does not exist or you do not have sufficient authorization")
	}
}

