module github.com/daystram/audit/audit-tr

go 1.16

require (
	github.com/daystram/audit/proto v0.0.0
	github.com/go-ping/ping v0.0.0-20210506233800-ff8be3320020
	github.com/spf13/viper v1.8.1
	google.golang.org/grpc v1.39.0
)

replace github.com/daystram/audit/proto v0.0.0 => ../proto
