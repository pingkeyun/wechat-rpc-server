package app

import (
	"strconv"
	"sync"
	"time"

	"github.com/prometheus/common/log"
)

// AccessTokenItem 令牌结构
type AccessTokenItem struct {
	AccessToken string
	ExpiredAt   time.Time
}

// IsExpired 令牌过期
func (ac AccessTokenItem) IsExpired() bool {
	return time.Now().After(ac.ExpiredAt)
}

var singletonServer *Server

// Server AccessToken专用服务
type Server struct {
	mu   sync.Mutex
	data map[string]AccessTokenItem
}

var once sync.Once

// NewServer 服务
func NewServer() *Server {
	once.Do(func() {
		singletonServer = &Server{
			data: make(map[string]AccessTokenItem),
		}
	})

	return singletonServer
}

// GetItem 获取token项
func (s *Server) GetItem(appID string) (AccessTokenItem, bool) {
	s.mu.Lock()

	ac, ok := s.data[appID]
	if ok && !ac.IsExpired() {
		s.mu.Unlock()
		return ac, ok
	}

	if err := s.refresh(appID); err != nil {
		log.Error(err, appID)
		ac, ok = AccessTokenItem{
			AccessToken: "",
			ExpiredAt:   time.Now(),
		}, true
	} else {
		ac, ok = s.data[appID]
	}

	s.mu.Unlock()

	return ac, ok
}

// refresh 刷新token
func (s *Server) refresh(appID string) error {
	result, err := GetAccessToken(appID)
	if err != nil {
		return err
	}

	t, _ := strconv.ParseInt(result.Data["expired_at"].(string), 10, 32)
	s.data[appID] = AccessTokenItem{
		AccessToken: result.Data["access_token"].(string),
		ExpiredAt:   time.Unix(t, 0),
	}

	return nil
}

// GetAccessToken 远程获取api接口
func (s *Server) GetAccessToken(appID string) string {
	ac, ok := s.GetItem(appID)
	if ok {
		return ac.AccessToken
	}

	return ""
}
