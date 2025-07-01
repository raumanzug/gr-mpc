// Package workpad contains the entire front page of this app except
// the part intended for displaying error messages.
package workpad

import (
	"github.com/raumanzug/gr-mpc/globals"
	"github.com/raumanzug/gr-mpc/services/ui/tabedit"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (sbd BackingData) SetupView() (fyne.CanvasObject, error, error) {
	return sbd.SetupErrlesslyView(), nil, nil
}

func (sbd BackingData) SetupErrlesslyView() (content fyne.CanvasObject) {

	newAction := func() {
		bd := tabedit.BackingData{
			PExc: nil,
		}
		view := bd.SetupErrlesslyView()
		globals.ApplicationContext.UI().PushView(view)
	}
	serverNew := widget.NewToolbarAction(
		theme.ContentAddIcon(),
		newAction,
	)

	editAction := func() {
		pExc := globals.ApplicationContext.Preferences().SelectedServerTabKey()
		bd := tabedit.BackingData{
			PExc: pExc,
		}
		view := bd.SetupErrlesslyView()
		if pExc != nil {
			serverTab, err := globals.ApplicationContext.Preferences().LookupServerTab(*pExc)
			if err != nil {
				globals.ApplicationContext.UI().AddErr(err)
				return
			}
			err = bd.UpdateSTEF(serverTab)
			if err != nil {
				globals.ApplicationContext.UI().AddErr(err)
				return
			}
		}
		globals.ApplicationContext.UI().PushView(view)
	}
	serverEdit := widget.NewToolbarAction(
		theme.SettingsIcon(),
		editAction,
	)

	srvToolbar := widget.NewToolbar(
		serverNew,
		serverEdit,
	)

	srvToolbarCentered := container.NewCenter(srvToolbar)

	ct := globals.ApplicationContext.UI().GetDocTabs()

	content = container.NewVBox(
		srvToolbarCentered,
		ct,
	)

	return
}
