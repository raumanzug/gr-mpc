package interfaces

import (
	"github.com/raumanzug/gr-mpc/interfaces"
)

// IBackendControls represents methods for controlling
// MPD servers.
type IBackendControls interface {
	interfaces.IBackendControls

	UpdateCurrentSong() error
	UpdateSongList() error
	UpdateVolumeControls() error
}
