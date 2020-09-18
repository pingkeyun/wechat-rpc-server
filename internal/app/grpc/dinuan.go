package grpc

import (
	"context"

	"github.com/pingkeyun/wechat-rpc-server/internal/app/models"
	"github.com/pingkeyun/wechat-rpc-server/pkg/pb"
)

type DinuanServiceServer struct {
}

// getDeviceIdByStartDate 根据日期获取所有设备ID
func (d DinuanServiceServer) GetDeviceIdByStartDate(ctx context.Context, r *pb.GetDeviceIdByDateRequest) (*pb.GetDeviceIdByDateResponse, error) {
	model := models.Dinuan{}
	result := model.GetDeviceIdByStartDate(r.GetDate())

	return &pb.GetDeviceIdByDateResponse{
		DeviceId: result,
	}, nil
}

// getDeviceIdByEndDate 根据日期获取所有设备ID
func (d DinuanServiceServer) GetDeviceIdByEndDate(ctx context.Context, r *pb.GetDeviceIdByDateRequest) (*pb.GetDeviceIdByDateResponse, error) {
	model := models.Dinuan{}
	result := model.GetDeviceIdByEndDate(r.GetDate())

	return &pb.GetDeviceIdByDateResponse{
		DeviceId: result,
	}, nil
}
