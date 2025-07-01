package controls

import (
	"github.com/raumanzug/gr-mpc/interfaces"
)

type BackingData struct {
	PUiState     *interfaces.ControlsState
	BackendState interfaces.IState
}

func IState2pBackingData(state interfaces.IState) *BackingData {
	backingData := BackingData{
		PUiState:     state.GetControlsState(),
		BackendState: state,
	}

	return &backingData
}
