package interfaces

import (
	"strings"
)

// IServer represents a media rendering engine.
type IServer interface {

	// GetDisplayName returns the display name of this engine.
	GetDisplayName() string

	// GetProtocolId returns the ProtocolId from type of
	// media rendering engine, e.g. "mpd" for MPD servers.
	GetProtocolId() string

	// Open turns a connection to this media rendering engine
	// into operation.
	Open() (IState, error)

	// MarshalJSON serializes an object of this type into json.
	MarshalJSON() ([]byte, error)

	UpdateSTEF(stef ISTEF) error
}

// ServerCompare is used for sorting media rendering engines.
// ServerCompare behaves similar to strcmp in C for strings.
func ServerCompare(x, y IServer) int {
	return strings.Compare(
		x.GetDisplayName(),
		y.GetDisplayName(),
	)
}

// IState represents operation on media rendering engine.
type IState interface {
	IControlsStateEquipped
	IMutexEquipped

	// OpenBackendControls opens a connection related to
	// this media rendering engine.
	OpenBackendControls() (IBackendControls, error)

	// Close ends the work with this media rendering engine.
	Close() error

	// Lifecycle methods.

	// Start can be called whenever this object can be
	// safely used because energy management of
	// smartphone operating systems set this app active.
	Start() error

	// Stop prepares this object for inactive state set
	// by energy management of smartphone operating systems.
	Stop() error
}
