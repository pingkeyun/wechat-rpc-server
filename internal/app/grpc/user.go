package grpc

import (
	"context"

	"github.com/pingkeyun/wechat-rpc-server/internal/app/models"
	"github.com/pingkeyun/wechat-rpc-server/pkg/pb"
)

type UserServiceServer struct {
}

// GetUser 读取用户基本信息
func (u UserServiceServer) GetUser(ctx context.Context, r *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	model := models.User{}
	user, err := model.GetUser(r.GetUserId())
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		Result: u.ItemToPbItem(user),
	}, nil
}

// PbItemToStructItem 将model数据类型转换为pb数据类型
func (u UserServiceServer) ItemToPbItem(item models.User) *pb.User {
	return &pb.User{
		Userid:        item.Userid,
		Nickname:      item.Nickname,
		Realname:      item.Realname,
		Openid:        item.Openid,
		Headimgurl:    item.Headimgurl,
		Mobile:        item.Mobile,
		Province:      item.Province,
		City:          item.City,
		Sex:           item.Sex,
		SubscribeTime: item.SubscribeTime,
		Subscribe:     item.Subscribe,
		AuthorizerId:  item.AuthorizerId,
		GmtModified:   item.GmtModified.Format("2006-01-02 15:04:05"),
	}
}
