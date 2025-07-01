package interfaces

type IServerTabs interface {
	// Activate sets up a server in the app.
	Activate(IServer) (IState, error)

	// GetActivatedState returns the state with display name given
	// in the arg if there is such a state that is open.
	GetActivatedState(string) IState

	// GetActivatedState returns open states.
	GetActivatedStates() []IState

	// RemoveServerTabCallback acts if a doc tab is erased.
	//
	// does not affect doc tabs view.
	// instead it is callback for close hook.
	RemoveServerTabCallback(string) error

	// SelectCallback acts if a doc tab making the relating
	// media rendering engine the media rendering engine currently used.
	//
	// does not affect doc tabs view.
	// instead it is callback for onSelected hook in doc tabs container.
	SelectCallback(string) error
}
