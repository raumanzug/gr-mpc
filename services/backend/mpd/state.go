package mpd

import (
	"errors"
	"fmt"
	"sync"

	"github.com/raumanzug/gr-mpc/globals"
	"github.com/raumanzug/gr-mpc/interfaces"
	mpdInterfaces "github.com/raumanzug/gr-mpc/services/backend/mpd/interfaces"
	"github.com/raumanzug/gr-mpc/utils/dbj"

	"fyne.io/fyne/v2"
	"github.com/fhs/gompd/v2/mpd"
)

type state_t struct {
	interfaces.ControlsState
	interfaces.MutexEquippedBase_t

	server    server_t
	isOpen    bool
	waitGroup sync.WaitGroup
	pWatcher  *mpd.Watcher
}

func NewState(server server_t) interfaces.IState {
	state := state_t{
		server: server,
		isOpen: false,
	}

	state.InitControlsState()

	return &state
}

func (pState *state_t) updateSubsystem(subsystem string) func() error {
	return func() (err error) {
		var controls mpdInterfaces.IBackendControls
		controls, err = pState.OpenMpdBackendControls()
		if err != nil {
			return
		}
		defer func() {
			err = errors.Join(err, controls.Close())
		}()
		switch subsystem {
		case "mixer":
			err = errors.Join(
				err,
				controls.UpdateVolumeControls(),
			)
		case "player":
			err = errors.Join(
				err,
				controls.UpdateCurrentSong(),
			)
		case "playlist":
			err = errors.Join(
				err,
				controls.UpdateSongList(),
			)
		default:
			cerr := errors.New(
				"subsystem not supported",
			)
			err = errors.Join(err, cerr)
		}

		return
	}
}

func (pState *state_t) preStart() (err error) {
	var password string
	if pState.server.password != nil {
		password = *pState.server.password
	}
	pState.pWatcher, err = mpd.NewWatcher(
		"tcp",
		pState.server.socket,
		password,
		"mixer",
		"player",
		"playlist",
	)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			err = errors.Join(err, pState.pWatcher.Close())
		}
	}()

	var backendControls interfaces.IBackendControls
	backendControls, err = pState.OpenBackendControls()
	if err != nil {
		return
	}
	defer func() {
		err = errors.Join(err, backendControls.Close())
	}()

	err = backendControls.UpdateAllControls()
	if err != nil {
		return
	}

	return
}

func (pState *state_t) Start() (err error) {
	if pState.isOpen {
		return
	}

	err = pState.preStart()
	if err != nil {
		return
	}

	pState.isOpen = true
	pState.waitGroup.Add(2)

	go func() {
		defer pState.waitGroup.Done()
		var prevErr error
		var counter uint = 0
		for err := range pState.pWatcher.Error {
			if prevErr != nil && prevErr == err {
				counter++
				continue
			}
			prevErr = err
			counter++
			if counter > 1 {
				errMsg := fmt.Sprintf("...%d times...", counter)
				err = errors.Join(err, errors.New(errMsg))
			}
			counter = 0
			fyne.Do(
				func() {
					globals.ApplicationContext.UI().AddErr(err)
				},
			)
		}
	}()
	go func() {
		defer pState.waitGroup.Done()
		for subsystem := range pState.pWatcher.Event {
			dbj.DoBackgroundJob(
				pState,
				pState.updateSubsystem(subsystem),
			)()
		}
	}()

	return
}

func (pState *state_t) Stop() (err error) {
	if !pState.isOpen {
		return
	}
	if pState.pWatcher == nil {
		err = errors.New("stop denied; no watcher active")
		return
	}
	err = pState.pWatcher.Close()
	pState.isOpen = false
	pState.waitGroup.Wait()

	return
}

func (pState *state_t) OpenBackendControls() (
	controls interfaces.IBackendControls,
	err error,
) {
	controls, err = pState.OpenMpdBackendControls()

	return
}

func (pState *state_t) OpenMpdBackendControls() (
	controls mpdInterfaces.IBackendControls,
	err error,
) {
	pControls := &mpdControls_t{}

	if pState.server.password == nil {
		pControls.pMpdClient, err = mpd.Dial(
			"tcp",
			pState.server.socket,
		)
	} else {
		pControls.pMpdClient, err = mpd.DialAuthenticated(
			"tcp",
			pState.server.socket,
			*pState.server.password,
		)
	}
	pControls.pControlsState = pState.GetControlsState()

	controls = pControls

	return
}

func (pState *state_t) Close() (err error) {

	return
}
