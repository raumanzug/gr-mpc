// server tab service manages server tabs in the ui.
package servertabs

import (
	"errors"
	"maps"
	"slices"

	"github.com/raumanzug/gr-mpc/interfaces"
	"github.com/raumanzug/gr-mpc/utils/dbj"
)

type servertabs_t struct {
	context interfaces.IApplicationContext

	pMarkedTab    *string
	stateRegistry map[string]interfaces.IState
}

func (pServerTabs *servertabs_t) close(serverKey string) (err error) {
	state, nonvoid := pServerTabs.stateRegistry[serverKey]
	if !nonvoid {
		err = errors.New("Close before Open")
		return
	}
	dbj.ForceBackgroundJob(
		state,
		state.Close,
	)()
	delete(pServerTabs.stateRegistry, serverKey)

	return
}

func (pServerTabs *servertabs_t) markAsSelected(serverKey string) (err error) {
	if pServerTabs.pMarkedTab != nil &&
		*(pServerTabs.pMarkedTab) == serverKey {
		return
	}
	err = pServerTabs.markAsUnselected()
	if err != nil {
		return
	}
	if pServerTabs.context.Lifecycle().IsActive() {
		state, nonvoid := pServerTabs.stateRegistry[serverKey]
		if !nonvoid {
			err = errors.New("non-existent server specified")
			return
		}
		pServerTabs.pMarkedTab = &serverKey
		dbj.DoBackgroundJob(
			state,
			state.Start,
		)()
	}

	return
}

func (pServerTabs *servertabs_t) markAsUnselected() (err error) {
	if pServerTabs.pMarkedTab == nil {
		return
	}

	state, nonvoid := pServerTabs.stateRegistry[*pServerTabs.pMarkedTab]
	if !nonvoid {
		err = errors.New("non-existent server specified")
		return
	}

	dbj.DoBackgroundJob(
		state,
		state.Stop,
	)()

	return
}

func (pServerTabs *servertabs_t) Activate(server interfaces.IServer) (
	state interfaces.IState,
	err error,
) {
	state, err = server.Open()
	if err != nil {
		return
	}
	pServerTabs.stateRegistry[server.GetDisplayName()] = state

	return
}

func (pServerTabs *servertabs_t) GetActivatedState(serverKey string) (state interfaces.IState) {
	val, nonvoid := pServerTabs.stateRegistry[serverKey]
	if nonvoid {
		state = val
	}

	return
}

func (pServerTabs *servertabs_t) GetActivatedStates() []interfaces.IState {
	return slices.Collect(maps.Values(pServerTabs.stateRegistry))
}

func (pServerTabs *servertabs_t) RemoveServerTabCallback(key string) (err error) {
	err = pServerTabs.context.Preferences().RemoveServerTab(key)
	if err != nil {
		return
	}

	defer func() {
		err = errors.Join(err, pServerTabs.close(key))
	}()

	var pServerKey *string
	pServerKey = pServerTabs.context.Preferences().SelectedServerTabKey()
	if pServerKey != nil && *pServerKey != key {
		return
	}

	pNewTabItem := pServerTabs.context.UI().GetDocTabs().Selected()
	if pNewTabItem == nil {
		err = pServerTabs.markAsUnselected()
		if err != nil {
			return
		}
		err = pServerTabs.context.UI().UpdateScreen()
		return
	}

	err = pServerTabs.SelectCallback(pNewTabItem.Text)
	if err != nil {
		return
	}

	return
}

func (pServerTabs *servertabs_t) SelectCallback(key string) (err error) {
	err = pServerTabs.markAsSelected(key)
	if err != nil {
		return
	}

	err = pServerTabs.context.Preferences().SaveSelectedServerTab(key)
	if err != nil {
		return
	}

	return
}

func New(
	context interfaces.IApplicationContext,
) interfaces.IServerTabs {
	servertabs := servertabs_t{
		context: context,

		stateRegistry: make(map[string]interfaces.IState),
	}

	return &servertabs
}
