package nvidiasmi

import (
	"reflect"
	"sync"
)

var (
	MemoryStoreData MemoryStore
	mu              sync.RWMutex
	UpdateChan      = make(chan struct{})
)

type MemoryStore struct {
	NvidiaSMIInfo NvidiaSMIInfo
}

func (ms *MemoryStore) staticIsEmpty() bool {
	return reflect.DeepEqual(ms.NvidiaSMIInfo.StaticNvidiaSMIInfo, StaticNvidiaSMIInfo{})
}
