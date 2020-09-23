package models

import "sync"

type MessageTemplate struct {
	ID              uint64 `gorm:"primaryKey"`
	AuthorizerID    uint32
	TemplateName    string
	TemplateId      string
	State           uint8
	TemplateIdShort string
}

func (MessageTemplate) TableName() string {
	return "tbl_authorizer_msg_template"
}

func (t MessageTemplate) GetTemplateId(authorizerId uint32, templateName string) string {
	var ret struct {
		TemplateId string
	}
	db.Model(&t).Select("template_id").Where("authorizer_id = ? AND template_name = ? AND state = 1", authorizerId, templateName).Scan(&ret)

	return ret.TemplateId
}

type MessageTemplateResultItem struct {
	TemplateName, TemplateId string
}

func (t MessageTemplate) GetTemplateList(authorizerId uint32) map[string]string {
	rows, _ := db.Model(&t).Select("template_name", "template_id").Where("authorizer_id = ?", authorizerId).Rows()

	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	var result = make(map[string]string)
	for rows.Next() {
		var item MessageTemplateResultItem
		db.ScanRows(rows, &item)

		result[item.TemplateName] = item.TemplateId
	}

	return result
}
