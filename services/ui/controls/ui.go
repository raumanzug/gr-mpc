// Package controls contains a generator of a view consisting in a
// label indicating the name of station currently used and control elements
// that can be used to control the media rendering engine currently used, e.g.
// pause, start, stop, station selection, volume.
package controls

import (
	"errors"

	"github.com/raumanzug/gr-mpc/constants"
	"github.com/raumanzug/gr-mpc/globals"
	"github.com/raumanzug/gr-mpc/interfaces"
	"github.com/raumanzug/gr-mpc/utils/dbj"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/lang"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (pBD *BackingData) stationListCallback(id widget.ListItemID) func() error {
	return func() (err error) {
		if pBD.BackendState == nil {
			return
		}

		var controls interfaces.IBackendControls
		controls, err = pBD.BackendState.OpenBackendControls()
		if err != nil {
			return
		}
		defer func() {
			err = errors.Join(controls.Close())
		}()

		err = controls.PlayStation(uint(id))

		return
	}
}

func (pBD *BackingData) mediaPauseCallback() (err error) {
	if pBD.BackendState == nil {
		return
	}

	var controls interfaces.IBackendControls
	controls, err = pBD.BackendState.OpenBackendControls()
	if err != nil {
		return
	}
	defer func() {
		err = errors.Join(err, controls.Close())
	}()

	err = controls.Pause()

	return
}

func (pBD *BackingData) mediaPlayCallback() (err error) {
	if pBD.BackendState == nil {
		return
	}

	var controls interfaces.IBackendControls
	controls, err = pBD.BackendState.OpenBackendControls()
	if err != nil {
		return
	}
	defer func() {
		err = errors.Join(err, controls.Close())
	}()

	err = controls.Play()

	return
}

func (pBD *BackingData) mediaSkipNextCallback() (err error) {
	if pBD.BackendState == nil {
		return
	}

	var controls interfaces.IBackendControls
	controls, err = pBD.BackendState.OpenBackendControls()
	if err != nil {
		return
	}
	defer func() {
		err = errors.Join(err, controls.Close())
	}()

	err = controls.SkipNext()

	return
}

func (pBD *BackingData) mediaSkipPreviousCallback() (err error) {
	if pBD.BackendState == nil {
		return
	}

	var controls interfaces.IBackendControls
	controls, err = pBD.BackendState.OpenBackendControls()
	if err != nil {
		return
	}
	defer func() {
		err = errors.Join(err, controls.Close())
	}()

	err = controls.SkipPrevious()

	return
}

func (pBD *BackingData) mediaStopCallback() (err error) {
	if pBD.BackendState == nil {
		return
	}

	var controls interfaces.IBackendControls
	controls, err = pBD.BackendState.OpenBackendControls()
	if err != nil {
		return
	}
	defer func() {
		err = errors.Join(err, controls.Close())
	}()

	err = controls.Stop()

	return
}

func (pBD *BackingData) getVolumeSliderCallback(
	PUiState *interfaces.ControlsState,
) func() error {
	return func() (err error) {
		if pBD.BackendState == nil {
			return
		}

		var controls interfaces.IBackendControls
		controls, err = pBD.BackendState.OpenBackendControls()
		if err != nil {
			return
		}
		defer func() {
			err = errors.Join(err, controls.Close())
		}()

		volume, cerr := PUiState.Volume.Get()
		if cerr != nil {
			err = errors.Join(err, cerr)
			return
		}
		err = errors.Join(
			err,
			controls.SetVolume(volume),
		)

		return
	}
}

func (pBD *BackingData) RegisterVolumeSliderCallback() {
	volumeSliderListener := binding.NewDataListener(
		dbj.DoBackgroundJob(
			pBD.BackendState,
			pBD.getVolumeSliderCallback(pBD.PUiState),
		),
	)
	pBD.PUiState.Volume.AddListener(volumeSliderListener)
}

func (pBD *BackingData) SetupView() (content fyne.CanvasObject, err error, terr error) {
	var controlGroups = []fyne.CanvasObject{}

	if pBD.PUiState.Station != nil {
		stationLabel := widget.NewLabelWithData(pBD.PUiState.Station)
		stationLabel.Wrapping = fyne.TextWrapWord
		controlGroups = append(controlGroups, stationLabel)
	}

	if pBD.PUiState.Station != nil && pBD.PUiState.Title != nil {
		titleLabel := widget.NewLabelWithData(pBD.PUiState.Title)
		titleLabelSize := titleLabel.MinSize()
		titleLabel.Wrapping = fyne.TextWrapWord
		titleItem := widget.NewAccordionItem(
			lang.X(
				"program",
				"Now playing",
			),
			titleLabel,
		)

		titleLabelH, titleLabelV := titleLabelSize.Components()
		titleLabelV *= constants.StationSelectorResizeFactor
		stationListSize := fyne.NewSize(titleLabelH, titleLabelV)
		backgroundWidget := canvas.NewRectangle(
			constants.StationSelectorBackgroundColor,
		)
		backgroundWidget.SetMinSize(stationListSize)
		stationList := widget.NewListWithData(
			pBD.PUiState.Stations,
			func() fyne.CanvasObject {
				return widget.NewLabel("template")
			},
			func(d binding.DataItem, o fyne.CanvasObject) {
				dString, ok := d.(binding.String)
				var err error
				if !ok {
					err = errors.New("data item does not bind to string.")
				}
				o.(*widget.Label).Bind(dString)
				globals.ApplicationContext.UI().AddErr(err)
			})
		stationList.OnSelected = func(id widget.ListItemID) {
			dbj.DoBackgroundJob(
				pBD.BackendState,
				pBD.stationListCallback(id),
			)()
		}
		stationListStack := container.NewStack(
			backgroundWidget,
			stationList,
		)
		stationListItem := widget.NewAccordionItem(
			lang.X("stations", "Stations"),
			stationListStack,
		)

		accordion := widget.NewAccordion(
			titleItem,
			stationListItem,
		)
		accordion.MultiOpen = false
		controlGroups = append(controlGroups, accordion)
	}

	mediaPause := widget.NewToolbarAction(
		theme.MediaPauseIcon(),
		dbj.DoBackgroundJob(
			pBD.BackendState,
			pBD.mediaPauseCallback,
		),
	)

	mediaPlay := widget.NewToolbarAction(
		theme.MediaPlayIcon(),
		dbj.DoBackgroundJob(
			pBD.BackendState,
			pBD.mediaPlayCallback,
		),
	)

	mediaSkipNext := widget.NewToolbarAction(
		theme.MediaSkipNextIcon(),
		dbj.DoBackgroundJob(
			pBD.BackendState,
			pBD.mediaSkipNextCallback,
		),
	)

	mediaSkipPrevious := widget.NewToolbarAction(
		theme.MediaSkipPreviousIcon(),
		dbj.DoBackgroundJob(
			pBD.BackendState,
			pBD.mediaSkipPreviousCallback,
		),
	)

	mediaStop := widget.NewToolbarAction(
		theme.MediaStopIcon(),
		dbj.DoBackgroundJob(
			pBD.BackendState,
			pBD.mediaStopCallback,
		),
	)

	naviToolbar := container.NewCenter(
		widget.NewToolbar(
			mediaSkipPrevious,
			mediaSkipNext,
		),
	)
	controlGroups = append(controlGroups, naviToolbar)

	if pBD.PUiState.Volume != nil {
		volumeSlider := widget.NewSlider(0.0, 1.0)
		volumeSlider.Step = 0.01
		volumeSlider.Bind(pBD.PUiState.Volume)
		controlGroups = append(controlGroups, volumeSlider)
	}

	playToolbar := container.NewCenter(
		widget.NewToolbar(
			mediaPlay,
			mediaPause,
			mediaStop,
		),
	)

	controlGroups = append(controlGroups, playToolbar)

	content = container.NewVBox(controlGroups...)

	return
}
