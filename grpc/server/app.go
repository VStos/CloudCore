package server

import (
	"context"
	"time"

	grpc_v1 "github.com/VStos/CloudCore/grpc"
	"google.golang.org/grpc"
)

type Options struct {
	TcpPort int
	TimeOut time.Duration
}

func (s *Options) Init() {
	s.TcpPort = 443
	s.TimeOut = time.Hour
}

type serverAPI struct {
	grpc_v1.UnimplementedAuthServer // Хитрая штука, о ней ниже
	auth                            Auth
}

// Тот самый интерфейс, котрый мы передавали в grpcApp
type Auth interface {
	Login(ctx context.Context, email string, password string) (token string, err error)
	RegisterNewUser(ctx context.Context, email string, password string) (user_id int64, err error)
}

func Register(gRPCServer *grpc.Server, auth Auth) {
	grpc_v1.RegisterAuthServer(gRPCServer, &serverAPI{auth: auth})
}

func (s *serverAPI) Login(ctx context.Context, in *grpc_v1.LoginRequest) (*grpc_v1.LoginResponse, error) {
	token, _ := s.auth.Login(ctx, in.GetEmail(), in.GetPassword())
	return &grpc_v1.LoginResponse{Token: token}, nil
}

func (s *serverAPI) Register(ctx context.Context, in *grpc_v1.RegisterRequest) (*grpc_v1.RegisterResponse, error) {
	user_id, err := s.auth.RegisterNewUser(ctx, in.GetEmail(), in.GetPassword())
	return user_id, err
}
