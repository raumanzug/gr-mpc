package interfaces

import (
	"strconv"
	"strings"

	"github.com/raumanzug/gr-mpc/interfaces"
)

type IEvent interface {
	GetDisplayName() string
	GetState() interfaces.IState
	GetType() string
}

func IEvent2String(event IEvent) string {
	return event.GetDisplayName() +
		": " +
		event.GetType()
}

func IEventCompare(x, y IEvent) int {
	return strings.Compare(
		IEvent2String(x),
		IEvent2String(y),
	)
}

type statebase_t struct {
	dn   string
	conn interfaces.IState
}

func (pB *statebase_t) GetDisplayName() string {
	return pB.dn
}

func (pB *statebase_t) GetState() interfaces.IState {
	return pB.conn
}

func (pB *statebase_t) SetDisplayName(dn string) {
	pB.dn = dn
}

func (pB *statebase_t) SetState(state interfaces.IState) {
	pB.conn = state
}

type OpenBackendControlsEvent_t struct {
	statebase_t
}

func (event OpenBackendControlsEvent_t) GetType() string {
	return "open backend controls"
}

type CloseBackendControlsEvent_t struct {
	statebase_t
}

func (event CloseBackendControlsEvent_t) GetType() string {
	return "close backend controls"
}

type CloseEvent_t struct {
	statebase_t
}

func (event CloseEvent_t) GetType() string {
	return "close"
}

type PauseEvent_t struct {
	statebase_t
}

func (event PauseEvent_t) GetType() string {
	return "pause"
}

type PlayEvent_t struct {
	statebase_t
}

func (event PlayEvent_t) GetType() string {
	return "play"
}

type PlayStationEvent_t struct {
	statebase_t

	Station uint
}

func (event PlayStationEvent_t) GetType() string {
	return "play station nr. " +
		strconv.FormatUint(uint64(event.Station), 10)
}

type SkipNextEvent_t struct {
	statebase_t
}

func (event SkipNextEvent_t) GetType() string {
	return "skip next"
}

type SkipPreviousEvent_t struct {
	statebase_t
}

func (event SkipPreviousEvent_t) GetType() string {
	return "skip previous"
}

type StopEvent_t struct {
	statebase_t
}

func (event StopEvent_t) GetType() string {
	return "stop playback"
}

type SetVolumeEvent_t struct {
	statebase_t

	Volume float64
}

func (event SetVolumeEvent_t) GetType() string {
	return "set volume to " +
		strconv.FormatFloat(event.Volume, 'f', 5, 64)
}

type UpdateAllControlsEvent_t struct {
	statebase_t
}

func (event UpdateAllControlsEvent_t) GetType() string {
	return "update all controls"
}

type LifecycleStartEvent_t struct {
	statebase_t
}

func (event LifecycleStartEvent_t) GetType() string {
	return "lifecycle start"
}

type LifecycleStopEvent_t struct {
	statebase_t
}

func (event LifecycleStopEvent_t) GetType() string {
	return "lifecycle stop"
}
