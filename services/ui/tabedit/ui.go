// Provisorium
package tabedit

import (
	"github.com/raumanzug/gr-mpc/globals"
	"github.com/raumanzug/gr-mpc/interfaces"
	"github.com/raumanzug/gr-mpc/services/ui/tetab"

	"fyne.io/fyne/v2"
)

func (pBD *BackingData) SetupView() (fyne.CanvasObject, error, error) {
	return pBD.SetupErrlesslyView(), nil, nil
}

func tabedit2tetab(pBD *BackingData) *tetab.BackingData {
	protocol := globals.MainProtocol // Provisorium
	stef := protocol.CreateSTEF()
	bd := tetab.BackingData{
		PExc: pBD.PExc,
		STEF: stef,
	}
	pBD.STEF = stef

	return &bd
}

func (pBD *BackingData) SetupErrlesslyView() (content fyne.CanvasObject) {
	return tabedit2tetab(pBD).SetupErrlesslyView()
}

func (pBD *BackingData) UpdateSTEF(server interfaces.IServer) (err error) {
	err = server.UpdateSTEF(pBD.STEF)

	return
}
