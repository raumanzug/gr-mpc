// Package globals contains global variables.
package globals

import (
	"github.com/raumanzug/gr-mpc/interfaces"
)

// ApplicationContext is the point where all services are accessible.
var ApplicationContext interfaces.IApplicationContext

// Provisorium
var MainProtocol interfaces.IProtocol
