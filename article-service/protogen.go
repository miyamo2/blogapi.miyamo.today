//go:build protogen

//go:generate protoc --proto_path=./.proto/article --go_out=. --go_opt=Marticle.proto=internal/infra/grpc --go-grpc_out=. --go-grpc_opt=Marticle.proto=internal/infra/grpc ./.proto/article/article.proto
package article_service
