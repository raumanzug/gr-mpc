package mpd

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
	var mpdServer server_t
	{
		item, ok := m["dn"]
		var cerr error
		if !ok {
			cerr = errors.New("no display name given.")
		} else {
			mpdServer.dn, ok = item.(string)
			if !ok {
				cerr = errors.New("display name must be string.")
			}
		}
		err = errors.Join(err, cerr)
	}
	{
		item, ok := m["socket"]
		var cerr error
		if !ok {
			cerr = errors.New("no socket given.")
		} else {
			mpdServer.socket, ok = item.(string)
			if !ok {
				cerr = errors.New("socket must be string.")
			}
		}
		err = errors.Join(err, cerr)
	}
	{
		item, ok := m["password"]
		var cerr error
		if ok {
			var mpdPassword string
			mpdPassword, ok = item.(string)
			if !ok {
				cerr = errors.New("password must be string.")
			} else {
				mpdServer.password = &mpdPassword
			}
		}
		err = errors.Join(err, cerr)
	}

	server = mpdServer

	return
}

func New(context interfaces.IApplicationContext) interfaces.IProtocol {
	protocol := protocol_t{
		context: context,
	}

	return &protocol
}
