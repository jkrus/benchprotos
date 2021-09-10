.PHONY: pre-push go-mod
pre-push: go-mod

go-mod:
	go mod tidy -v
	go mod download -x

.PHONY: protoc-upgrade protoc-update generate go-mod
protoc-upgrade: protoc-update generate go-mod

protoc-update:
	go get -u github.com/golang/protobuf/protoc-gen-go@latest
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

generate:
	protoc \
		-I=. \
		--go_out=. \
		--go_opt=module=github.com/jkrus/benchprotos \
		--go-grpc_out=. \
		--go-grpc_opt=module=github.com/jkrus/benchprotos \
		--proto_path=grpc/proto grpc/proto/*.proto
