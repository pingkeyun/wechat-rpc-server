package app

import (
	"strconv"
	"sync"
	"time"
)

type AccessTokenItem struct {
	AccessToken string
	ExpiredAt   time.Time
}

// 令牌过期
func (ac AccessTokenItem) IsExpired() bool {
	return time.Now().After(ac.ExpiredAt)
}

var singletonServer *Server

type Server struct {
	mu   sync.Mutex
	data map[string]AccessTokenItem
}

var once sync.Once

func NewServer() *Server {
	once.Do(func() {
		singletonServer = &Server{
			data: make(map[string]AccessTokenItem),
		}
	})

	return singletonServer
}

// Get 获取token项
func (s *Server) GetItem(appId string) (AccessTokenItem, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ac, ok := s.data[appId]
	if ok && !ac.IsExpired() {
		return ac, ok
	}

	s.refresh(appId)
	ac, ok = s.data[appId]

	return ac, ok
}

// refresh 刷新token
func (s *Server) refresh(appId string) {
	result, err := GetAccessToken(appId)
	if err != nil {
		return
	}

	t, _ := strconv.ParseInt(result.Data["expired_at"].(string), 10, 32)
	s.data[appId] = AccessTokenItem{
		AccessToken: result.Data["access_token"].(string),
		ExpiredAt:   time.Unix(t, 0),
	}
}

// GetAccessToken 远程获取api接口
func (s *Server) GetAccessToken(appId string) string {
	ac, ok := s.GetItem(appId)
	if ok {
		return ac.AccessToken
	}

	return ""
}
