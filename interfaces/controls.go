package interfaces

// IBackendControls represents methods for controlling
// media rendering engines.
type IBackendControls interface {
	Pause() error
	Play() error
	PlayStation(uint) error
	SkipNext() error
	SkipPrevious() error
	Stop() error

	// SetVolume sets the volume; range [0.0, 1.0]
	SetVolume(float64) error

	// UpdateAllControls set up initial data to control view
	UpdateAllControls() error

	// Close indicate that we stop working with this engine
	Close() error
}
