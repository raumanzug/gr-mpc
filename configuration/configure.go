package configuration

import (
	"embed"

	"github.com/raumanzug/gr-mpc/globals"
	"github.com/raumanzug/gr-mpc/interfaces"

	"fyne.io/fyne/v2/lang"
)

//go:embed translations
var translations embed.FS

// Configure configures services and set globals.ApplicationContext
// to the application context given as arg.
func Configure(context interfaces.IApplicationContext) {
	lang.AddTranslationsFS(translations, "translations")
	globals.ApplicationContext = context
	context.Lifecycle().Configure()
	context.UI().Configure()
}
