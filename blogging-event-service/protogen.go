//go:build protogen

//go:generate protoc --proto_path=./.proto/blogging_event --go_out=. --go_opt=Mblogging_event.proto=internal/infra/grpc --go-grpc_out=. --go-grpc_opt=Mblogging_event.proto=internal/infra/grpc ./.proto/blogging_event/blogging_event.proto
package blogging_event_service
