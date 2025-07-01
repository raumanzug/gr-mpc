package test

import (
	"encoding/json"

	"github.com/raumanzug/gr-mpc/interfaces"
)

type server_t struct {
	dn string
}

func (server server_t) GetDisplayName() string {
	return server.dn
}

func (server server_t) GetProtocolId() string {
	return "test"
}

func (server server_t) MarshalJSON() (out []byte, err error) {

	in := map[string]interface{}{
		"protocol": "test",
		"dn":       server.dn,
	}

	out, err = json.Marshal(in)

	return
}

func (server server_t) Open() (state interfaces.IState, err error) {
	state = newState(server)

	return
}

func (server server_t) UpdateSTEF(stef interfaces.ISTEF) (err error) {
	stef.DisplayNameBinding().Set(server.dn)

	return
}
