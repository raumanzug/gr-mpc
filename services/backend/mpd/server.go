package mpd

import (
	"encoding/json"
	"net"

	"github.com/raumanzug/gr-mpc/interfaces"
)

type server_t struct {
	dn       string
	socket   string
	password *string
}

func (server server_t) GetDisplayName() string {
	return server.dn
}

func (server server_t) GetProtocolId() string {
	return "mpd"
}

func (server server_t) MarshalJSON() (out []byte, err error) {

	in := map[string]interface{}{
		"protocol": "mpd",
		"dn":       server.dn,
		"socket":   server.socket,
	}
	if server.password != nil {
		in["password"] = server.password
	}

	out, err = json.Marshal(in)

	return
}

func (server server_t) Open() (state interfaces.IState, err error) {
	state = NewState(server)

	return
}

func (server server_t) UpdateSTEF(stef interfaces.ISTEF) (err error) {
	stef.DisplayNameBinding().Set(server.dn)
	switch s := stef.(type) {
	case *stef_t:
		var host, port string
		host, port, err = net.SplitHostPort(server.socket)
		s.host.Set(host)
		s.port.Set(port)

		s.isPassword.Set(server.password != nil)
		if server.password != nil {
			s.password.Set(*server.password)
		}
	}

	return
}
