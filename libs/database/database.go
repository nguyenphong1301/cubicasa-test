package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

func NewDatabase(extension, conStr string, OpenConnect, IdleCon int, lifeTime time.Duration, showLog bool) (*gorm.DB, error) {
	db, err := gorm.Open(extension, conStr)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxOpenConns(OpenConnect)
	db.DB().SetMaxIdleConns(IdleCon)
	db.DB().SetConnMaxLifetime(lifeTime * time.Minute)
	db.LogMode(showLog)
	return db, nil
}
