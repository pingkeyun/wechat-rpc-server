package models

import "time"

type User struct {
	Userid        uint64 `gorm:"primaryKey"`
	Nickname      string
	Realname      string
	Openid        string
	Headimgurl    string
	Mobile        string
	Province      string
	City          string
	Sex           uint32
	SubscribeTime uint64
	Subscribe     uint32
	AuthorizerId  uint32
	GmtModified   time.Time
	Unionid       string
}

func (User) TableName() string {
	return "tbl_users"
}

// GetUser 读取用户信息
func (u User) GetUser(userId uint64) (User, error) {
	result := db.Take(&u, userId)

	return u, result.Error
}
