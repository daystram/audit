package models

import (
	"gorm.io/gorm"
)

type applicationOrm struct {
	db *gorm.DB
}

// CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
type Application struct {
	ID          string `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()" json:"-"`
	Name        string `gorm:"column:name;type:varchar(20);not null" json:"-"`
	Description string `gorm:"column:description;type:varchar(50)" json:"-"`
	CreatedAt   int64  `gorm:"column:created_at;autoCreateTime" json:"-"`
	UpdatedAt   int64  `gorm:"column:updated_at;autoUpdateTime" json:"-"`
}

type ApplicationOrmer interface {
	GetAll() (applications []Application, err error)
	GetOneByID(ID string) (application Application, err error)
	Insert(application Application) (ID string, err error)
	Update(application Application) (err error)
	DeleteByID(ID string) (err error)
}

func NewApplicationOrmer(db *gorm.DB) ApplicationOrmer {
	_ = db.AutoMigrate(&Application{})
	return &applicationOrm{db}
}

func (o *applicationOrm) GetAll() (applications []Application, err error) {
	result := o.db.Model(&Application{}).Order("name").Find(&applications)
	return applications, result.Error
}

func (o *applicationOrm) GetOneByID(ID string) (application Application, err error) {
	result := o.db.Model(&Application{}).Where("id = ?", ID).First(&application)
	return application, result.Error
}

func (o *applicationOrm) Insert(application Application) (ID string, err error) {
	result := o.db.Model(&Application{}).Create(&application)
	return application.ID, result.Error
}

func (o *applicationOrm) Update(application Application) (err error) {
	result := o.db.Model(&Application{}).Where("id = ?", application.ID).Updates(&application)
	return result.Error
}

func (o *applicationOrm) DeleteByID(ID string) (err error) {
	result := o.db.Model(&Application{}).Where("id = ?", ID).Delete(Application{})
	return result.Error
}
