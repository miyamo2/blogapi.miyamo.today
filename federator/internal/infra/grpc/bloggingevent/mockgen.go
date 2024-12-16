//go:generate mockgen -source=blogging_event_grpc.pb.go -destination=../../../mock/infra/grpc/bloggingevent/mock_blogging_event_grpc.pb.go -package=$GOPACKAGE
package bloggingevent
