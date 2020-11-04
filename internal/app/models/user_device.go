package models

type UserDevice struct {
	Model
	Userid       uint64 `gorm:column:userid`
	Openid       string
	Title        string
	DeviceId     uint64
	DeviceNumber string
	DeviceType   uint8
	Founder      uint8
	AuthorizerId uint32
	ProvinceCode string
	CityCode     string
	CityName     string
	AreaCode     string
	AreaName     string
}

func (UserDevice) TableName() string {
	return "tbl_user_device"
}

// UserDevice 根据主键获取记录
func (ud UserDevice) Get(id uint64) (UserDevice, error) {
	var user UserDevice

	result := db.Take(&user, id)

	return user, result.Error
}

// GetUserByDeviceId 根据设备ID 获取所有绑定的用户列表
func (ud UserDevice) GetUserByDeviceId(deviceId uint64) []UserDevice {
	var users []UserDevice
	db.Model(&ud).Where("device_id = ?", deviceId).Find(&users)

	return users
}

// GetUserByDevivceNumber 根据设备号获取用户列表
func (ud UserDevice) GetUserByDeviceNumber(deviceNumber string) []UserDevice {
	var users []UserDevice
	db.Model(&ud).Where("device_number = ?", deviceNumber).Find(&users)

	return users
}
