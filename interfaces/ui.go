package interfaces

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// IUI represents UI service's methods.
type IUI interface {

	// AddErr adds an error to error showing widget.
	//
	// Call it in fyne's thread only!
	AddErr(error)

	// AddServerTab adds a server tab to doc tabs.
	//
	// Call it in fyne's thread only!
	// It does not check whether another tab with
	// same display name already exists.
	// The second error denote temporary errors
	// as network problems.
	AddServerTab(server IServer) (error, error)

	// BottomReached checks whether view stack has a top most
	// object which could be closed and deleted.
	BottomReached() bool

	// CloseTopmostView drops the top most object from view stack
	// and replaces the screen by the view located next to this
	// view.
	CloseTopmostView() error

	// Configure initializes this service.  Call it before
	// you use this service!
	Configure()

	// GetApp returns the representation of this app.
	GetApp() fyne.App

	// GetDocTabs returns the doc tabs widget consisting in
	// a doc tab for each media rendering engine.
	GetDocTabs() *container.DocTabs

	// GetMainWindow returns the outmost window of this app.
	GetMainWindow() fyne.Window

	// PushView pushes a view on top of the view stack.
	// If view stack overcedes the capacity an error arises.
	PushView(fyne.CanvasObject) error

	// RemoveServerTab removes doc tab.
	RemoveServerTab(string) error

	// Reset the front page of the smartphone.
	// The view stack contains only the front page after
	// calling this procedure.
	//
	// Call it in fyne's thread only!
	UpdateScreen() error

	// UpdateRoot erases error box and replaces all the
	// errors in it by the error in the parameter.
	// If error is nil the error box wanes from smartphone's screen.
	//
	// Call it in fyne's thread only!
	UpdateRoot(error)
}
