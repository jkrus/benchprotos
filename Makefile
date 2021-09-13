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
		--go_opt=module=benchprotos \
		--go-grpc_out=. \
		--go-grpc_opt=module=benchprotos \
		--proto_path=grpc/proto grpc/proto/*.proto

	capnp compile \
		--import-path=. \
		--output=go   crpc/cproto/*.capnp

	mv crpc/cproto/*.go crpc/cpb

gen id:
			capnp compile \
            --import-path=. \
             -ocapnp   crpc/cproto/*.capnp



