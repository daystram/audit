.PHONY: all build_proto generate_be_mock clean

all: build_proto generate_be_mock

build_proto:
	@echo "Building protobufs..."
	@protoc --go_out=. proto/*.proto --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative

generate_be_mock:
	@echo "Building audit-be mocks..."
	@find audit-be/handlers audit-be/models -type f ! -name "*_test*" | xargs -L1 -I{}  sh -c 'OUT=$${1%%/*}/mocks/$${1#*/} && mkdir -p $${OUT%/*} && touch $${OUT} && mockgen --source $${1} > $${OUT}' -- {}

clean:
	rm -f proto/*.go
	rm -rf audit-be/mocks 
