all: pb

pb:
	protoc -I=. \
	-I=${GOPATH}/src \
	-I=${GOPATH}/src/github.com/meshplus/bitxhub-model/pb \
	--gogofaster_out=:. \
	pb.proto

#grpc:
#	cd pb && protoc -I=. \
#	-I=${GOPATH}/src \
#	-I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
#	-I=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
#	--grpc-gateway_out=logtostderr=true:. \
#	--swagger_out=logtostderr=true:. \
#	--gogofaster_out=plugins=grpc:. \
#	broker.proto plugin.proto

## make linter: Run golanci-lint
linter:
	golangci-lint run -E goimports -E bodyclose --skip-dirs-use-default


clean:
	rm pb/*.pb.go
	rm pb/*.json
	rm pb/*.gw.go

.PHONY: pb
