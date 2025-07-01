package mpd

import (
	"errors"
	"net"

	"github.com/raumanzug/gr-mpc/globals"
	"github.com/raumanzug/gr-mpc/interfaces"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/lang"
	"fyne.io/fyne/v2/widget"
)

type stef_t struct {
	interfaces.STEFBase_t

	host       binding.String
	port       binding.String
	isPassword binding.Bool
	password   binding.String
}

func NewSTEF() interfaces.ISTEF {
	stef := stef_t{
		STEFBase_t: *(interfaces.NewSTEFBase()),
		host:       binding.NewString(),
		port:       binding.NewString(),
		isPassword: binding.NewBool(),
		password:   binding.NewString(),
	}

	stef.prefill()

	return &stef
}

func (pStef *stef_t) prefill() {
	pStef.port.Set("6600")
}

func (pStef *stef_t) UpdateForm(pForm *widget.Form) {
	passwordContainer := container.NewStack()
	{
		listener := func() {
			isPassword, err := pStef.isPassword.Get()
			passwordContainer.RemoveAll()
			var view fyne.CanvasObject
			if isPassword {
				view = (*passwordBackingData_t)(pStef).SetupErrlesslyView()
			} else {
				view = (*nopwBackingData_t)(pStef).SetupErrlesslyView()
			}
			passwordContainer.Add(view)
			globals.ApplicationContext.UI().AddErr(err)
		}
		pStef.isPassword.AddListener(binding.NewDataListener(listener))
	}
	pForm.Append(
		lang.X("host", "host"),
		widget.NewEntryWithData(pStef.host),
	)
	pForm.Append(
		lang.X("port", "port"),
		widget.NewEntryWithData(pStef.port),
	)
	pForm.Append(
		lang.X("password?", "password?"),
		passwordContainer,
	)
}

func (pStef *stef_t) Edited() (server interfaces.IServer, err error) {
	var mpdServer server_t
	mpdServer.dn, err = pStef.DisplayNameBinding().Get()
	var cerr error
	var host, port string
	host, cerr = pStef.host.Get()
	err = errors.Join(err, cerr)
	port, cerr = pStef.port.Get()
	err = errors.Join(err, cerr)
	mpdServer.socket = net.JoinHostPort(host, port)
	var isPassword bool
	isPassword, cerr = pStef.isPassword.Get()
	err = errors.Join(err, cerr)
	if isPassword {
		var password string
		password, cerr = pStef.password.Get()
		err = errors.Join(err, cerr)
		mpdServer.password = &password
	}

	server = mpdServer

	return
}
