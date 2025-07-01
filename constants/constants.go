// Package constants contains global constants.
package constants

import (
	"image/color"
)

// AppName denotes the name of the application.
const AppName = "gr-mpc"

// keys in preferences database

// SelectedServerTabKey specifies the key in preferences database
// which refers to the currently used media rendering engine.
const SelectedServerTabKey = "selectedServerTab"

// TabbedServersKey specifies the key in preferences database which
// refers to a list of all media rendering engines known by this application.
const TabbedServersKey = "tabbedServers"

// StationSelectorBackgroundColor specifies background color for
// station selector.
var StationSelectorBackgroundColor = color.CMYK{0, 0, 0, 12}

// ErrorBackgroundColor specifies background color for error messages.
var ErrorBackgroundColor = color.CMYK{0, 12, 12, 0}

// StationsSelectorResizeFactor specifies resize factor for station selector.
const StationSelectorResizeFactor = 5.0
