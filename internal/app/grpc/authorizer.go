package grpc

import (
	"context"

	"github.com/pingkeyun/wechat-rpc-server/internal/app/models"
	"github.com/pingkeyun/wechat-rpc-server/pkg/pb"
)

type AuthorizerServiceServer struct{}

// 基本信息
func (a AuthorizerServiceServer) GetBaseInfo(ctx context.Context, r *pb.GetBaseInfoRequest) (*pb.GetBaseInfoResponse, error) {
	model := models.Authorizer{}
	m, err := model.GetBaseInfo(r.GetAuthorizerId())
	if err != nil {
		return nil, err
	}

	return &pb.GetBaseInfoResponse{
		AuthorizerAppid:   m["AuthorizerAppid"],
		AuthorizerAppname: m["AuthorizerAppname"],
	}, nil
}

// 授权信息
func (a AuthorizerServiceServer) GetAuthorizeInfo(ctx context.Context, r *pb.GetAuthorizeInfoRequest) (*pb.GetAuthorizeInfoResponse, error) {
	model := models.Authorizer{}
	var row models.Authorizer
	var err error
	if r.AuthorizerId > 0 {
		row, err = model.GetAuthorizeInfo(r.AuthorizerId)
	} else {
		row, err = model.GetAuthorizeInfo(r.AuthorizerAppid)
	}
	if err != nil {
		return nil, err
	}

	return &pb.GetAuthorizeInfoResponse{
		AuthorizerId:    row.AuthorizerId,
		AuthorizerAppid: row.AuthorizerAppid,
		NickName:        row.NickName,
		HeadImg:         row.HeadImg,
		ServiceTypeInfo: row.ServiceTypeInfo,
		VerifyTypeInfo:  row.VerifyTypeInfo,
		OriginId:        row.OriginId,
		PrincipalName:   row.PrincipalName,
		Alias:           row.Alias,
		BussinessInfo:   row.BussinessInfo,
		QrcodeUrl:       row.QrcodeUrl,
	}, nil
}
