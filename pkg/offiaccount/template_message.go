package offiaccount

import (
	"errors"
	"github.com/levigross/grequests"
	"strconv"
)

type TemplateSendResult struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg`
	Msgid   int    `json:"msgid"`
}

// Send 发送模板消息
func SendTemplateMessage(accessToken string, json string) (TemplateSendResult, error) {
	url := "https://api.weixin.qq.com/cgi-bin/message/template/send"

	resp, err := grequests.Post(url, &grequests.RequestOptions{
		Params: map[string]string{
			"access_token": accessToken,
		},
		JSON: json,
	})

	if err != nil {
		return TemplateSendResult{}, err
	}

	if !resp.Ok {
		return TemplateSendResult{}, errors.New("httpCode:" + strconv.Itoa(resp.StatusCode))
	}

	result := TemplateSendResult{}
	if err := resp.JSON(&result); err != nil {
		return TemplateSendResult{}, err
	}

	return result, nil
}
