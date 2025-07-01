package dbj

import (
	"errors"

	"github.com/raumanzug/gr-mpc/globals"
	"github.com/raumanzug/gr-mpc/interfaces"

	"fyne.io/fyne/v2"
)

func _doBackgroundJob(
	state interfaces.IMutexEquipped,
	job func() error,
) {
	pMutex := state.GetMutex()

	go func() {
		var err error
		defer func() {
			fyne.Do(
				func() {
					globals.ApplicationContext.UI().AddErr(err)
				},
			)
		}()

		if !pMutex.TryLock() {
			err = errors.New("backend blocked; request ignored.")
			return
		}
		defer pMutex.Unlock()

		err = job()
	}()

}

func _forceBackgroundJob(
	state interfaces.IMutexEquipped,
	job func() error,
) {
	pMutex := state.GetMutex()

	go func() {
		var err error
		defer func() {
			fyne.Do(
				func() {
					globals.ApplicationContext.UI().AddErr(err)
				},
			)
		}()

		pMutex.Lock()
		defer pMutex.Unlock()

		err = job()
	}()

}

// DoBackgroundJob performs a job in an own thread.  During the
// lifetime of this job a mutex prevents starting another job
// which is assigned to the same mutex.  Instead this another job
// is simply dropping.
func DoBackgroundJob(
	state interfaces.IMutexEquipped,
	job func() error,
) func() {
	return func() {
		_doBackgroundJob(state, job)
	}
}

// ForceBackgroundJob performs a job in an own thread.  During the
// lifetime of this job a mutex prevents starting another job
// which is assigned to the same mutex.  Instead this another job
// will waiting til this mutex is free.
func ForceBackgroundJob(
	state interfaces.IMutexEquipped,
	job func() error,
) func() {
	return func() {
		_forceBackgroundJob(state, job)
	}
}
