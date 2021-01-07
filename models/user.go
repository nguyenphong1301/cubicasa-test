package models

import "time"

type User struct {
	ID        int64      `gorm:"column=id;PRIMARY;" json:"id"`
	Role      string     `gorm:"column=role" json:"role"`
	Email     string     `gorm:"column=email;index;not null;unique" json:"email"`
	TeamID    int64      `gorm:"column=team_id" json:"team_id"`
	CreatedAt *time.Time `gorm:"column=created_at;autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column=updated_at;autoUpdateTime" json:"updated_at"`
}

func (m *User) TableName() string {
	return `users`
}

func (m *User) AssignTeam(teamID int64) {
	m.TeamID = teamID
}
