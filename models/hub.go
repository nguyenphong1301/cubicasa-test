package models

import "time"

type Hub struct {
	ID          int64      `gorm:"column=id;PRIMARY;" json:"id"`
	Name        string     `gorm:"column=name;index;not null;unique" json:"name"`
	GeoLocation string     `gorm:"column=geo_location; not null" json:"geo_location"`
	CreatedAt   *time.Time `gorm:"column=created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column=updated_at;autoUpdateTime" json:"updated_at"`
}

func (m *Hub) TableName() string {
	return `hubs`
}
