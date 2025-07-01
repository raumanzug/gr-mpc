// Mpc is a graphical app for controlling MPD servers
// that is intended for use on smartphones.
package main

import (
	"github.com/raumanzug/gr-mpc/configuration"
	"github.com/raumanzug/gr-mpc/globals"
	"github.com/raumanzug/gr-mpc/services/application"
	"github.com/raumanzug/gr-mpc/services/backend/mpd" // Provisorium
)

func main() {
	context := application.New()
	globals.MainProtocol = mpd.New(context) // Provisorium
	configuration.Configure(context)
	mainWindow := globals.ApplicationContext.UI().GetMainWindow()
	mainWindow.ShowAndRun()
}
