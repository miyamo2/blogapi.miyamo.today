//go:generate mockgen -source=blogging_event.connect.go -destination=../../../../mock/infra/grpc/blogging_event/blogging_eventconnect/mock_blogging_event.connect.go -package=$GOPACKAGE
package blogging_eventconnect
