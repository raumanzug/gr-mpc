package servertabs

import (
	"errors"
	"slices"

	"github.com/raumanzug/gr-mpc/globals"
	"github.com/raumanzug/gr-mpc/interfaces"
	"github.com/raumanzug/gr-mpc/services/ui/controls"
	"github.com/raumanzug/gr-mpc/utils/dbj"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func iServer2TabItem(server interfaces.IServer) (
	pTabItem *container.TabItem,
	err error,
	terr error,
) {
	var state interfaces.IState
	state, err = globals.ApplicationContext.ServerTabs().Activate(server)

	var view fyne.CanvasObject
	if err == nil {
		pStack := container.NewStack(
			widget.NewProgressBarInfinite(),
		)
		var controlsView fyne.CanvasObject
		pControlsMView := controls.IState2pBackingData(state)
		controlsView, err, terr = pControlsMView.SetupView()
		if err == nil {
			view = pStack

			dbj.DoBackgroundJob(
				state,
				updateControls(
					state,
					pStack,
					terr,
					controlsView,
					pControlsMView,
				),
			)()
		}
	}
	if err != nil {
		view = widget.NewIcon(theme.BrokenImageIcon())
	}

	pTabItem = container.NewTabItem(
		server.GetDisplayName(),
		view,
	)

	return
}

func closeCallback(pTabItem *container.TabItem) {
	err := globals.ApplicationContext.ServerTabs().
		RemoveServerTabCallback(pTabItem.Text)
	globals.ApplicationContext.UI().AddErr(err)
}

func selectCallback(pTabItem *container.TabItem) {
	err := globals.ApplicationContext.ServerTabs().
		SelectCallback(pTabItem.Text)
	globals.ApplicationContext.UI().AddErr(err)
}

func updateControls(
	state interfaces.IState,
	pStack *fyne.Container,
	terr error,
	controlsView fyne.CanvasObject,
	pControlsMView *controls.BackingData,
) func() error {
	return func() (err error) {
		var backendControls interfaces.IBackendControls
		backendControls, err = state.OpenBackendControls()
		if err != nil {
			return
		}
		defer func() {
			err = errors.Join(err, backendControls.Close())
		}()

		var tcerr error
		tcerr = backendControls.UpdateAllControls()
		terr = errors.Join(terr, tcerr)
		err = errors.Join(err, terr)
		fyne.Do(
			func() {
				pControlsMView.RegisterVolumeSliderCallback()
				pStack.RemoveAll()
				pStack.Add(controlsView)
			},
		)

		return
	}
}

func (pBD *BackingData) SetupView() (content fyne.CanvasObject, err error, terr error) {

	tabItemGenerator := func(yield func(*container.TabItem) bool) {
		for _, server := range pBD.servers {
			pTabItem, cerr, tcerr := iServer2TabItem(server)
			err = errors.Join(err, cerr)
			terr = errors.Join(terr, tcerr)
			if !yield(pTabItem) {
				return
			}
		}
	}
	tabItems := slices.Collect(tabItemGenerator)
	pBD.docTabs = container.NewDocTabs(tabItems...)
	pBD.docTabs.OnClosed = closeCallback
	pBD.docTabs.OnSelected = selectCallback

	content = pBD.docTabs

	return
}
