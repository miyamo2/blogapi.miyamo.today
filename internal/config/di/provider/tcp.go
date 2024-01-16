package provider

import (
	"github.com/miyamo2/blogapi-tag-service/internal/infra/tcp"
	"go.uber.org/fx"
	"net"
)

var Tcp = fx.Options(
	fx.Provide(func() net.Listener { return tcp.MustListen(tcp.WithPort("8080")) }),
)
