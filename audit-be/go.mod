module github.com/daystram/audit/audit-be

go 1.16

require (
	github.com/daystram/audit/proto v0.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/mock v1.6.0
	github.com/influxdata/influxdb-client-go/v2 v2.4.0
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.5.1
	google.golang.org/grpc v1.39.0
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.8

)

replace github.com/daystram/audit/proto v0.0.0 => ../proto
