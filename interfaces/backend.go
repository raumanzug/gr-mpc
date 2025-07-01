package interfaces

// IBackend represents Backend service's methods.
type IBackend interface {
	IBase

	// GetProtocol returns IProtocol object related
	// to protocol, e.g. "mpd".
	GetProtocol(string) (IProtocol, error)

	// Unmarshal transform json data representing connection data of
	// a media rendering engine into an IServer instance that
	// represents the same connection data.
	Unmarshal(in []byte) (IServer, error)
}

// IProtocol represents methods related to a backend.
type IProtocol interface {
	IBase

	CreateSTEF() ISTEF
}

// IBase represents all that IBackend and IProtocol has in common.
type IBase interface {

	// Map2Server returns IServer object from data gained from
	// json representation.
	Map2Server(map[string]interface{}) (IServer, error)
}
