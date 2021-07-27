package handlers

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"

	"github.com/daystram/audit/audit-be/constants"
	"github.com/daystram/audit/audit-be/models"
	"github.com/daystram/audit/audit-be/utils"
	"github.com/daystram/audit/proto"
)

func (m *module) InitializeScheduler() (err error) {
	m.schedulerCron = cron.New()
	if _, err = m.schedulerCron.AddFunc(constants.ShedluerCronSpec, m.TriggerTracking); err != nil {
		return err
	}
	m.schedulerCron.Start()
	return
}

func (m *module) TriggerTracking() {
	var err error
	var services []models.Service
	if services, err = m.db.serviceOrmer.GetAllEnabled(); err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err)
		return
	}
	for _, service := range services {
		if err = m.trackerServer.SendTrackingRequest(&proto.TrackingRequest{
			ApplicationId: service.ApplicationID,
			ServiceId:     service.ID,
			Endpoint:      service.Endpoint,
			Type:          utils.ServiceTypeToProto(service.Type),
			Config:        service.Config,
			RequestedAt:   time.Now().UnixNano(),
		}); err != nil {
			log.Println(err)
		}
	}
}
