package backend

import (
	"encoding/json"
	"errors"

	"github.com/raumanzug/gr-mpc/interfaces"
)

type backend_t struct {
	context interfaces.IApplicationContext

	backends map[string]interfaces.IProtocol
}

func (pBackend *backend_t) GetProtocol(id string) (
	protocol interfaces.IProtocol,
	err error,
) {
	var nonvoid bool
	protocol, nonvoid = pBackend.backends[id]
	if !nonvoid {
		err = errors.New("protocol not supported.")
	}

	return
}

func (pBackend *backend_t) Unmarshal(in []byte) (
	server interfaces.IServer,
	err error,
) {
	var serverData map[string]interface{}
	err = json.Unmarshal(in, &serverData)
	if err != nil {
		return
	}

	server, err = pBackend.Map2Server(serverData)

	return
}

func (pBackend *backend_t) Map2Server(serverData map[string]interface{}) (
	server interfaces.IServer,
	err error,
) {
	item, ok := serverData["protocol"]
	if !ok {
		err = errors.New("server data lacks for protocol.")
		return
	}
	var protocolId string
	protocolId, ok = item.(string)
	if !ok {
		err = errors.New("protocol must be string.")
		return
	}

	var protocol interfaces.IProtocol
	protocol, err = pBackend.GetProtocol(protocolId)
	if err != nil {
		return
	}
	server, err = protocol.Map2Server(serverData)

	return
}

// New initializes Backend service.
func New(
	context interfaces.IApplicationContext,
	backends map[string]interfaces.IProtocol,
) interfaces.IBackend {
	backend := backend_t{
		context: context,

		backends: backends,
	}

	return &backend
}
