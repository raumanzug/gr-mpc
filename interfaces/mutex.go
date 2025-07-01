package interfaces

import (
	"sync"
)

type IMutexEquipped interface {
	GetMutex() *sync.Mutex
}

type MutexEquippedBase_t struct {
	mutex sync.Mutex
}

func (meBase MutexEquippedBase_t) GetMutex() *sync.Mutex {
	return &meBase.mutex
}
