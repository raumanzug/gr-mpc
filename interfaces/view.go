package interfaces

import (
	"fyne.io/fyne/v2"
)

// MView represents a type which is associated by generating ui view.
type MView interface {

	// SetupView generates the associated ui view.
	//
	// the second error result denotes temporary
	// errors as network problems
	SetupView() (fyne.CanvasObject, error, error)
}

// MErrlessView represents a type associated by generating ui view whose
// associated ui view can be generated without throwing errors.
type MErrlessView interface {
	MView

	// SetupErrlesslyView generates this view without throwing errors.
	SetupErrlesslyView() fyne.CanvasObject
}

type MTabEditorView interface {
	MErrlessView

	UpdateSTEF(IServer) error
}
