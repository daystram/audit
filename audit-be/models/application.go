package models

import (
	"gorm.io/gorm"
)

type applicationOrm struct {
	db *gorm.DB
}

// CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
type Application struct {
	ID          string `gorm:"column:ID;primaryKey;type:uuid;default:uuid_generate_v4()" json:"-"`
	Name        string `gorm:"column:Name;type:varchar(20);not null" json:"-"`
	Description string `gorm:"column:Description;type:varchar(50)" json:"-"`
	CreatedAt   int64  `gorm:"column:CreatedAt;autoCreateTime" json:"-"`
	UpdatedAt   int64  `gorm:"column:UpdatedAt;autoUpdateTime" json:"-"`
}

type ApplicationOrmer interface {
}

func NewApplicationOrmer(db *gorm.DB) ApplicationOrmer {
	_ = db.AutoMigrate(&Application{})
	return &applicationOrm{db}
}
