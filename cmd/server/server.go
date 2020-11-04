package main

import (
	"log"
	"net"

	"github.com/pingkeyun/wechat-rpc-server/internal/app/config"
	"github.com/pingkeyun/wechat-rpc-server/internal/app/models"

	"google.golang.org/grpc"

	"net/http"

	_ "net/http/pprof"

	rpcsrv "github.com/pingkeyun/wechat-rpc-server/internal/app/grpc"
	"github.com/pingkeyun/wechat-rpc-server/pkg/pb"
)

const PORT = "6001"

func main() {
	go http.ListenAndServe("0.0.0.0:6060", nil)

	// db
	config.Setup()
	models.Init()

	grpcServer := grpc.NewServer()

	// accessToken
	pb.RegisterAccessTokenServiceServer(grpcServer, rpcsrv.AccessTokenServer{})

	// UserDevice
	pb.RegisterUserDeviceServiceServer(grpcServer, rpcsrv.UserDeviceServiceServer{})

	// 地暖
	pb.RegisterDinuanServiceServer(grpcServer, rpcsrv.DinuanServiceServer{})

	// User
	pb.RegisterUserServiceServer(grpcServer, rpcsrv.UserServiceServer{})

	// 模板消息
	pb.RegisterMsgTemplateServiceServer(grpcServer, rpcsrv.MessageTemplateServiceServer{})

	// authorizer
	pb.RegisterAuthorizerServiceServer(grpcServer, rpcsrv.AuthorizerServiceServer{})

	// WechatPush
	pb.RegisterWechatPushServiceServer(grpcServer, rpcsrv.WechatPushServiceServer{})

	listen, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	log.Println("服务已启动...")

	grpcServer.Serve(listen)
}
