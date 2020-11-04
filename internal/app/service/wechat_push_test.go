package service_test

import (
	"github.com/pingkeyun/wechat-rpc-server/internal/app/config"
	"testing"
	"github.com/pingkeyun/wechat-rpc-server/internal/app/service"
)

func TestWechatPushPut(t *testing.T) {
	config.Setup()
	t.Run("TestWechatPushPut", func(t *testing.T) {
		for i:= 1; i< 10; i++ {
			ret := service.WechatPushPut("aa")
			if ret != "ok" {
				t.Error("写入消息队列失败")
			}
		}
	})
}
