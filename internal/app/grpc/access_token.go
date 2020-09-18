package grpc

import (
	"context"

	"github.com/pingkeyun/wechat-rpc-server/internal/app"
	"github.com/pingkeyun/wechat-rpc-server/pkg/pb"
)

type AccessTokenServer struct{}

func (a AccessTokenServer) Get(ctx context.Context, r *pb.AccessTokenRequest) (*pb.AccessTokenResponse, error) {
	var token string
	var expired_at uint32

	// server
	server := app.NewServer()
	ac, ok := server.GetItem(r.GetAppId())
	if ok {
		token = ac.AccessToken
		expired_at = uint32(ac.ExpiredAt.Unix())
	}

	return &pb.AccessTokenResponse{
		Token:     token,
		ExpiredAt: expired_at,
	}, nil
}
