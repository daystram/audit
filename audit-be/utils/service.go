package utils

import (
	"github.com/daystram/audit/audit-be/constants"
	"github.com/daystram/audit/proto"
)

func ServiceTypeToProto(serviceType string) proto.ServiceType {
	switch serviceType {
	case constants.ServiceTypeHTTP:
		return proto.ServiceType_SERVICE_TYPE_HTTP
	case constants.ServiceTypeTCP:
		return proto.ServiceType_SERVICE_TYPE_TCP
	case constants.ServiceTypePING:
		return proto.ServiceType_SERVICE_TYPE_PING
	default:
		return proto.ServiceType_SERVICE_TYPE_UNKNOWN
	}
}
