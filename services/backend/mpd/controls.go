package mpd

import (
	"errors"
	"math"
	"strconv"

	"github.com/raumanzug/gr-mpc/globals"
	"github.com/raumanzug/gr-mpc/interfaces"

	"fyne.io/fyne/v2"
	"github.com/fhs/gompd/v2/mpd"
)

type mpdControls_t struct {
	pMpdClient     *mpd.Client
	pControlsState *interfaces.ControlsState
}

func (pC *mpdControls_t) Close() (err error) {
	err = pC.pMpdClient.Close()

	return
}

func (pC *mpdControls_t) Pause() (err error) {
	err = pC.pMpdClient.Pause(true)

	return
}

func (pC *mpdControls_t) Play() (err error) {
	err = pC.pMpdClient.Play(-1)

	return
}

func (pC *mpdControls_t) PlayStation(nr uint) (err error) {
	err = pC.pMpdClient.Play(int(nr))

	return
}

func (pC *mpdControls_t) SkipNext() (err error) {
	err = pC.pMpdClient.Play(-1)
	if err != nil {
		return
	}
	err = pC.pMpdClient.Next()

	return
}

func (pC *mpdControls_t) SkipPrevious() (err error) {
	err = pC.pMpdClient.Play(-1)
	if err != nil {
		return
	}
	err = pC.pMpdClient.Previous()

	return
}

func (pC *mpdControls_t) Stop() (err error) {
	err = pC.pMpdClient.Stop()

	return
}

func (pC *mpdControls_t) SetVolume(volume float64) (err error) {
	volInt := int(math.Round(volume * 100.0))
	err = pC.pMpdClient.SetVolume(volInt)

	return
}

func (pC *mpdControls_t) UpdateCurrentSong() (err error) {
	var attrs mpd.Attrs
	attrs, err = pC.pMpdClient.CurrentSong()
	if err != nil {
		return
	}

	{
		name, nonvoid := attrs["Name"]
		if !nonvoid {
			name = "???"
		}
		fyne.Do(
			func() {
				err := pC.pControlsState.Station.Set(name)
				globals.ApplicationContext.UI().AddErr(err)
			},
		)
	}
	{
		title, nonvoid := attrs["Title"]
		if !nonvoid {
			title = "???"
		}
		fyne.Do(
			func() {
				err := pC.pControlsState.Title.Set(title)
				globals.ApplicationContext.UI().AddErr(err)
			},
		)
	}

	return
}

func (pC *mpdControls_t) UpdateSongList() (err error) {
	var attrsl []mpd.Attrs
	attrsl, err = pC.pMpdClient.PlaylistInfo(-1, -1)
	if err != nil {
		return
	}

	fyne.Do(
		func() {
			emptyList := make([]string, len(attrsl))
			err := pC.pControlsState.Stations.Set(emptyList)
			defer func() {
				globals.ApplicationContext.UI().AddErr(err)
			}()
			if err != nil {
				return
			}
			for _, attrs := range attrsl {
				posString, nonvoid := attrs["Pos"]
				if !nonvoid {
					continue
				}
				pos, cerr := strconv.Atoi(posString)
				if cerr != nil {
					err = errors.Join(err, cerr)
					continue
				}
				if len(attrsl) <= pos {
					cerr := errors.New("pos behind buffer lgth.")
					err = errors.Join(err, cerr)
					continue
				}
				stationName, nonvoid := attrs["Name"]
				if !nonvoid {
					continue
				}
				err = errors.Join(
					err,
					pC.pControlsState.
						Stations.
						SetValue(pos, stationName),
				)
			}
		},
	)

	return
}

func (pC *mpdControls_t) updateMpdStatus() (err error) {
	var attrs mpd.Attrs
	attrs, err = pC.pMpdClient.Status()
	if err != nil {
		return
	}

	{
		volString, nonvoid := attrs["volume"]
		if !nonvoid {
			return
		}
		volInt, cerr := strconv.Atoi(volString)
		if cerr != nil {
			err = errors.Join(err, cerr)
			return
		}
		volFloat := float64(volInt) / 100.0
		fyne.Do(
			func() {
				err = pC.pControlsState.Volume.Set(volFloat)
				globals.ApplicationContext.UI().AddErr(err)
			},
		)
	}

	return
}

func (pC *mpdControls_t) UpdateVolumeControls() (err error) {
	err = pC.updateMpdStatus()

	return
}

func (pC *mpdControls_t) UpdateAllControls() (err error) {
	err = pC.UpdateCurrentSong()
	err = errors.Join(err, pC.UpdateSongList())
	err = errors.Join(err, pC.updateMpdStatus())

	return
}
