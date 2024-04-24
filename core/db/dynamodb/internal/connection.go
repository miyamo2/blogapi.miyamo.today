package internal

import (
	"github.com/miyamo2/api.miyamo.today/core/db/dynamodb/client"
	"sync"
)

var (
	Lock   sync.RWMutex
	Client client.Client
)
