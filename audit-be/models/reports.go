package models

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

const (
	_ReportMeasurement  = "report"
	_ReportFieldLatency = "latency"
	_ReportQueryWindow  = "-3h"
)

type Report struct {
	ApplicationID string `json:"-"`
	ServiceID     string `json:"-"`
	Latency       int64  `json:"-"`
	Timestamp     int64  `json:"-"`
}

type reportOrm struct {
	bucket   string
	writeAPI api.WriteAPIBlocking
	queryAPI api.QueryAPI
}

type ReportOrmer interface {
	GetWindowByApplicationIDAndServiceID(applicationID, serviceID string) (reports []Report, err error)
	Insert(report Report) (err error)
}

func NewReportOrmer(influx influxdb2.Client, organization, bucket string) ReportOrmer {
	return &reportOrm{
		bucket:   bucket,
		writeAPI: influx.WriteAPIBlocking(organization, bucket),
		queryAPI: influx.QueryAPI(organization),
	}
}

func (o *reportOrm) GetWindowByApplicationIDAndServiceID(applicationID, serviceID string) (reports []Report, err error) {
	var result *api.QueryTableResult
	reports = make([]Report, 0)
	if result, err = o.queryAPI.Query(context.Background(), fmt.Sprintf(`
	from(bucket: "%s")
		|> range(start: %s)
		|> filter(fn: (r) => r["_measurement"] == "%s")
		|> filter(fn: (r) => r["application_id"] == "%s")
		|> filter(fn: (r) => r["service_id"] == "%s")
		|> filter(fn: (r) => r["_field"] == "%s")
		|> aggregateWindow(every: 10s, fn: mean, createEmpty: false)
		|> yield(name: "mean")
	`, o.bucket, _ReportQueryWindow, _ReportMeasurement, applicationID, serviceID, _ReportFieldLatency)); err == nil {
		for result.Next() {
			record := result.Record()
			reports = append(reports, Report{
				ApplicationID: applicationID,
				ServiceID:     serviceID,
				Latency:       int64(result.Record().Value().(float64)),
				Timestamp:     record.Time().UnixNano(),
			})
		}
		if result.Err() != nil {
			return make([]Report, 0), fmt.Errorf("query parsing error: %s", result.Err().Error())
		}
	}
	return
}

func (o *reportOrm) Insert(report Report) (err error) {
	point := influxdb2.NewPoint(_ReportMeasurement,
		map[string]string{
			"application_id": report.ApplicationID,
			"service_id":     report.ServiceID,
		},
		map[string]interface{}{
			_ReportFieldLatency: report.Latency,
		},
		time.Unix(0, report.Timestamp))
	return o.writeAPI.WritePoint(context.Background(), point)
}
