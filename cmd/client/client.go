package main

import (
	"context"
	"fmt"
	"log"

	"github.com/pingkeyun/wechat-rpc-server/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:6001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ctx := context.Background()
	user := pb.NewUserDeviceServiceClient(conn)
	request := pb.GetUserByDeviceNumberRequest{
		DeviceNumber: "730000F514AF",
	}
	resp, err := user.GetUserByDeviceNumber(ctx, &request)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(len(resp.Results))
	for _, v := range resp.Results {
		fmt.Println(v)
		fmt.Println(v.GetDeviceId(), v.GetDeviceNumber(), v.GetDeviceType(), v.GetOpenid(), v.GetFounder())
		// fmt.Printf("%d:, %#v", i, v)
	}

}
