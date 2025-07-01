package mpd

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (pPBD *passwordBackingData_t) SetupView() (fyne.CanvasObject, error, error) {
	return pPBD.SetupErrlesslyView(), nil, nil
}

func (pPBD *passwordBackingData_t) SetupErrlesslyView() fyne.CanvasObject {
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.Bind(pPBD.password)
	listener := func() {
		pPBD.isPassword.Set(false)
	}
	content := container.NewHBox(
		passwordEntry,
		widget.NewToolbar(
			widget.NewToolbarAction(
				theme.CancelIcon(),
				listener,
			),
		),
	)

	return content
}
