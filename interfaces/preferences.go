package interfaces

import (
	"fyne.io/fyne/v2"
)

// IPreference represents Preference service's methods.
type IPreferences interface {
	// AddServTab adds a IServer object, i.e. a representation
	// of a media rendering engine, to the list of known
	// media rendering engines in the preference database.
	//
	// Method does not check whether display name already exists.
	AddServerTab(server IServer) error

	// DoesNotExist is used as form validator for ensuring that
	// display name for media rendering engine is elected
	// unambigously.
	DoesNotExist(*string) fyne.StringValidator

	// LookupServerTab looks up the list of known media rendering
	// engines for a media rendering engine with a given display name.
	// Call results in nil if this method does not found such a
	// media rendering engine.
	LookupServerTab(string) (IServer, error)

	// RemoveServerTab deletes a media rendering engine from the
	// list of known media rendering engines in the preferences
	// database.
	RemoveServerTab(string) error

	// SaveSelectedServerTab registers a media rendering engine
	// as the media rendering engine currently used.
	//
	// Method does not check whether the list of known
	// media rendering engines in the preferences database contains
	// such a media rendering engine.
	SaveSelectedServerTab(string) error

	// SelectedServerTabKey returns nil if there are no known
	// media rendering engines and a pointer to the display name
	// of that server that is selected.
	SelectedServerTabKey() *string

	// TabbedServers returns the list of knonwn media rendering engines
	// in the preferences database.
	TabbedServers() ([]IServer, error)
}
