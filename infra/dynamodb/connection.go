package dynamodb

import (
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi-core/infra/dynamodb/client"
	"github.com/miyamo2/blogapi-core/infra/dynamodb/internal"
	"github.com/miyamo2/blogapi-core/log"
)

var ErrClientNotInitialized = errors.New("dynamodb client is not initialized")

func Get() (client.Client, error) {
	internal.Lock.RLock()
	defer internal.Lock.RUnlock()
	if internal.Client == nil {
		return nil, ErrClientNotInitialized
	}
	return internal.Client, nil
}

func Initialize(client client.Client) {
	log.DefaultLogger().Info("initialize dynamodb client")
	internal.Lock.Lock()
	defer internal.Lock.Unlock()
	if internal.Client != nil {
		log.DefaultLogger().Warn("dynamodb client is already initialized")
		return
	}
	internal.Client = client
	log.DefaultLogger().Info("completed dynamodb client initialization")
}

func Invalidate() {
	log.DefaultLogger().Info("invalidate dynamodb client")
	internal.Lock.Lock()
	defer internal.Lock.Unlock()
	if internal.Client == nil {
		log.DefaultLogger().Warn("dynamodb client is not initialized")
		return
	}
	internal.Client = nil
	log.DefaultLogger().Info("completed dynamodb client invalidation")
}
