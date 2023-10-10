.PHONY: proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ${PROTO_FILE}

.PHONY: build-grpc
build-grpc:
	go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main cmd/server/main.go

.PHONY: test
test:
	go test -v ./test...