package internal

import (
	"github.com/miyamo2/blogapi-core/db/dynamodb/client"
	"sync"
)

var (
	Lock   sync.RWMutex
	Client client.Client
)
