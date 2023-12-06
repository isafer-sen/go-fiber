package model

import "time"

// Devices 设备表
type Devices struct {
	ID                int64     `json:"id" gorm:"id"`                                   // 编号
	RoomId            int64     `json:"room_id" gorm:"room_id"`                         // 房间ID
	DeviceAddr        string    `json:"device_addr" gorm:"device_addr"`                 // 设备地址
	DeviceTypeId      int64     `json:"device_type_id" gorm:"device_type_id"`           // 设备类型：1视频监控 2电子班牌 3空调 4感应器 5照明开关 6电动窗帘 7电源插座 8温湿度 9门禁
	EntityId          string    `json:"entity_id" gorm:"entity_id"`                     // 实体id
	DeviceName        string    `json:"device_name" gorm:"device_name"`                 // 设备名称
	AsDeviceName      string    `json:"as_device_name" gorm:"as_device_name"`           // 设备别名
	DeviceStatus      string    `json:"device_status" gorm:"device_status"`             // off 离线 on 在线
	Temperature       string    `json:"temperature" gorm:"temperature"`                 // 空调温度
	LockTopic         string    `json:"lock_topic" gorm:"lock_topic"`                   // 锁下发主题
	TasmotaTag        string    `json:"tasmota_tag" gorm:"tasmota_tag"`                 // tasmota 设备标识
	CurrentPosition   string    `json:"current_position" gorm:"current_position"`       // 窗帘位置
	UnitOfMeasurement string    `json:"unit_of_measurement" gorm:"unit_of_measurement"` // 状态后缀
	VideoPic          string    `json:"video_pic" gorm:"video_pic"`                     // 视频封面
	VideoStreamPic    string    `json:"video_stream_pic" gorm:"video_stream_pic"`       // 视频流地址
	LastChanged       time.Time `json:"last_changed" gorm:"last_changed"`               // 最后改变时间
	LastUpdated       time.Time `json:"last_updated" gorm:"last_updated"`               // 最后更改时间
	ElectricQuantity  string    `json:"electric_quantity" gorm:"electric_quantity"`     // 锁电量
	CreatedAt         time.Time `json:"created_at" gorm:"created_at"`                   // 创建时间
	UpdatedAt         time.Time `json:"updated_at" gorm:"updated_at"`                   // 更新时间
}
