package handlers

import (
	"gorm.io/gorm"

	"github.com/daystram/audit/audit-be/datatransfers"
	"github.com/daystram/audit/audit-be/models"
)

func (m *module) MonitorGetAll() (applicationInfos []datatransfers.ApplicationInfo, err error) {
	var applications []models.Application
	applicationInfos = make([]datatransfers.ApplicationInfo, 0)
	if applications, err = m.db.applicationOrmer.GetAllShowcaseWithServices(); err == gorm.ErrRecordNotFound {
		return applicationInfos, nil
	} else if err != nil {
		return nil, err
	}
	for _, application := range applications {
		serviceInfos := make([]datatransfers.ServiceInfo, 0)
		for _, service := range application.Services {
			serviceInfos = append(serviceInfos, datatransfers.ServiceInfo{
				ID:            service.ID,
				ApplicationID: service.ApplicationID,
				Name:          service.Name,
				Description:   service.Description,
				Endpoint:      service.Endpoint,
				Type:          service.Type,
				Config:        service.Config,
				Enabled:       service.Enabled,
				Showcase:      service.Showcase,
				CreatedAt:     service.CreatedAt,
				UpdatedAt:     service.UpdatedAt,
			})
		}
		applicationInfos = append(applicationInfos, datatransfers.ApplicationInfo{
			ID:          application.ID,
			Services:    serviceInfos,
			Name:        application.Name,
			Description: application.Description,
			CreatedAt:   application.CreatedAt,
			UpdatedAt:   application.UpdatedAt,
		})
	}
	return
}
