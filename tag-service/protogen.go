//go:build protogen

//go:generate protoc --proto_path=./.proto/tag --go_out=. --go_opt=Mtag.proto=internal/infra/grpc --go-grpc_out=. --go-grpc_opt=Mtag.proto=internal/infra/grpc ./.proto/tag/tag.proto
package tag_service
