module github.com/daystram/audit/audit-tr

go 1.16

require (
	github.com/daystram/audit/proto v0.0.0
	google.golang.org/grpc v1.39.0
)

replace github.com/daystram/audit/proto v0.0.0 => ../proto
