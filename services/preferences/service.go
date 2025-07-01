// Package preferences implements Preferences service.
package preferences

import (
	"encoding/json"
	"errors"
	"slices"
	"strings"

	"github.com/raumanzug/gr-mpc/constants"
	"github.com/raumanzug/gr-mpc/interfaces"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/lang"
)

type preferences_t struct {
	context interfaces.IApplicationContext

	preferences fyne.Preferences
}

func bsf(serverTab interfaces.IServer, key string) int {
	return strings.Compare(serverTab.GetDisplayName(), key)
}

func (pPref *preferences_t) lookupServerTabIdx(
	serverTabs []interfaces.IServer,
	key string,
) (
	idx int,
	isFound bool,
	err error) {

	idx, isFound = slices.BinarySearchFunc(
		serverTabs,
		key,
		bsf,
	)

	return
}

func (pPref *preferences_t) saveServerTabs(serverTabs []interfaces.IServer) (
	err error) {

	var jsonServerTabs []string
	for _, serverTab := range serverTabs {
		jsonServerTab, cerr := json.Marshal(serverTab)
		if cerr != nil {
			err = errors.Join(err, cerr)
			continue
		}
		jsonServerTabs = append(jsonServerTabs, string(jsonServerTab))
	}

	pPref.preferences.SetStringList(
		constants.TabbedServersKey,
		jsonServerTabs,
	)

	return
}

func removeServerTabFromServerTabs(
	pServerTabs *[]interfaces.IServer,
	idx int,
) {
	copy((*pServerTabs)[idx:], (*pServerTabs)[idx+1:])
	*pServerTabs = (*pServerTabs)[:len(*pServerTabs)-1]
}

func (pPref *preferences_t) AddServerTab(serverTab interfaces.IServer) (err error) {
	var serverTabs []interfaces.IServer
	serverTabs, err = pPref.TabbedServers()
	if err != nil {
		return
	}
	serverTabs = append(serverTabs, serverTab)
	err = pPref.saveServerTabs(serverTabs)

	return
}

func (pPref *preferences_t) DoesNotExist(pExc *string) fyne.StringValidator {
	return func(key string) (err error) {
		var serverTabs []interfaces.IServer
		serverTabs, err = pPref.TabbedServers()
		if err != nil {
			return
		}
		if pExc != nil {
			var idx int
			var excFound bool
			idx, excFound, err = pPref.lookupServerTabIdx(
				serverTabs,
				*pExc,
			)
			if !excFound {
				err = errors.Join(err,
					errors.New("exc not found."),
				)
			}
			if err != nil {
				return
			}
			removeServerTabFromServerTabs(&serverTabs, idx)
		}

		_, isFound := slices.BinarySearchFunc(
			serverTabs,
			key,
			bsf,
		)

		if isFound {
			err = errors.New(
				lang.X(
					"dn already exists",
					"display name already exists",
				),
			)
		}

		return
	}
}

func (pPref *preferences_t) LookupServerTab(key string) (
	serverTab interfaces.IServer,
	err error) {
	var idx int
	var serverTabs []interfaces.IServer
	var isFound bool
	serverTabs, err = pPref.TabbedServers()
	if err != nil {
		return
	}
	idx, isFound, err = pPref.lookupServerTabIdx(serverTabs, key)
	if err != nil {
		return
	}

	if isFound {
		serverTab = serverTabs[idx]
	}

	return
}

func (pPref *preferences_t) RemoveServerTab(key string) (err error) {
	var serverTabs []interfaces.IServer
	var idx int
	var isFound bool
	serverTabs, err = pPref.TabbedServers()
	if err != nil {
		return
	}
	idx, isFound, err = pPref.lookupServerTabIdx(
		serverTabs,
		key,
	)
	if err != nil {
		return
	}
	if !isFound {
		err = errors.New("no server tab removed because server tab with given key could not be found.")
		return
	}
	removeServerTabFromServerTabs(&serverTabs, idx)
	err = pPref.saveServerTabs(serverTabs)

	return
}

func (pPref *preferences_t) SaveSelectedServerTab(key string) (err error) {
	pPref.preferences.SetString(constants.SelectedServerTabKey, key)

	return
}

func (pPref *preferences_t) SelectedServerTabKey() (pServerKey *string) {
	jsonTabbedServers := pPref.preferences.StringListWithFallback(
		constants.TabbedServersKey,
		[]string{},
	)

	if len(jsonTabbedServers) == 0 {
		return
	}

	selectedServerTab := pPref.preferences.String(
		constants.SelectedServerTabKey,
	)
	pServerKey = &selectedServerTab

	return
}

func (pPref *preferences_t) TabbedServers() (
	serverTabs []interfaces.IServer,
	err error,
) {
	jsonTabbedServers := pPref.preferences.StringListWithFallback(
		constants.TabbedServersKey,
		[]string{},
	)
	for _, jsonServerTab := range jsonTabbedServers {
		serverTab, cerr := pPref.context.Backend().
			Unmarshal([]byte(jsonServerTab))
		err = errors.Join(err, cerr)
		if cerr != nil {
			continue
		}
		serverTabs = append(serverTabs, serverTab)
	}

	slices.SortStableFunc(serverTabs, interfaces.ServerCompare)

	return
}

// New initializes Preferences service.
func New(
	context interfaces.IApplicationContext,
	prefs fyne.Preferences,
) interfaces.IPreferences {
	preferences := preferences_t{
		context: context,

		preferences: prefs,
	}

	return &preferences
}
