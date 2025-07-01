// Package ui implements UI service.
package ui

import (
	"errors"

	"github.com/raumanzug/gr-mpc/constants"
	"github.com/raumanzug/gr-mpc/interfaces"
	"github.com/raumanzug/gr-mpc/services/ui/report"
	"github.com/raumanzug/gr-mpc/services/ui/servertabs"
	"github.com/raumanzug/gr-mpc/services/ui/tabedit"
	"github.com/raumanzug/gr-mpc/services/ui/workpad"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

const stackHeight = 3

type ui_t struct {
	context interfaces.IApplicationContext

	app                fyne.App
	mainWindow         fyne.Window
	pDocTabsBD         *servertabs.BackingData
	pRoot              *fyne.Container
	pScreen            *fyne.Container
	pReportBD          *report.BackingData
	reportBox          fyne.CanvasObject
	viewStack          []fyne.CanvasObject
	viewStackContainer [stackHeight]fyne.CanvasObject
	workpad            fyne.CanvasObject
}

// adjusting the selected server tab settings in preferences to those
func (pUi *ui_t) adjustServerTabs(serverTabs []interfaces.IServer) (err error, terr error) {
	pSst := pUi.context.Preferences().SelectedServerTabKey()
	if pSst == nil {
		pickSth := serverTabs[0].GetDisplayName()
		pSst = &pickSth
		err = pUi.context.Preferences().SaveSelectedServerTab(pickSth)
		if err != nil {
			return
		}
	}

	pDocTabs, idx, isFound := pUi.searchDocTab(*pSst)
	if isFound {
		pDocTabs.SelectIndex(idx)
	} else {
		pTabItem := pDocTabs.Selected()
		if pTabItem == nil && len(pDocTabs.Items) != 0 {
			pTabItem = pDocTabs.Items[0]
			pDocTabs.Select(pTabItem)
		}
		if pTabItem != nil {
			*pSst = pTabItem.Text
			err = pUi.context.Preferences().
				SaveSelectedServerTab(pTabItem.Text)
		}
	}

	return
}

func (pUi *ui_t) clearViewStack() {
	pUi.viewStack = pUi.viewStackContainer[:0]
}

// in the doc tabs.
// read the server tabs from preferences database and initializes
// doc tabs.
func (pUi *ui_t) initializeServerTabs() (err error, terr error) {
	var serverTabs []interfaces.IServer
	serverTabs, err = pUi.context.Preferences().TabbedServers()
	if err != nil {
		return
	}

	pUi.pDocTabsBD = servertabs.NewBackingData(serverTabs)

	_, err, terr = pUi.pDocTabsBD.SetupView()
	if err != nil {
		return
	}

	if len(serverTabs) == 0 {
		return
	}

	var tcerr error
	err, tcerr = pUi.adjustServerTabs(serverTabs)
	terr = errors.Join(terr, tcerr)

	return
}

func (pUi *ui_t) replaceScreen(view fyne.CanvasObject) {
	pUi.pScreen.RemoveAll()
	pUi.pScreen.Add(view)
}

func (pUi *ui_t) searchDocTab(key string) (
	pDocTabs *container.DocTabs,
	idx int,
	isFound bool,
) {
	pDocTabs = pUi.pDocTabsBD.GetDocTabs()
	isFound = false

	for i, pTabItem := range pDocTabs.Items {
		if pTabItem.Text == key {
			idx = i
			isFound = true
			return
		}
	}

	return
}

func (pUi *ui_t) showErr() {
	pScreen := pUi.pScreen
	if !pUi.pReportBD.IsEmpty() {
		pScreen = container.NewVBox(
			pUi.reportBox,
			pScreen,
		)
	}
	pUi.pRoot.RemoveAll()
	pUi.pRoot.Add(pScreen)
}

func (pUi *ui_t) AddErr(err error) {
	pUi.pReportBD.AddErr(err)
	pUi.showErr()
}

func (pUi *ui_t) AddServerTab(server interfaces.IServer) (err error, terr error) {
	err, terr = pUi.pDocTabsBD.AddDocTab(server)
	if err != nil {
		return
	}

	err = pUi.context.Preferences().AddServerTab(server)
	if err != nil {
		return
	}

	err = pUi.context.ServerTabs().
		SelectCallback(server.GetDisplayName())

	return
}

func (pUi *ui_t) BottomReached() bool {
	return len(pUi.viewStack) < 1
}

func (pUi *ui_t) CloseTopmostView() (err error) {
	idx := len(pUi.viewStack) - 1
	if idx < 1 {
		err = errors.New("no views on stack that could be closed.")
		return
	}
	pUi.replaceScreen(pUi.viewStackContainer[idx-1])
	pUi.viewStack = pUi.viewStackContainer[:idx]

	return
}

func (pUi *ui_t) Configure() {

	err, terr := pUi.initializeServerTabs()
	if err != nil {
		return
	}

	var tcerr error
	wpbd := workpad.BackingData(struct{}{})
	pUi.workpad, err, tcerr = wpbd.SetupView() // called after initializeServerTabs
	// because wpdb.SetupView demands for doc tabs produced by it.
	if err != nil {
		return
	}
	terr = errors.Join(terr, tcerr)

	pUi.pReportBD.ClearCallback = func() {
		pUi.showErr()
	}
	pUi.reportBox = pUi.pReportBD.SetupErrlesslyView()
	err = errors.Join(
		err,
		pUi.UpdateScreen(),
	)
	pUi.showErr()

	pUi.mainWindow.SetContent(pUi.pRoot)

	pUi.AddErr(err)
	pUi.AddErr(terr)
}

func (pUi *ui_t) GetApp() fyne.App {
	return pUi.app
}

func (pUi *ui_t) GetDocTabs() *container.DocTabs {
	return pUi.pDocTabsBD.GetDocTabs()
}

func (pUi *ui_t) GetMainWindow() fyne.Window {
	return pUi.mainWindow
}

func (pUi *ui_t) PushView(view fyne.CanvasObject) (err error) {
	idx := len(pUi.viewStack)
	if idx == stackHeight {
		err = errors.New("top reached; no view can be pushed on top of view stack.")
		return
	}

	pUi.viewStackContainer[idx] = view
	pUi.viewStack = pUi.viewStackContainer[:idx+1]

	pUi.replaceScreen(view)

	return
}

func (pUi *ui_t) RemoveServerTab(key string) (err error) {
	pDocTabs, idx, isFound := pUi.searchDocTab(key)
	if !isFound {
		err = errors.New("doc tab could not be removed because it was not found.")
		return
	}
	pDocTabs.RemoveIndex(idx)
	err = pUi.context.ServerTabs().RemoveServerTabCallback(key)

	return
}

func (pUi *ui_t) UpdateScreen() (err error) {
	var serverTabs []interfaces.IServer
	serverTabs, err = pUi.context.Preferences().TabbedServers()
	if err != nil {
		return
	}
	pUi.clearViewStack()
	if len(serverTabs) == 0 {
		bd := tabedit.BackingData{
			PExc: nil,
		}
		view := bd.SetupErrlesslyView()
		err = pUi.PushView(view)

		return
	}
	var terr error

	err, terr = pUi.adjustServerTabs(serverTabs)
	if err != nil {
		return
	}
	pUi.AddErr(terr)
	err = pUi.PushView(pUi.workpad)

	return
}

func (pUi *ui_t) UpdateRoot(err error) {
	pUi.pReportBD.Clear()
	pUi.pReportBD.AddErr(err)
	pUi.showErr()
}

// New initializes UI service.
func New(
	context interfaces.IApplicationContext,
	app fyne.App,
) interfaces.IUI {
	pRbd := report.NewBackingData()
	ui := ui_t{
		context: context,

		app:        app,
		mainWindow: app.NewWindow(constants.AppName),
		pReportBD:  pRbd,
		pRoot:      container.NewStack(),
		pScreen:    container.NewStack(),
	}
	ui.viewStack = ui.viewStackContainer[:0]

	return &ui
}
