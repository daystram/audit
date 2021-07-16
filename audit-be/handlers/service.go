package handlers

import (
	"gorm.io/gorm"

	"github.com/daystram/audit/audit-be/datatransfers"
	"github.com/daystram/audit/audit-be/models"
)

func (m *module) ServiceGetAll(applicationID string) (serviceInfos []datatransfers.ServiceInfo, err error) {
	var services []models.Service
	if services, err = m.db.serviceOrmer.GetAllByApplicationID(applicationID); err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	for _, service := range services {
		serviceInfos = append(serviceInfos, datatransfers.ServiceInfo{
			ID:            service.ID,
			ApplicationID: service.ApplicationID,
			Name:          service.Name,
			Description:   service.Description,
			Endpoint:      service.Endpoint,
			Type:          service.Type,
			Config:        service.Config,
			Showcase:      service.Showcase,
			CreatedAt:     service.CreatedAt,
			UpdatedAt:     service.UpdatedAt,
		})
	}
	return
}

func (m *module) ServiceGetOne(serviceID, applicationID string) (serviceInfo datatransfers.ServiceInfo, err error) {
	var service models.Service
	if service, err = m.db.serviceOrmer.GetOneByIDAndApplicationID(serviceID, applicationID); err != nil {
		return datatransfers.ServiceInfo{}, err
	}
	serviceInfo = datatransfers.ServiceInfo{
		ID:            service.ID,
		ApplicationID: service.ApplicationID,
		Name:          service.Name,
		Description:   service.Description,
		Endpoint:      service.Endpoint,
		Type:          service.Type,
		Config:        service.Config,
		Showcase:      service.Showcase,
		CreatedAt:     service.CreatedAt,
		UpdatedAt:     service.UpdatedAt,
	}
	return
}

func (m *module) ServiceCreate(serviceInfo datatransfers.ServiceInfo) (serviceID string, err error) {
	if serviceID, err = m.db.serviceOrmer.Insert(models.Service{
		ApplicationID: serviceInfo.ApplicationID,
		Name:          serviceInfo.Name,
		Description:   serviceInfo.Description,
		Endpoint:      serviceInfo.Endpoint,
		Type:          serviceInfo.Type,
		Config:        serviceInfo.Config,
		Showcase:      serviceInfo.Showcase,
	}); err != nil {
		return "", err
	}
	return
}

func (m *module) ServiceUpdate(serviceInfo datatransfers.ServiceInfo) (err error) {
	if err = m.db.serviceOrmer.Update(models.Service{
		ID:            serviceInfo.ID,
		ApplicationID: serviceInfo.ApplicationID,
		Name:          serviceInfo.Name,
		Description:   serviceInfo.Description,
		Endpoint:      serviceInfo.Endpoint,
		Type:          serviceInfo.Type,
		Config:        serviceInfo.Config,
		Showcase:      serviceInfo.Showcase,
	}); err != nil {
		return err
	}
	return
}

func (m *module) ServiceDelete(serviceID, applicationID string) (err error) {
	if err = m.db.serviceOrmer.DeleteByIDAndApplicationID(serviceID, applicationID); err != nil {
		return err
	}
	return
}
