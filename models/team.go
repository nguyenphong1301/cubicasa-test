package models

import (
	"time"
)

type Team struct {
	ID        int64      `gorm:"column=id;PRIMARY;" json:"id"`
	Name      string     `gorm:"column=name;index;not null;" json:"name"`
	TeamType  string     `gorm:"column=team_type;index;not null;" json:"team_type"`
	HubID     int64      `gorm:"column=hub_id;" json:"hub_id"`
	CreatedAt *time.Time `gorm:"column=created_at;autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column=updated_at;autoUpdateTime" json:"updated_at"`
}

func (m *Team) TableName() string {
	return `teams`
}

func (m *Team) AssignHub(hubId int64) {
	m.HubID = hubId
}
