run-genproto-auth:
	protoc pkg/proto/auth-proto.proto --go_out=pkg/proto-pb --go-grpc_out=pkg/proto-pb

run-go-auth:
	go run cmd/authentication/main.go