// Package mpd contains a generator of a view in which connection data
// of a MPD server can be edited/entered.
package tetab

import (
	"errors"

	"github.com/raumanzug/gr-mpc/globals"
	"github.com/raumanzug/gr-mpc/interfaces"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/lang"
	"fyne.io/fyne/v2/widget"
)

func (pBD *BackingData) SetupView() (fyne.CanvasObject, error, error) {
	return pBD.SetupErrlesslyView(), nil, nil
}

func (pBD *BackingData) SetupErrlesslyView() (content fyne.CanvasObject) {
	dnRawEntry := widget.NewEntryWithData(pBD.STEF.DisplayNameBinding())
	dnEntry := widget.NewFormItem(
		lang.X("display name", "display name"),
		dnRawEntry,
	)
	dnRawEntry.Validator = globals.ApplicationContext.Preferences().DoesNotExist(pBD.PExc)
	pForm := widget.NewForm(
		dnEntry,
	)
	if !globals.ApplicationContext.UI().BottomReached() {
		pForm.OnCancel = func() {
			err := globals.ApplicationContext.UI().CloseTopmostView()
			globals.ApplicationContext.UI().AddErr(err)
		}
	}
	pForm.OnSubmit = func() {
		var err error
		defer func() {
			err = errors.Join(
				err,
				globals.ApplicationContext.UI().UpdateScreen(),
			)
			globals.ApplicationContext.UI().AddErr(err)
		}()

		if pBD.PExc != nil {
			err = globals.ApplicationContext.UI().RemoveServerTab(*pBD.PExc)
			if err != nil {
				return
			}
		}

		var dnServer interfaces.IServer
		dnServer, err = pBD.STEF.Edited()
		if err != nil {
			return
		}

		var terr error
		terr, err = globals.ApplicationContext.UI().AddServerTab(dnServer)
		if err != nil {
			return
		}
		globals.ApplicationContext.UI().AddErr(terr)
	}
	pForm.SubmitText = lang.X("submit form", "submit")
	pForm.CancelText = lang.X("cancel form", "cancel")
	pBD.STEF.UpdateForm(pForm)

	content = pForm

	return
}
