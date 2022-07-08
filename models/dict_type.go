package models

import "time"

type DictType struct {
	Id         int64     `gorm:"column:id"`
	Name       string    `gorm:"column:name"`        // 字典类型名称
	Status     int       `gorm:"column:status"`      // 状态，1 未删除，0 禁用，-1 删除
	CreateTime time.Time `gorm:"column:create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time"` // 修改时间
}

func (dt DictType) TableName() string {
	return "sys_dict_type"
}
