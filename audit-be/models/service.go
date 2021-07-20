package models

import (
	"gorm.io/gorm"
)

// CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
type Service struct {
	ID            string      `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()" json:"-"`
	Application   Application `gorm:"foreignKey:ApplicationID;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	ApplicationID string      `json:"-"`

	Name        string `gorm:"column:name;type:varchar(20);not null" json:"-"`
	Description string `gorm:"column:description;type:varchar(50)" json:"-"`
	Endpoint    string `gorm:"column:endpoint;type:varchar(50)" json:"-"`
	Type        string `gorm:"column:type;type:varchar(5)" json:"-"`
	Config      string `gorm:"column:config;type:text" json:"-"`
	Showcase    bool   `gorm:"column:showcase;default:false" json:"-"`

	CreatedAt int64 `gorm:"column:created_at;autoCreateTime" json:"-"`
	UpdatedAt int64 `gorm:"column:updated_at;autoUpdateTime" json:"-"`
}

type serviceOrm struct {
	db *gorm.DB
}

type ServiceOrmer interface {
	GetAllByApplicationID(applicationID string) (services []Service, err error)
	GetOneByIDAndApplicationID(ID, applicationID string) (service Service, err error)
	Insert(service Service) (ID string, err error)
	Update(service Service) (err error)
	DeleteByIDAndApplicationID(ID, applicationID string) (err error)
}

func NewServiceOrmer(db *gorm.DB) ServiceOrmer {
	_ = db.AutoMigrate(&Service{})
	return &serviceOrm{db}
}

func (o *serviceOrm) GetAllByApplicationID(applicationID string) (services []Service, err error) {
	result := o.db.Model(&Service{}).Where("application_id = ?", applicationID).Order("name").Find(&services)
	return services, result.Error
}

func (o *serviceOrm) GetOneByIDAndApplicationID(ID, applicationID string) (service Service, err error) {
	result := o.db.Model(&Service{}).Where("id = ? AND application_id = ?", ID, applicationID).First(&service)
	return service, result.Error
}

func (o *serviceOrm) Insert(service Service) (ID string, err error) {
	result := o.db.Model(&Service{}).Create(&service)
	return service.ID, result.Error
}

func (o *serviceOrm) Update(service Service) (err error) {

	result := o.db.Model(&Service{}).Where("id = ? AND application_id = ?", service.ID, service.ApplicationID).Select(
		"Name",
		"Description",
		"Endpoint",
		"Type",
		"Config",
		"Showcase",
		"CreatedAt",
		"UpdatedAt",
	).Updates(&service)
	return result.Error
}

func (o *serviceOrm) DeleteByIDAndApplicationID(ID, applicationID string) (err error) {
	result := o.db.Model(&Service{}).Where("id = ? AND application_id = ?", ID, applicationID).Delete(Service{})
	return result.Error
}
