gen_grapql:
	go generate ./...

fmt:
	gofmt -s -l -w .

gen_proto:
	mkdir -p go
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	export PATH="$PATH:$(go env GOPATH)/bin"
	protoc-gen-go --version
	protoc -I=./../../proto --go_out=go --go-grpc_out=go ./../../proto/*.proto
	go mod tidy
