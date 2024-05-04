package provider

import (
	"net"
	"os"

	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/infra/tcp"
	"go.uber.org/fx"
)

var Tcp = fx.Options(
	fx.Provide(func() net.Listener { return tcp.MustListen(tcp.WithPort(os.Getenv("PORT"))) }),
)
