package mpd

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/lang"
	"fyne.io/fyne/v2/widget"
)

func (pPBD *nopwBackingData_t) SetupView() (fyne.CanvasObject, error, error) {
	return pPBD.SetupErrlesslyView(), nil, nil
}

func (pPBD *nopwBackingData_t) SetupErrlesslyView() fyne.CanvasObject {
	listener := func() {
		pPBD.isPassword.Set(true)
	}
	content := widget.NewButton(
		lang.X("enter a password", "Enter a password"),
		listener,
	)

	return content
}
