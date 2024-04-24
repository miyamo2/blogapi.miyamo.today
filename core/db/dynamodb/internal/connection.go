package internal

import (
	"sync"

	"github.com/miyamo2/blogapi.miyamo.today/core/db/dynamodb/client"
)

var (
	Lock   sync.RWMutex
	Client client.Client
)
