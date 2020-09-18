package grpc

import (
	"context"

	"github.com/pingkeyun/wechat-rpc-server/internal/app/models"
	"github.com/pingkeyun/wechat-rpc-server/pkg/pb"
)

type MessageTemplateServiceServer struct {
}

// GetTemplateId 读取模板ID
func (t MessageTemplateServiceServer) GetTemplateId(ctx context.Context, r *pb.GetTemplateIdRequest) (*pb.GetTemplateIdResponse, error) {
	model := models.MessageTemplate{}
	templateId := model.GetTemplateId(r.GetAuthorizerId(), r.GetTemplateName())

	return &pb.GetTemplateIdResponse{
		TemplateId: templateId,
	}, nil
}

// GetTemplateList 读取配置模板列表
func (t MessageTemplateServiceServer) GetTemplateList(ctx context.Context, r *pb.GetTemplateListRequest) (*pb.GetTemplateListResponse, error) {
	model := models.MessageTemplate{}
	items := model.GetTemplateList(r.GetAuthorizerId())

	var result []*pb.TemplateItem
	for name, tid := range items {
		result = append(result, &pb.TemplateItem{
			TemplateName: name,
			TemplateId:   tid,
		})
	}

	return &pb.GetTemplateListResponse{
		Results: result,
	}, nil
}
