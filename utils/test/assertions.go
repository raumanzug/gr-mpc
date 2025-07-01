package test

import (
	"slices"
	"testing"

	"github.com/raumanzug/gr-mpc/globals"
	"github.com/raumanzug/gr-mpc/interfaces"
)

func AssertSelectedServerTabConsistency(t *testing.T) {
	pSelectedServerTabKey := globals.ApplicationContext.Preferences().
		SelectedServerTabKey()
	if pSelectedServerTabKey == nil {
		return
	}

	pTabItem := globals.ApplicationContext.UI().
		GetDocTabs().Selected()

	if pTabItem == nil {
		t.Error("nothing selected")
		return
	}

	if *pSelectedServerTabKey != pTabItem.Text {
		t.Error("selected server tab in preferences and in doc tabs do not coindice")
	}
}

func AssertTabbedServicesConsistency(t *testing.T) {
	servers, err := globals.ApplicationContext.Preferences().
		TabbedServers()
	if err != nil {
		t.Error(err.Error())
		return
	}
	if len(servers) == 0 {
		return
	}
	prefServerTabs := slices.Collect(
		func(yield func(string) bool) {
			for _, server := range servers {
				if !yield(server.GetDisplayName()) {
					return
				}
			}
		},
	)
	slices.Sort(prefServerTabs)

	docTabs := globals.ApplicationContext.UI().
		GetDocTabs()
	docServerTabs := slices.Collect(
		func(yield func(string) bool) {
			for _, item := range docTabs.Items {
				if !yield(item.Text) {
					return
				}
			}
		},
	)
	slices.Sort(docServerTabs)

	if !slices.Equal(prefServerTabs, docServerTabs) {
		t.Error("serverTabs in preferences and in doc tabs do not coindice")
	}
}

func AssertTabbedServersContainsSelectedServerTab(t *testing.T) {
	pSelectedServerTabKey := globals.ApplicationContext.Preferences().
		SelectedServerTabKey()
	if pSelectedServerTabKey == nil {
		return
	}

	servers, err := globals.ApplicationContext.Preferences().
		TabbedServers()
	if err != nil {
		t.Error(err.Error())
		return
	}

	if !slices.ContainsFunc(
		servers,
		func(server interfaces.IServer) bool {
			return *pSelectedServerTabKey == server.GetDisplayName()
		},
	) {
		t.Error("tabbed servers does not contain selected tab key.")
	}
}

func AssertTabbedServersServerTabsServiceCoincidence(t *testing.T) {
	prefServers, err := globals.ApplicationContext.Preferences().
		TabbedServers()
	if err != nil {
		t.Error(err.Error())
		return
	}

	activatedStates := globals.ApplicationContext.ServerTabs().
		GetActivatedStates()

	if len(prefServers) != len(activatedStates) {
		t.Error("len of tabbed servers gallery does not coindice to len of server tabs activated states gallery.")
	}

	for _, server := range prefServers {
		if globals.ApplicationContext.ServerTabs().
			GetActivatedState(server.GetDisplayName()) == nil {
			t.Error("there are servers in tabbed servers gallery that are not related to any state in server tabs activated states gallery.")
		}
	}
}
