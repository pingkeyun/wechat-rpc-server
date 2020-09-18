package offiaccount

import (
	"encoding/json"
	"errors"
	"sync"
)

type Item struct {
	Value string `json:"value""`
	Color string `json:"color"`
}

// NewItem 创建消息模板数据项
func NewItem(value, color string) Item {
	if color == "" {
		color = "#000000"
	}
	return Item{
		Value: value,
		Color: color,
	}
}

type TemplateMessage struct {
	mu          sync.Mutex
	ToUser      string            `json:"touser"`
	TemplateId  string            `json:"template_id"`
	Url         string            `json:"url,omitempty"`
	Miniprogram map[string]string `json:"miniprogram,omitempty"'`
	Data        map[string]Item   `json:"data"`
}

func NewTemplateMessage(templateID string) *TemplateMessage {
	return &TemplateMessage{
		TemplateId: templateID,
		Data:       make(map[string]Item),
	}
}

// SetUser 设置用户OPENID
func (t *TemplateMessage) SetUser(openID string) *TemplateMessage {
	t.ToUser = openID
	return t
}

// SetUrl 设置跳转URL
func (t *TemplateMessage) SetUrl(url string) *TemplateMessage {
	t.Url = url
	return t
}

// SetMiniprogram 设置小程序
func (t *TemplateMessage) SetMiniprogram(appid, pagepath string) *TemplateMessage {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.Miniprogram["appid"] = appid
	t.Miniprogram["pagepath"] = pagepath

	return t
}

// addDataItem 添加消息内容
func (t *TemplateMessage) AddItem(key string, value Item) *TemplateMessage {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.Data[key] = value

	return t
}

// Json 返回模板消息JSON数据格式
func (t *TemplateMessage) Json() ([]byte, error) {
	if len(t.Data) < 0 || len(t.Data) > 5 {
		return nil, errors.New("模板消息个数不合法")
	}
	return json.Marshal(t)
}
