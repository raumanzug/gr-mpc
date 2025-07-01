package lifecycle

import (
	"errors"

	"github.com/raumanzug/gr-mpc/interfaces"
	"github.com/raumanzug/gr-mpc/utils/dbj"

	"fyne.io/fyne/v2"
)

type lifecycle_t struct {
	context interfaces.IApplicationContext

	lifecycle fyne.Lifecycle

	isActive bool
}

func (pLc *lifecycle_t) start() (err error) {
	pLc.isActive = true
	var pServerKey *string
	pServerKey = pLc.context.Preferences().SelectedServerTabKey()
	if pServerKey == nil {
		return
	}
	state := pLc.context.ServerTabs().GetActivatedState(*pServerKey)
	if state == nil {
		err = errors.New("Start called before Open")
		return
	}
	dbj.DoBackgroundJob(
		state,
		state.Start,
	)()

	return
}

func (pLc *lifecycle_t) stop() (err error) {
	pLc.isActive = false
	for _, state := range pLc.context.ServerTabs().GetActivatedStates() {
		dbj.DoBackgroundJob(
			state,
			state.Stop,
		)()
	}

	return
}

func (pLc *lifecycle_t) Configure() {
	pLc.lifecycle.SetOnStarted(
		func() {
			err := pLc.start()
			pLc.context.UI().AddErr(err)
		},
	)
	pLc.lifecycle.SetOnStopped(
		func() {
			err := pLc.stop()
			pLc.context.UI().AddErr(err)
		},
	)

	return
}

func (pLc *lifecycle_t) IsActive() (isActive bool) {
	isActive = pLc.isActive

	return
}

func New(
	context interfaces.IApplicationContext,
	lc fyne.Lifecycle,
) interfaces.ILifecycle {
	lifecycle := lifecycle_t{
		context:   context,
		lifecycle: lc,
		isActive:  false,
	}

	return &lifecycle
}
