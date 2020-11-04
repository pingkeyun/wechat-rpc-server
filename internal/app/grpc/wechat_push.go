package grpc

import (
	"context"

	"github.com/pingkeyun/wechat-rpc-server/pkg/pb"
	"github.com/pingkeyun/wechat-rpc-server/internal/app/service"
)

type WechatPushServiceServer struct {
}

// Put 写放消息队列
func (w WechatPushServiceServer)Put(ctx context.Context, r *pb.WechatPushRequest) (*pb.WechatPushResponse, error){
	// 连接rabbitmq，将消息写入队列，由推送服务消费
	ret := service.WechatPushPut(r.GetJsonstr())

	return &pb.WechatPushResponse{
		Ret: ret,
	}, nil
}

