//go:generate mockgen -source=article_grpc.pb.go -destination=../../../mock/infra/grpc/article/mock_article_grpc.pb.go -package=$GOPACKAGE
package article
