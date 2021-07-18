.PHONY: all build_proto generate_be_mock clean

all: build_proto generate_audit-be_mock generate_proto_mock

build_proto:
	@echo "Building protobufs..."
	@protoc --go_out=. proto/*.proto --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative

generate_audit-be_mock:
	@echo "Generating audit-be mocks..."
	@cd audit-be && find handlers models -type f ! -name "*_test*" | xargs -L1 -I{}  sh -c 'OUT=mocks/$${1} && mkdir -p $${OUT%/*} && touch $${OUT} && mockgen --source $${1} > $${OUT} && echo $$OUT' -- {}

generate_proto_mock:
	@echo "Generating proto mocks..."
	@cd proto && find . -type f -name "*.go" | xargs -L1 -I{}  sh -c 'OUT=mocks/$${1} && mkdir -p $${OUT%/*} && touch $${OUT} && mockgen --source $${1} > $${OUT}' -- {}

clean:
	rm -f proto/*.go
	rm -rf audit-be/mocks 
	rm -rf proto/mocks 
