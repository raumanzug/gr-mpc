// Package application implements the application context.
package application

import (
	"github.com/raumanzug/gr-mpc/interfaces"
	"github.com/raumanzug/gr-mpc/services/backend"
	"github.com/raumanzug/gr-mpc/services/backend/mpd"
	"github.com/raumanzug/gr-mpc/services/lifecycle"
	"github.com/raumanzug/gr-mpc/services/preferences"
	"github.com/raumanzug/gr-mpc/services/servertabs"
	"github.com/raumanzug/gr-mpc/services/ui"

	"fyne.io/fyne/v2/app"
)

type context_t struct {
	backend     interfaces.IBackend
	lifecycle   interfaces.ILifecycle
	preferences interfaces.IPreferences
	servertabs  interfaces.IServerTabs
	ui          interfaces.IUI
}

// New initializes all the services.
func New() interfaces.IApplicationContext {
	a := app.New()
	l := a.Lifecycle()
	p := a.Preferences()

	context := context_t{}

	backends := map[string]interfaces.IProtocol{
		"mpd": mpd.New(&context),
	}

	context.backend = backend.New(&context, backends)
	context.lifecycle = lifecycle.New(&context, l)
	context.preferences = preferences.New(&context, p)
	context.servertabs = servertabs.New(&context)
	context.ui = ui.New(&context, a)

	return &context
}

func (pAc *context_t) Backend() interfaces.IBackend {
	return pAc.backend
}

func (pAc *context_t) Lifecycle() interfaces.ILifecycle {
	return pAc.lifecycle
}

func (pAc *context_t) Preferences() interfaces.IPreferences {
	return pAc.preferences
}

func (pAc *context_t) ServerTabs() interfaces.IServerTabs {
	return pAc.servertabs
}

func (pAc *context_t) UI() interfaces.IUI {
	return pAc.ui
}
