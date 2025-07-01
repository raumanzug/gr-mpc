package test

import (
	"github.com/raumanzug/gr-mpc/configuration"
	"github.com/raumanzug/gr-mpc/globals"
	testService "github.com/raumanzug/gr-mpc/services/backend/test"
	testUtils "github.com/raumanzug/gr-mpc/services/test"
	testUtilsGlobals "github.com/raumanzug/gr-mpc/utils/test/globals"
)

func Test() {
	context := testUtils.New()
	globals.MainProtocol = testService.New(context)
	testUtilsGlobals.Gateway = newGateway()
	configuration.Configure(context)
	mainWindow := globals.ApplicationContext.UI().GetMainWindow()
	mainWindow.ShowAndRun()
}
