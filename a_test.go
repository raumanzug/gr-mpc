package main

import (
	"testing"

	"github.com/raumanzug/gr-mpc/globals"

	testUtils "github.com/raumanzug/gr-mpc/utils/test"
	testGlobals "github.com/raumanzug/gr-mpc/utils/test/globals"
	testUtilsInterfaces "github.com/raumanzug/gr-mpc/utils/test/interfaces"

	fyneTest "fyne.io/fyne/v2/test"
)

func TestA(t *testing.T) {
	testUtils.Test()
	gateway := testGlobals.Gateway
	gateway.Mutex().Lock()
	defer gateway.Mutex().Unlock()

	pForm := testUtils.ExtractForm(t)
	if pForm == nil {
		return
	}

	pEntry := testUtils.ExtractDNEntry(t, pForm)
	if pEntry == nil {
		return
	}

	if pForm.OnCancel != nil {
		t.Error("cancel button should be disabled but is not.")
	}

	if len(pEntry.Text) != 0 {
		t.Error("dn entry contains text but should not.")
	}

	if pForm.OnSubmit == nil {
		t.Error("submit button should be enabled but is not.")
		return
	}

	{
		fyneTest.Type(pEntry, "elephant")
		_, submitButton := testUtils.ExtractFormButtons(t, pForm)
		fyneTest.Tap(submitButton)
	}

	{
		servers, err := globals.ApplicationContext.Preferences().
			TabbedServers()
		if err != nil {
			t.Error(err.Error())
			return
		}
		if len(servers) != 1 {
			t.Error("one server expected.")
			return
		}
		if servers[0].GetDisplayName() != "elephant" {
			t.Errorf(
				"server elephant was expected in preferences, server %s given",
				servers[0].GetDisplayName(),
			)
		}
	}

	{
		pSelectedServerTabKey := globals.ApplicationContext.Preferences().
			SelectedServerTabKey()
		if pSelectedServerTabKey == nil ||
			*pSelectedServerTabKey != "elephant" {
			t.Error("a tab elephant should be selected")
		}
	}

	{
		docTabs := globals.ApplicationContext.UI().
			GetDocTabs()
		if len(docTabs.Items) != 1 {
			t.Error("one doc tab expected")
			return
		}
		itemText := docTabs.Items[0].Text
		if itemText != "elephant" {
			t.Errorf(
				"only doc tab should be elephant. %s given",
				itemText,
			)
		}
		pSelectedDocTabitem := docTabs.Selected()
		if pSelectedDocTabitem == nil ||
			pSelectedDocTabitem.Text != "elephant" {
			t.Error("doc tab elephant should be selected. is not.")
		}
	}
	testUtils.AssertTabbedServersServerTabsServiceCoincidence(t)

	state := globals.ApplicationContext.ServerTabs().GetActivatedState(
		"elephant",
	)
	if state == nil {
		t.Error("state elephant is nil.")
		return
	}

	testUtils.ExpectEventConnection[*testUtilsInterfaces.UpdateAllControlsEvent_t](
		"elephant",
		state,
	)
	testUtils.ExpectEventConnection[*testUtilsInterfaces.SetVolumeEvent_t](
		"elephant",
		state,
	)

	{
		settingsButton := testUtils.ExtractSettingsButton(t)
		if settingsButton == nil {
			t.Error("no settings button found.")
			return
		}
		fyneTest.Tap(settingsButton)
	}

	pForm = testUtils.ExtractForm(t)
	if pForm == nil {
		return
	}

	pEntry = testUtils.ExtractDNEntry(t, pForm)
	if pEntry == nil {
		return
	}

	if pEntry.Text != "elephant" {
		t.Error("dn entry should contain elephant.")
	}

	if pForm.OnSubmit == nil {
		t.Error("submit button should be enabled but is not.")
	}

	if pForm.OnCancel == nil {
		t.Error("cancel button should be enabled but is not.")
		return
	}

	{
		cancelButton, _ := testUtils.ExtractFormButtons(t, pForm)
		if cancelButton == nil {
			t.Error("no cancel button found.")
			return
		}
		fyneTest.Tap(cancelButton)
	}

	{
		servers, err := globals.ApplicationContext.Preferences().
			TabbedServers()
		if err != nil {
			t.Error(err.Error())
			return
		}
		if len(servers) != 1 {
			t.Error("one server expected.")
			return
		}
		if servers[0].GetDisplayName() != "elephant" {
			t.Errorf(
				"server elephant was expected in preferences, server %s given",
				servers[0].GetDisplayName(),
			)
		}
	}

	{
		pSelectedServerTabKey := globals.ApplicationContext.Preferences().
			SelectedServerTabKey()
		if pSelectedServerTabKey == nil ||
			*pSelectedServerTabKey != "elephant" {
			t.Error("a tab elephant should be selected")
		}
	}

	{
		docTabs := globals.ApplicationContext.UI().
			GetDocTabs()
		if len(docTabs.Items) != 1 {
			t.Error("one doc tab expected")
			return
		}
		itemText := docTabs.Items[0].Text
		if itemText != "elephant" {
			t.Errorf(
				"only doc tab should be elephant. %s given",
				itemText,
			)
		}
		pSelectedDocTabitem := docTabs.Selected()
		if pSelectedDocTabitem == nil ||
			pSelectedDocTabitem.Text != "elephant" {
			t.Error("doc tab elephant should be selected. is not.")
		}
	}
	testUtils.AssertTabbedServersServerTabsServiceCoincidence(t)

	{
		settingsButton := testUtils.ExtractSettingsButton(t)
		if settingsButton == nil {
			t.Error("no settings button found.")
			return
		}
		fyneTest.Tap(settingsButton)
	}

	pForm = testUtils.ExtractForm(t)
	if pForm == nil {
		return
	}

	pEntry = testUtils.ExtractDNEntry(t, pForm)
	if pEntry == nil {
		return
	}

	if pEntry.Text != "elephant" {
		t.Error("dn entry should contain elephant.")
	}

	if pForm.OnCancel == nil {
		t.Error("cancel button should be enabled but is not.")
	}

	if pForm.OnSubmit == nil {
		t.Error("submit button should be enabled but is not.")
		return
	}

	{
		fyneTest.Type(pEntry, "big ")
		_, submitButton := testUtils.ExtractFormButtons(t, pForm)
		fyneTest.Tap(submitButton)
	}

	testUtils.ExpectCloseEvent(
		"elephant",
		state,
	)

	state = globals.ApplicationContext.ServerTabs().GetActivatedState(
		"big elephant",
	)
	if state == nil {
		t.Error("state big elephant is nil.")
		return
	}

	testUtils.ExpectEventConnection[*testUtilsInterfaces.UpdateAllControlsEvent_t](
		"big elephant",
		state,
	)
	testUtils.ExpectEventConnection[*testUtilsInterfaces.SetVolumeEvent_t](
		"big elephant",
		state,
	)

	testUtils.AssertSelectedServerTabConsistency(t)
	testUtils.AssertTabbedServicesConsistency(t)
	testUtils.AssertTabbedServersContainsSelectedServerTab(t)
	testUtils.AssertTabbedServersServerTabsServiceCoincidence(t)

	{
		servers, err := globals.ApplicationContext.Preferences().
			TabbedServers()
		if err != nil {
			t.Error(err.Error())
			return
		}
		if len(servers) != 1 {
			t.Error("one server expected.")
			return
		}
	}

	// fyne does not call event handler when Remove/RemoveIndex method
	// is called on DocTabs.
	globals.ApplicationContext.UI().RemoveServerTab("big elephant")

	pForm = testUtils.ExtractForm(t)
	if pForm == nil {
		return
	}

	pEntry = testUtils.ExtractDNEntry(t, pForm)
	if pEntry == nil {
		return
	}

	if pForm.OnCancel != nil {
		t.Error("cancel button should be disabled but is not.")
	}

	if len(pEntry.Text) != 0 {
		t.Error("dn entry contains text but should not.")
	}

	if pForm.OnSubmit == nil {
		t.Error("submit button should be enabled but is not.")
		return
	}

	{
		fyneTest.Type(pEntry, "elephant")
		_, submitButton := testUtils.ExtractFormButtons(t, pForm)
		fyneTest.Tap(submitButton)
	}

	{
		servers, err := globals.ApplicationContext.Preferences().
			TabbedServers()
		if err != nil {
			t.Error(err.Error())
			return
		}
		if len(servers) != 1 {
			t.Error("one server expected.")
			return
		}
		if servers[0].GetDisplayName() != "elephant" {
			t.Errorf(
				"server elephant was expected in preferences, server %s given",
				servers[0].GetDisplayName(),
			)
		}
	}

	{
		pSelectedServerTabKey := globals.ApplicationContext.Preferences().
			SelectedServerTabKey()
		if pSelectedServerTabKey == nil ||
			*pSelectedServerTabKey != "elephant" {
			t.Error("a tab elephant should be selected")
		}
	}

	{
		docTabs := globals.ApplicationContext.UI().
			GetDocTabs()
		if len(docTabs.Items) != 1 {
			t.Error("one doc tab expected")
			return
		}
		itemText := docTabs.Items[0].Text
		if itemText != "elephant" {
			t.Errorf(
				"only doc tab should be elephant. %s given",
				itemText,
			)
		}
		pSelectedDocTabitem := docTabs.Selected()
		if pSelectedDocTabitem == nil ||
			pSelectedDocTabitem.Text != "elephant" {
			t.Error("doc tab elephant should be selected. is not.")
		}
	}
	testUtils.AssertTabbedServersServerTabsServiceCoincidence(t)

	testUtils.ExpectCloseEvent(
		"big elephant",
		state,
	)

	state = globals.ApplicationContext.ServerTabs().GetActivatedState(
		"elephant",
	)
	if state == nil {
		t.Error("state elephant is nil.")
		return
	}

	testUtils.ExpectEventConnection[*testUtilsInterfaces.UpdateAllControlsEvent_t](
		"elephant",
		state,
	)
	testUtils.ExpectEventConnection[*testUtilsInterfaces.SetVolumeEvent_t](
		"elephant",
		state,
	)

	{
		contentAddButton := testUtils.ExtractContentAddButton(t)
		if contentAddButton == nil {
			t.Error("no content add button found.")
			return
		}
		fyneTest.Tap(contentAddButton)
	}

	pForm = testUtils.ExtractForm(t)
	if pForm == nil {
		return
	}

	pEntry = testUtils.ExtractDNEntry(t, pForm)
	if pEntry == nil {
		return
	}

	if pForm.OnCancel == nil {
		t.Error("cancel button should be enabled but is not.")
	}

	if len(pEntry.Text) != 0 {
		t.Error("dn entry contains text but should not.")
	}

	if pForm.OnSubmit == nil {
		t.Error("submit button should be enabled but is not.")
		return
	}

	{
		fyneTest.Type(pEntry, "hippopotamus")
		_, submitButton := testUtils.ExtractFormButtons(t, pForm)
		fyneTest.Tap(submitButton)
	}
	testUtils.AssertTabbedServersServerTabsServiceCoincidence(t)

	state = globals.ApplicationContext.ServerTabs().GetActivatedState(
		"hippopotamus",
	)
	if state == nil {
		t.Error("state hippopotamus is nil.")
		return
	}

	testUtils.ExpectEventConnection[*testUtilsInterfaces.UpdateAllControlsEvent_t](
		"hippopotamus",
		state,
	)
	testUtils.ExpectEventConnection[*testUtilsInterfaces.SetVolumeEvent_t](
		"hippopotamus",
		state,
	)

	{
		servers, err := globals.ApplicationContext.Preferences().
			TabbedServers()
		if err != nil {
			t.Error(err.Error())
			return
		}
		if len(servers) != 2 {
			t.Error("two servers expected.")
			return
		}
	}

	{
		pSelectedServerTabKey := globals.ApplicationContext.Preferences().
			SelectedServerTabKey()
		if pSelectedServerTabKey == nil ||
			*pSelectedServerTabKey != "hippopotamus" {
			t.Error("a tab hippopotamus should be selected")
		}
	}

	testUtils.AssertSelectedServerTabConsistency(t)
	testUtils.AssertTabbedServicesConsistency(t)
	testUtils.AssertTabbedServersContainsSelectedServerTab(t)
	testUtils.AssertTabbedServersServerTabsServiceCoincidence(t)

	{
		docTabs := globals.ApplicationContext.UI().GetDocTabs()
		oldIndex := docTabs.SelectedIndex()
		newIndex := (oldIndex + 1) % 2
		docTabs.SelectIndex(newIndex)
	}

	{
		servers, err := globals.ApplicationContext.Preferences().
			TabbedServers()
		if err != nil {
			t.Error(err.Error())
			return
		}
		if len(servers) != 2 {
			t.Error("two servers expected.")
			return
		}
	}

	{
		pSelectedServerTabKey := globals.ApplicationContext.Preferences().
			SelectedServerTabKey()
		if pSelectedServerTabKey == nil ||
			*pSelectedServerTabKey != "elephant" {
			t.Error("a tab elephant should be selected")
		}
	}

	testUtils.AssertSelectedServerTabConsistency(t)
	testUtils.AssertTabbedServicesConsistency(t)
	testUtils.AssertTabbedServersContainsSelectedServerTab(t)
	testUtils.AssertTabbedServersServerTabsServiceCoincidence(t)

	globals.ApplicationContext.UI().RemoveServerTab("hippopotamus")

	{
		servers, err := globals.ApplicationContext.Preferences().
			TabbedServers()
		if err != nil {
			t.Error(err.Error())
			return
		}
		if len(servers) != 1 {
			t.Error("one server expected.")
			return
		}
		if servers[0].GetDisplayName() != "elephant" {
			t.Errorf(
				"server elephant was expected in preferences, server %s given",
				servers[0].GetDisplayName(),
			)
		}
	}

	{
		docTabs := globals.ApplicationContext.UI().
			GetDocTabs()
		if len(docTabs.Items) != 1 {
			t.Error("one doc tab expected")
			return
		}
		itemText := docTabs.Items[0].Text
		if itemText != "elephant" {
			t.Errorf(
				"only doc tab should be elephant. %s given",
				itemText,
			)
		}
	}
	testUtils.AssertTabbedServersServerTabsServiceCoincidence(t)

	{
		contentAddButton := testUtils.ExtractContentAddButton(t)
		if contentAddButton == nil {
			t.Error("no content add button found.")
			return
		}
		fyneTest.Tap(contentAddButton)
	}

	pForm = testUtils.ExtractForm(t)
	if pForm == nil {
		return
	}

	pEntry = testUtils.ExtractDNEntry(t, pForm)
	if pEntry == nil {
		return
	}

	if pForm.OnCancel == nil {
		t.Error("cancel button should be enabled but is not.")
	}

	if len(pEntry.Text) != 0 {
		t.Error("dn entry contains text but should not.")
	}

	if pForm.OnSubmit == nil {
		t.Error("submit button should be enabled but is not.")
		return
	}

	{
		fyneTest.Type(pEntry, "hippopotamus")
		_, submitButton := testUtils.ExtractFormButtons(t, pForm)
		fyneTest.Tap(submitButton)
	}
	testUtils.AssertTabbedServersServerTabsServiceCoincidence(t)

	testUtils.ExpectCloseEvent(
		"hippopotamus",
		state,
	)

	state = globals.ApplicationContext.ServerTabs().GetActivatedState(
		"hippopotamus",
	)
	if state == nil {
		t.Error("state hippopotamus is nil.")
		return
	}

	testUtils.ExpectEventConnection[*testUtilsInterfaces.UpdateAllControlsEvent_t](
		"hippopotamus",
		state,
	)
	testUtils.ExpectEventConnection[*testUtilsInterfaces.SetVolumeEvent_t](
		"hippopotamus",
		state,
	)

	{
		servers, err := globals.ApplicationContext.Preferences().
			TabbedServers()
		if err != nil {
			t.Error(err.Error())
			return
		}
		if len(servers) != 2 {
			t.Error("two servers expected.")
			return
		}
	}

	{
		pSelectedServerTabKey := globals.ApplicationContext.Preferences().
			SelectedServerTabKey()
		if pSelectedServerTabKey == nil ||
			*pSelectedServerTabKey != "hippopotamus" {
			t.Error("a tab hippopotamus should be selected")
		}
	}

	testUtils.AssertSelectedServerTabConsistency(t)
	testUtils.AssertTabbedServicesConsistency(t)
	testUtils.AssertTabbedServersContainsSelectedServerTab(t)
	testUtils.AssertTabbedServersServerTabsServiceCoincidence(t)

	globals.ApplicationContext.UI().RemoveServerTab("hippopotamus")

	{
		servers, err := globals.ApplicationContext.Preferences().
			TabbedServers()
		if err != nil {
			t.Error(err.Error())
			return
		}
		if len(servers) != 1 {
			t.Error("one server expected.")
			return
		}
		if servers[0].GetDisplayName() != "elephant" {
			t.Errorf(
				"server elephant was expected in preferences, server %s given",
				servers[0].GetDisplayName(),
			)
		}
	}

	{
		docTabs := globals.ApplicationContext.UI().
			GetDocTabs()
		if len(docTabs.Items) != 1 {
			t.Error("one doc tab expected")
			return
		}
		itemText := docTabs.Items[0].Text
		if itemText != "elephant" {
			t.Errorf(
				"only doc tab should be elephant. %s given",
				itemText,
			)
		}
	}
	testUtils.AssertTabbedServersServerTabsServiceCoincidence(t)

	// testing controls buttons
	{
		playButton := testUtils.ExtractPlayButton(t)
		if playButton == nil {
			t.Error("no play button found.")
			return
		}
		fyneTest.Tap(playButton)
	}
	testUtils.AssertTabbedServersServerTabsServiceCoincidence(t)

	testUtils.ExpectCloseEvent(
		"hippopotamus",
		state,
	)

	state = globals.ApplicationContext.ServerTabs().GetActivatedState(
		"elephant",
	)
	if state == nil {
		t.Error("state elephant is nil.")
		return
	}

	testUtils.ExpectEventConnection[*testUtilsInterfaces.PlayEvent_t](
		"elephant",
		state,
	)

	{
		pauseButton := testUtils.ExtractPauseButton(t)
		if pauseButton == nil {
			t.Error("no pause button found.")
			return
		}
		fyneTest.Tap(pauseButton)
	}
	testUtils.AssertTabbedServersServerTabsServiceCoincidence(t)

	testUtils.ExpectEventConnection[*testUtilsInterfaces.PauseEvent_t](
		"elephant",
		state,
	)

	{
		stopButton := testUtils.ExtractStopButton(t)
		if stopButton == nil {
			t.Error("no stop button found.")
			return
		}
		fyneTest.Tap(stopButton)
	}
	testUtils.AssertTabbedServersServerTabsServiceCoincidence(t)

	testUtils.ExpectEventConnection[*testUtilsInterfaces.StopEvent_t](
		"elephant",
		state,
	)

	{
		skipNextButton := testUtils.ExtractSkipNextButton(t)
		if skipNextButton == nil {
			t.Error("no skip next button found.")
			return
		}
		fyneTest.Tap(skipNextButton)
	}
	testUtils.AssertTabbedServersServerTabsServiceCoincidence(t)

	testUtils.ExpectEventConnection[*testUtilsInterfaces.SkipNextEvent_t](
		"elephant",
		state,
	)

	{
		skipPreviousButton := testUtils.ExtractSkipPreviousButton(t)
		if skipPreviousButton == nil {
			t.Error("no skip previous button found.")
			return
		}
		fyneTest.Tap(skipPreviousButton)
	}
	testUtils.AssertTabbedServersServerTabsServiceCoincidence(t)

	testUtils.ExpectEventConnection[*testUtilsInterfaces.SkipPreviousEvent_t](
		"elephant",
		state,
	)

	// testing whether events are dropped if mpd is blocked.
	{
		playButton := testUtils.ExtractPlayButton(t)
		if playButton == nil {
			t.Error("no play button found.")
			return
		}
		fyneTest.Tap(playButton)
	}
	testUtils.AssertTabbedServersServerTabsServiceCoincidence(t)

	{
		pauseButton := testUtils.ExtractPauseButton(t)
		if pauseButton == nil {
			t.Error("no pause button found.")
			return
		}
		fyneTest.Tap(pauseButton)
	}
	testUtils.AssertTabbedServersServerTabsServiceCoincidence(t)

	testUtils.ExpectEventConnection[*testUtilsInterfaces.PlayEvent_t](
		"elephant",
		state,
	)

}
