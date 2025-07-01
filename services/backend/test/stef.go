package test

import (
	"github.com/raumanzug/gr-mpc/interfaces"

	"fyne.io/fyne/v2/widget"
)

type stef_t struct {
	interfaces.STEFBase_t
}

func NewSTEF() interfaces.ISTEF {
	stef := stef_t{
		STEFBase_t: *(interfaces.NewSTEFBase()),
	}

	return &stef
}

func (pStef *stef_t) Edited() (server interfaces.IServer, err error) {
	var testServer server_t
	testServer.dn, err = pStef.DisplayNameBinding().Get()

	server = testServer

	return
}

func (pStef *stef_t) UpdateForm(pForm *widget.Form) {
}
