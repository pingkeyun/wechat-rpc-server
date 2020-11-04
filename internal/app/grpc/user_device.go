package grpc

import (
	"context"

	"github.com/pingkeyun/wechat-rpc-server/internal/app/models"
	"github.com/pingkeyun/wechat-rpc-server/pkg/pb"
)

type UserDeviceServiceServer struct{}

// Get 根据信息主键获取记录
func (a UserDeviceServiceServer) Get(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {
	user := models.UserDevice{}
	result, err := user.Get(r.GetId())
	if err != nil {
		return nil, err
	}

	item := a.ItemToPbItem(result)

	return &pb.GetResponse{
		Result: item,
	}, nil
}

// GetUserByDeviceId 根据设备ID获取所有用户列表
func (a UserDeviceServiceServer) GetUserByDeviceId(ctx context.Context, r *pb.GetUserByDeviceIdRequest) (*pb.GetUserByDeviceIdResponse, error) {
	user := models.UserDevice{}
	result := user.GetUserByDeviceId(r.GetDeviceId())

	var ret []*pb.Item
	for _, item := range result {
		ret = append(ret, a.ItemToPbItem(item))
	}

	return &pb.GetUserByDeviceIdResponse{
		Results: ret,
	}, nil
}

// GetUserByDeviceNumber 根据设备号获取所有用户列表
func (a UserDeviceServiceServer) GetUserByDeviceNumber(ctx context.Context, r *pb.GetUserByDeviceNumberRequest) (*pb.GetUserByDeviceNumberResponse, error) {
	user := models.UserDevice{}
	result := user.GetUserByDeviceNumber(r.GetDeviceNumber())

	var ret []*pb.Item
	for _, item := range result {
		ret = append(ret, a.ItemToPbItem(item))
	}

	return &pb.GetUserByDeviceNumberResponse{
		Results: ret,
	}, nil
}

// PbItemToStructItem 将model数据类型转换为pb数据类型
func (a UserDeviceServiceServer) ItemToPbItem(item models.UserDevice) *pb.Item {
	// var founder pb.YN
	// founder = 0

	return &pb.Item{
		ID:           item.ID,
		CreatedAt:    item.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    item.UpdatedAt.Format("2006-01-02 15:04:05"),
		Userid:       item.Userid,
		Openid:       item.Openid,
		Title:        item.Title,
		DeviceId:     item.DeviceId,
		DeviceNumber: item.DeviceNumber,
		DeviceType:   uint32(item.DeviceType),
		Founder:      pb.YN(item.Founder),
		AuthorizerId: item.AuthorizerId,
		ProvinceCode: item.ProvinceCode,
		CityCode:     item.CityCode,
		CityName:     item.CityName,
		AreaCode:     item.AreaCode,
		AreaName:     item.AreaName,
	}
}
