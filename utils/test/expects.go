package test

import (
	"github.com/raumanzug/gr-mpc/interfaces"
	testGlobals "github.com/raumanzug/gr-mpc/utils/test/globals"

	testUtilsInterfaces "github.com/raumanzug/gr-mpc/utils/test/interfaces"
)

func EventCondition[T testUtilsInterfaces.IEvent](
	dn string,
	state interfaces.IState,
) func(testUtilsInterfaces.IEvent) bool {
	return func(event testUtilsInterfaces.IEvent) bool {
		pEvent, ok := event.(T)
		if !ok {
			return false
		}
		if pEvent.GetDisplayName() != dn {
			return false
		}
		if pEvent.GetState() != state {
			return false
		}

		return true
	}
}

func ExpectCloseEvent(
	dn string,
	state interfaces.IState,
) {
	gateway := testGlobals.Gateway
	gateway.Release(
		EventCondition[*testUtilsInterfaces.CloseEvent_t](
			dn,
			state,
		),
	)
}

func ExpectEventConnection[T testUtilsInterfaces.IEvent](
	dn string,
	state interfaces.IState,
) {
	gateway := testGlobals.Gateway
	gateway.Release(
		EventCondition[*testUtilsInterfaces.OpenBackendControlsEvent_t](
			dn,
			state,
		),
	)
	gateway.Release(EventCondition[T](dn, state))
	gateway.Release(EventCondition[*testUtilsInterfaces.CloseBackendControlsEvent_t](dn, state))
}
