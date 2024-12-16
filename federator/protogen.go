//go:build protogen

//go:generate protoc --proto_path=./.proto/article --go_out=. --go_opt=Marticle.proto=internal/infra/grpc/article --go-grpc_out=. --go-grpc_opt=Marticle.proto=internal/infra/grpc/article ./.proto/article/article.proto
//go:generate protoc --proto_path=./.proto/tag --go_out=. --go_opt=Mtag.proto=internal/infra/grpc/tag --go-grpc_out=. --go-grpc_opt=Mtag.proto=internal/infra/grpc/tag ./.proto/tag/tag.proto
//go:generate protoc --proto_path=./.proto/blogging_event --go_out=. --go_opt=Mblogging_event.proto=internal/infra/grpc/bloggingevent --go-grpc_out=. --go-grpc_opt=Mblogging_event.proto=internal/infra/grpc/bloggingevent ./.proto/blogging_event/blogging_event.proto
package federator
