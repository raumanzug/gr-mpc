package interfaces

import (
	"fyne.io/fyne/v2/data/binding"
)

// ControlsState is used for updating controls ui view.
type ControlsState struct {
	Station  binding.String
	Stations binding.StringList
	Title    binding.String
	Volume   binding.Float
}

// InitControlsState initializes a ControlsState object.
//
// Apply this function before you use a ControlsState objectQ
func (pControlsState *ControlsState) InitControlsState() {
	pControlsState.Station = binding.NewString()
	pControlsState.Stations = binding.NewStringList()
	pControlsState.Title = binding.NewString()
	pControlsState.Volume = binding.NewFloat()
}

type IControlsStateEquipped interface {
	GetControlsState() *ControlsState
}

func (ce ControlsState) GetControlsState() *ControlsState {
	return &ce
}
