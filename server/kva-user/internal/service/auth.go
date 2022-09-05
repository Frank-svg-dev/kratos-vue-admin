package service

import (
	"context"

	pb "kva-user/api/kva-user"
)

type AuthService struct {
	pb.UnimplementedAuthServer
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) CreateAuth(ctx context.Context, req *pb.UserLoginReq) (*pb.UserLoginRes, error) {
	return &pb.UserLoginRes{}, nil
}
