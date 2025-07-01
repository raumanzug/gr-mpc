package tetab

import (
	"github.com/raumanzug/gr-mpc/interfaces"
)

type BackingData struct {
	IsInitialPage bool
	PExc          *string
	STEF          interfaces.ISTEF
}
