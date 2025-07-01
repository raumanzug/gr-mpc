package interfaces

import (
	"sync"
)

type IGateway interface {
	Mutex() *sync.Mutex
	Release(func(IEvent) bool) IEvent
	Send(IEvent)
}
