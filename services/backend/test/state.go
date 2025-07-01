package test

import (
	"github.com/raumanzug/gr-mpc/interfaces"
	testUtilsGlobals "github.com/raumanzug/gr-mpc/utils/test/globals"
	testUtilsInterfaces "github.com/raumanzug/gr-mpc/utils/test/interfaces"
)

type state_t struct {
	interfaces.ControlsState
	interfaces.MutexEquippedBase_t

	server server_t
}

func newState(server server_t) interfaces.IState {
	state := state_t{
		server: server,
	}

	state.InitControlsState()

	return &state
}

func (pState *state_t) Close() (err error) {
	event := testUtilsInterfaces.CloseEvent_t{}
	event.SetDisplayName(pState.server.dn)
	event.SetState(pState)
	testUtilsGlobals.Gateway.Send(&event)

	return
}

func (pState *state_t) OpenBackendControls() (
	controls interfaces.IBackendControls,
	err error,
) {
	controls = newBC(pState)

	event := testUtilsInterfaces.OpenBackendControlsEvent_t{}
	event.SetDisplayName(pState.server.dn)
	event.SetState(pState)
	testUtilsGlobals.Gateway.Send(&event)

	return
}

func (pState *state_t) Start() (err error) {
	event := testUtilsInterfaces.LifecycleStartEvent_t{}
	event.SetDisplayName(pState.server.dn)
	event.SetState(pState)
	testUtilsGlobals.Gateway.Send(&event)

	return
}

func (pState *state_t) Stop() (err error) {
	event := testUtilsInterfaces.LifecycleStopEvent_t{}
	event.SetDisplayName(pState.server.dn)
	event.SetState(pState)
	testUtilsGlobals.Gateway.Send(&event)

	return
}
