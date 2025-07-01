package test

import (
	"errors"

	"github.com/raumanzug/gr-mpc/interfaces"
)

type protocol_t struct {
	context interfaces.IApplicationContext
}

func (pProtocol *protocol_t) CreateSTEF() (stef interfaces.ISTEF) {
	stef = NewSTEF()

	return
}

func (pProtocol *protocol_t) Map2Server(m map[string]interface{}) (
	server interfaces.IServer,
	err error,
) {
	var testServer server_t
	{
		item, ok := m["dn"]
		var cerr error
		if !ok {
			cerr = errors.New("no display name given.")
		} else {
			testServer.dn, ok = item.(string)
			if !ok {
				cerr = errors.New("display name must be string.")
			}
		}
		err = errors.Join(err, cerr)
	}

	server = testServer

	return
}

func New(context interfaces.IApplicationContext) interfaces.IProtocol {
	protocol := protocol_t{
		context: context,
	}

	return &protocol
}
