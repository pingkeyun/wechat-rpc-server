package models

type Dinuan struct {
	Model
	DeviceId  uint64
	StartDate string
	EndDate   string
}

func (Dinuan) TableName() string {
	return "tbl_dinuan_config"
}

// GetUsersByDeviceId 根据设备ID获取所有绑定此设备的用户列表
// startDate 格式：月-日， 如 03-15
func (d Dinuan) GetDeviceIdByStartDate(startDate string) []uint64 {
	var result []uint64
	db.Model(&d).Where("start_date = ?", startDate).Pluck("device_id", &result)

	return result
}

// GetDeviceIdByEndDate 根据结束日期查询
func (d Dinuan) GetDeviceIdByEndDate(endDate string) []uint64 {
	var result []uint64
	db.Model(&d).Where("end_date = ?", endDate).Pluck("device_id", &result)

	return result
}
