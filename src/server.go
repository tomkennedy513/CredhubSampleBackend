package src

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	"golang.org/x/sys/unix"
	"google.golang.org/grpc"
	"net"
	"os"
	"strings"
	. "test/proto"

)

const (
	netProtocol    = "unix"
)

type Server struct {
	pathToUnixSocket string
	net.Listener
	*grpc.Server
}

func New(pathToUnixSocketFile string) (*Server, error) {
	plugin := new(Server)
	plugin.pathToUnixSocket = pathToUnixSocketFile
	return plugin, nil
}

func (s *Server) Start() {
	s.mustServeRequests()
}

func (s *Server) mustServeRequests() {
	err := s.setupRPCServer()
	if err != nil {
		glog.Fatalf("failed to setup gRPC Server, %v", err)
	}

	err = s.Serve(s.Listener)
	if err != nil {
		glog.Fatalf("failed to serve gRPC, %v", err)
	}
	glog.Infof("Serving gRPC")
}

func (s *Server) setupRPCServer() error {
	if err := s.cleanSockFile(); err != nil {
		return err
	}

	listener, err := net.Listen(netProtocol, s.pathToUnixSocket)
	if err != nil {
		return fmt.Errorf("failed to start listener, error: %v", err)
	}
	s.Listener = listener
	glog.Infof("Listening on unix domain socket: %s", s.pathToUnixSocket)

	s.Server = grpc.NewServer()
	RegisterCredentialServiceServer(s.Server, s)

	return nil
}

func (s *Server) cleanSockFile() error {
	if strings.HasPrefix(s.pathToUnixSocket, "@") {
		return nil
	}

	err := unix.Unlink(s.pathToUnixSocket)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete the socket file, error: %v", err)
	}
	return nil
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

