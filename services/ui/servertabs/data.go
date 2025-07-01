package servertabs

import (
	"github.com/raumanzug/gr-mpc/interfaces"

	"fyne.io/fyne/v2/container"
)

type BackingData struct {
	servers []interfaces.IServer
	docTabs *container.DocTabs
}

// AddDocTab adds an doc tab.
//
// Call it in fyne's thread only!
// Call it only once for each server!
// Call it after setting up doc tabs with SetupView only!
func (pBd *BackingData) AddDocTab(server interfaces.IServer) (err, terr error) {
	var tabItem *container.TabItem
	tabItem, err, terr = iServer2TabItem(server)
	pBd.docTabs.Append(tabItem)

	return
}

// GetDocTabs returns doc tabs container produced by SetupView.
//
// Call it after setting up doc tabs with SetupView only!
func (pBd *BackingData) GetDocTabs() *container.DocTabs {
	return pBd.docTabs
}

func NewBackingData(servers []interfaces.IServer) *BackingData {
	bd := BackingData{
		servers: servers,
	}

	return &bd
}
