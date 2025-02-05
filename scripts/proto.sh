rm -rf pkg/proto/server/*.pb.go

protoc --go_out=. --go-grpc_out=. pkg/proto/server/server.proto
