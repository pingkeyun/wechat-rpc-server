package models

import (
	"errors"
	"gorm.io/gorm"
)

type Authorizer struct{
	ID        uint64 `gorm:"primaryKey"`
	AuthorizerId uint32
	ComponentAppid string
	AuthorizerAppid string
	NickName string
	HeadImg string
	ServiceTypeInfo uint32
	VerifyTypeInfo uint32
	OriginId string
	PrincipalName string
	Alias string
	BussinessInfo string
	QrcodeUrl string
	AuthorizationInfo string
}

func (Authorizer) TableName() string {
	return "tbl_authorizer_info"
}

func (a Authorizer)GetAuthorizeInfo(id interface{}) (Authorizer, error) {
	row := Authorizer{}
	var resp *gorm.DB

	switch id.(type) {
	case uint32:
		resp = db.Model(&a).Where("authorizer_id = ?", id).Take(&row)
	case string:
		resp = db.Model(&a).Where("authorizer_appid = ?", id).Take(&row)
	default:
		return row, errors.New("dataType invalid")
	}

	if resp.Error != nil {
		return Authorizer{}, resp.Error
	}

	return row, nil
}

func (a Authorizer)GetBaseInfo(authorizerId uint32) (map[string]string, error) {
	type ret struct{
		AuthorizerAppid string
		AuthorizerAppname string
	}
	row := ret{}
	resp:= db.Table("tbl_authorizer_account").
		Select("authorizer_appid", "authorizer_appname").
		Where("authorizer_id = ?", authorizerId).
		Take(&row)
	if resp.Error != nil {
		return nil, resp.Error
	}

	m := map[string]string{
		"AuthorizerAppid": row.AuthorizerAppid,
		"AuthorizerAppname": row.AuthorizerAppname,
	}
	return m, nil
}