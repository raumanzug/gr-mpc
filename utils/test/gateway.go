package test

import (
	"log"
	"sync"

	testUtilsInterfaces "github.com/raumanzug/gr-mpc/utils/test/interfaces"
)

type gateway_t struct {
	pMutex          *sync.Mutex
	pBackendCondVar *sync.Cond
	pTestCondVar    *sync.Cond

	condition func(testUtilsInterfaces.IEvent) bool
	out       testUtilsInterfaces.IEvent
}

func newGateway() testUtilsInterfaces.IGateway {
	var mutex sync.Mutex
	gateway := gateway_t{
		pMutex: &mutex,
	}
	gateway.pBackendCondVar = sync.NewCond(gateway.pMutex)
	gateway.pTestCondVar = sync.NewCond(gateway.pMutex)

	return &gateway
}

func (pGw *gateway_t) Mutex() *sync.Mutex {
	return pGw.pMutex
}

func (pGw *gateway_t) Release(
	condition func(testUtilsInterfaces.IEvent) bool,
) (out testUtilsInterfaces.IEvent) {
	pGw.condition = condition
	pGw.pBackendCondVar.Broadcast()
	pGw.pTestCondVar.Wait()
	out = pGw.out
	pGw.out = nil
	log.Println(
		"event " +
			testUtilsInterfaces.IEvent2String(out) +
			" released",
	)

	return
}

func (pGw *gateway_t) Send(event testUtilsInterfaces.IEvent) {
	log.Println(
		"event " +
			testUtilsInterfaces.IEvent2String(event) +
			" sent",
	)

	pGw.pMutex.Lock()
	defer pGw.pMutex.Unlock()

	for pGw.out != nil ||
		pGw.condition == nil ||
		!pGw.condition(event) {
		pGw.pBackendCondVar.Wait()
	}
	pGw.condition = nil
	pGw.out = event

	pGw.pTestCondVar.Signal()
}
