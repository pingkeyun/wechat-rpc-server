package app

import (
	"errors"
	"log"
	"strconv"

	"github.com/levigross/grequests"
)

const URL = "https://pt.pingkeyun.com/api/authorizer/accesstoken"

type Result struct {
	Errcode int                    `json:'errcode'`
	Errmsg  string                 `json:'errmsg'`
	Data    map[string]interface{} `json:'data'`
}

func GetAccessToken(appId string) (*Result, error) {
	requestOption := &grequests.RequestOptions{
		Data: map[string]string{
			"appid": appId,
		},
	}
	resp, err := grequests.Post(URL, requestOption)
	if err != nil {
		return &Result{}, err
	}
	defer resp.Close()

	// 200
	if !resp.Ok {
		return &Result{}, errors.New("响应状态码错误:" + strconv.Itoa(resp.StatusCode))
	}

	r := &Result{
		Data: make(map[string]interface{}),
	}
	err = resp.JSON(r)
	if err != nil {
		log.Fatal(err)
	}
	if r.Errcode == 1 {
		return r, errors.New(r.Errmsg)
	}

	return r, nil
}
