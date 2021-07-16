package models

import (
	"gorm.io/gorm"
)

// CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
type Service struct {
	ID            string      `gorm:"column:ID;primaryKey;type:uuid;default:uuid_generate_v4()" json:"-"`
	Application   Application `gorm:"foreignKey:ApplicationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	ApplicationID string      `json:"-"`

	Name        string `gorm:"column:Name;type:varchar(20);not null" json:"-"`
	Description string `gorm:"column:Description;type:varchar(50)" json:"-"`
	Endpoint    string `gorm:"column:Endpoint;type:varchar(50)" json:"-"`
	Type        byte   `gorm:"column:Type;type:smallint" json:"-"`
	Config      string `gorm:"column:Config;type:text" json:"-"`
	Showcase    bool   `gorm:"column:Showcase;default:false" json:"-"`

	CreatedAt int64 `gorm:"column:CreatedAt;autoCreateTime" json:"-"`
	UpdatedAt int64 `gorm:"column:UpdatedAt;autoUpdateTime" json:"-"`
}

type serviceOrm struct {
	db *gorm.DB
}

type ServiceOrmer interface {
}

func NewServiceOrmer(db *gorm.DB) ServiceOrmer {
	_ = db.AutoMigrate(&Service{})
	return &serviceOrm{db}
}
