package interfaces

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type ISTEF interface {
	DisplayNameBinding() binding.String
	UpdateForm(*widget.Form)
	Edited() (IServer, error)
}

type STEFBase_t struct {
	dn binding.String
}

func NewSTEFBase() *STEFBase_t {
	stefbase := STEFBase_t{
		dn: binding.NewString(),
	}

	return &stefbase
}

func (pSTEF *STEFBase_t) DisplayNameBinding() binding.String {
	return pSTEF.dn
}
