package test

import (
	"github.com/raumanzug/gr-mpc/interfaces"

	testUtilsData "github.com/raumanzug/gr-mpc/utils/test/data"
	testUtilsGlobals "github.com/raumanzug/gr-mpc/utils/test/globals"
	testUtilsInterfaces "github.com/raumanzug/gr-mpc/utils/test/interfaces"
)

type backendControls_t struct {
	dn             string
	state          interfaces.IState
	pControlsState *interfaces.ControlsState
}

func newBC(
	pState *state_t,
) interfaces.IBackendControls {
	backendControls := backendControls_t{
		dn:             pState.server.dn,
		state:          pState,
		pControlsState: pState.GetControlsState(),
	}

	return &backendControls
}

func (pBC *backendControls_t) Close() (err error) {
	event := testUtilsInterfaces.CloseBackendControlsEvent_t{}
	event.SetDisplayName(pBC.dn)
	event.SetState(pBC.state)
	testUtilsGlobals.Gateway.Send(&event)

	return
}

func (pBC *backendControls_t) Pause() (err error) {
	event := testUtilsInterfaces.PauseEvent_t{}
	event.SetDisplayName(pBC.dn)
	event.SetState(pBC.state)
	testUtilsGlobals.Gateway.Send(&event)

	return
}

func (pBC *backendControls_t) Play() (err error) {
	event := testUtilsInterfaces.PlayEvent_t{}
	event.SetDisplayName(pBC.dn)
	event.SetState(pBC.state)
	testUtilsGlobals.Gateway.Send(&event)

	return
}

func (pBC *backendControls_t) PlayStation(station uint) (err error) {
	event := testUtilsInterfaces.PlayStationEvent_t{
		Station: station,
	}
	event.SetDisplayName(pBC.dn)
	event.SetState(pBC.state)
	testUtilsGlobals.Gateway.Send(&event)

	return
}

func (pBC *backendControls_t) SetVolume(volume float64) (err error) {
	event := testUtilsInterfaces.SetVolumeEvent_t{
		Volume: volume,
	}
	event.SetDisplayName(pBC.dn)
	event.SetState(pBC.state)
	testUtilsGlobals.Gateway.Send(&event)

	return
}

func (pBC *backendControls_t) SkipNext() (err error) {
	event := testUtilsInterfaces.SkipNextEvent_t{}
	event.SetDisplayName(pBC.dn)
	event.SetState(pBC.state)
	testUtilsGlobals.Gateway.Send(&event)

	return
}

func (pBC *backendControls_t) SkipPrevious() (err error) {
	event := testUtilsInterfaces.SkipPreviousEvent_t{}
	event.SetDisplayName(pBC.dn)
	event.SetState(pBC.state)
	testUtilsGlobals.Gateway.Send(&event)

	return
}

func (pBC *backendControls_t) Stop() (err error) {
	event := testUtilsInterfaces.StopEvent_t{}
	event.SetDisplayName(pBC.dn)
	event.SetState(pBC.state)
	testUtilsGlobals.Gateway.Send(&event)

	return
}

func (pBC *backendControls_t) UpdateAllControls() (err error) {
	event := testUtilsInterfaces.UpdateAllControlsEvent_t{}
	event.SetDisplayName(pBC.dn)
	event.SetState(pBC.state)
	testUtilsGlobals.Gateway.Send(&event)

	pBC.pControlsState.Volume.Set(testUtilsData.InitialVolume)
	pBC.pControlsState.Stations.Set(testUtilsData.InitialStations[pBC.dn])

	return
}
