package handlers

import (
	"gorm.io/gorm"

	"github.com/daystram/audit/audit-be/datatransfers"
	"github.com/daystram/audit/audit-be/models"
)

func (m *module) ApplicationGetAll() (applicationInfos []datatransfers.ApplicationInfo, err error) {
	var applications []models.Application
	if applications, err = m.db.applicationOrmer.GetAll(); err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	for _, application := range applications {
		applicationInfos = append(applicationInfos, datatransfers.ApplicationInfo{
			ID:          application.ID,
			Name:        application.Name,
			Description: application.Description,
			CreatedAt:   application.CreatedAt,
			UpdatedAt:   application.UpdatedAt,
		})
	}
	return
}

func (m *module) ApplicationGetOne(applicationID string) (applicationInfo datatransfers.ApplicationInfo, err error) {
	var application models.Application
	if application, err = m.db.applicationOrmer.GetOneByID(applicationID); err != nil {
		return datatransfers.ApplicationInfo{}, err
	}
	applicationInfo = datatransfers.ApplicationInfo{
		ID:          application.ID,
		Name:        application.Name,
		Description: application.Description,
		CreatedAt:   application.CreatedAt,
		UpdatedAt:   application.UpdatedAt,
	}
	return
}

func (m *module) ApplicationCreate(applicationInfo datatransfers.ApplicationInfo) (applicationID string, err error) {
	if applicationID, err = m.db.applicationOrmer.Insert(models.Application{
		Name:        applicationInfo.Name,
		Description: applicationInfo.Description,
	}); err != nil {
		return "", err
	}
	return
}

func (m *module) ApplicationUpdate(applicationInfo datatransfers.ApplicationInfo) (err error) {
	if err = m.db.applicationOrmer.Update(models.Application{
		ID:          applicationInfo.ID,
		Name:        applicationInfo.Name,
		Description: applicationInfo.Description,
	}); err != nil {
		return err
	}
	return
}

func (m *module) ApplicationDelete(applicationID string) (err error) {
	if err = m.db.applicationOrmer.DeleteByID(applicationID); err != nil {
		return err
	}
	return
}
