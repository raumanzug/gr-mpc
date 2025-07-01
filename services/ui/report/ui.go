// Package report contains a generator for a view that shows error messages.
package report

import (
	"github.com/raumanzug/gr-mpc/constants"
	"github.com/raumanzug/gr-mpc/globals"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (pBD *BackingData) SetupView() (fyne.CanvasObject, error, error) {
	return pBD.SetupErrlesslyView(), nil, nil
}

func (pBD *BackingData) SetupErrlesslyView() (content fyne.CanvasObject) {
	mainWindow := globals.ApplicationContext.UI().GetMainWindow()
	mainWindow.Resize(fyne.Size{255, 255})
	mainWindowSize := mainWindow.Canvas().Size()
	backgroundSizeH, backgroundSizeV := mainWindowSize.Components()
	backgroundSizeH *= 0.62
	backgroundSizeV *= 0.236
	backgroundSize := fyne.Size{backgroundSizeH, backgroundSizeV}
	background := canvas.NewRectangle(constants.ErrorBackgroundColor)
	background.SetMinSize(backgroundSize)

	errorScroll := container.NewScroll(
		pBD.pErrorLabel,
	)

	errorWithBackground := container.NewStack(
		background,
		errorScroll,
	)

	clearObject := widget.NewToolbarAction(
		theme.ContentClearIcon(),
		pBD.Clear,
	)

	clearTB := widget.NewToolbar(
		clearObject,
	)

	content = container.NewHBox(
		widget.NewIcon(theme.ErrorIcon()),
		errorWithBackground,
		clearTB,
	)

	return
}
