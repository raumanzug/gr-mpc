package test

import (
	"iter"
	"slices"
	"testing"

	"github.com/raumanzug/gr-mpc/globals"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/lang"
	fyneTest "fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/raumanzug/gr-set"
)

func ExtractForm(t *testing.T) (pForm *widget.Form) {
	content := globals.ApplicationContext.UI().
		GetMainWindow().Content()
	laidOutObjects := fyneTest.LaidOutObjects(content)
	for _, obj := range laidOutObjects {
		switch v := obj.(type) {
		case *widget.Form:
			if pForm != nil {
				t.Error("there are more than one form on screen.")
				return
			}
			pForm = v
		}
	}
	if pForm == nil {
		t.Error("there is no form on screen.")
	}

	return
}

func ExtractDNEntry(t *testing.T, pForm *widget.Form) (pEntry *widget.Entry) {
	for _, item := range pForm.Items {
		if item.Text != lang.X("display name", "display name") {
			continue
		}
		switch v := item.Widget.(type) {
		case *widget.Entry:
			if pEntry != nil {
				t.Error("there is already a dn form entry in this form.")
				return
			}
			pEntry = v
		default:
			continue
		}
	}
	if pEntry == nil {
		t.Error("there is no dn form entry in this form.")
	}

	return
}

func ExtractFormButtons(t *testing.T, pForm *widget.Form) (
	cancelButton *widget.Button,
	submitButton *widget.Button,
) {
	laidOutObjectsSlice := fyneTest.LaidOutObjects(pForm)
	var laidOutObjects set.Set[fyne.CanvasObject] = set.NewSimpleSet[fyne.CanvasObject]()
	for _, laidOutObject := range laidOutObjectsSlice {
		laidOutObjects.Add(laidOutObject)
	}

	for _, formItem := range pForm.Items {
		formObjects := fyneTest.LaidOutObjects(formItem.Widget)
		for _, formObject := range formObjects {
			laidOutObjects.Remove(formObject)
		}
	}

	for laidOutObject := range laidOutObjects.Generator() {
		switch obj := laidOutObject.(type) {
		case *widget.Button:
			switch obj.Text {
			case lang.X("cancel form", "cancel"):
				if cancelButton != nil {
					t.Error("more than one cancel button.")
				}
				cancelButton = obj
			case lang.X("submit form", "submit"):
				if submitButton != nil {
					t.Error("more than one submit button.")
				}
				submitButton = obj
			}
		}
	}

	return
}

func toolbarButtonGenerator(name string) iter.Seq[*widget.Button] {
	content := globals.ApplicationContext.UI().
		GetMainWindow().Content()
	laidOutObjects := fyneTest.LaidOutObjects(content)

	return func(yield func(*widget.Button) bool) {
		for _, laidOutObject := range laidOutObjects {
			switch obj := laidOutObject.(type) {
			case *widget.Toolbar:
				for _, item := range obj.Items {
					switch tbobj := item.(type) {
					case *widget.ToolbarAction:
						if tbobj.Icon.Name() != name {
							continue
						}
						pButton, ok := tbobj.
							ToolbarObject().(*widget.Button)
						if !ok {
							continue
						}
						if !yield(pButton) {
							return
						}
					}
				}
			}
		}
	}
}

func ExtractContentAddButton(t *testing.T) (pButton *widget.Button) {
	contentAddButtons := slices.Collect(
		toolbarButtonGenerator(
			theme.ContentAddIcon().Name(),
		),
	)

	if len(contentAddButtons) != 1 {
		t.Error("only one content add button expected.")
		return
	}

	pButton = contentAddButtons[0]

	return
}

func ExtractSettingsButton(t *testing.T) (pButton *widget.Button) {
	contentAddButtons := slices.Collect(
		toolbarButtonGenerator(
			theme.SettingsIcon().Name(),
		),
	)

	if len(contentAddButtons) != 1 {
		t.Error("only one settings button expected.")
		return
	}

	pButton = contentAddButtons[0]

	return
}

func ExtractPlayButton(t *testing.T) (pButton *widget.Button) {
	mediaPlayButtons := slices.Collect(
		toolbarButtonGenerator(
			theme.MediaPlayIcon().Name(),
		),
	)

	if len(mediaPlayButtons) != 1 {
		t.Error("only one play button expected.")
		return
	}

	pButton = mediaPlayButtons[0]

	return
}

func ExtractPauseButton(t *testing.T) (pButton *widget.Button) {
	mediaPauseButtons := slices.Collect(
		toolbarButtonGenerator(
			theme.MediaPauseIcon().Name(),
		),
	)

	if len(mediaPauseButtons) != 1 {
		t.Error("only one pause button expected.")
		return
	}

	pButton = mediaPauseButtons[0]

	return
}

func ExtractStopButton(t *testing.T) (pButton *widget.Button) {
	mediaStopButtons := slices.Collect(
		toolbarButtonGenerator(
			theme.MediaStopIcon().Name(),
		),
	)

	if len(mediaStopButtons) != 1 {
		t.Error("only one stop button expected.")
		return
	}

	pButton = mediaStopButtons[0]

	return
}

func ExtractSkipNextButton(t *testing.T) (pButton *widget.Button) {
	mediaSkipNextButtons := slices.Collect(
		toolbarButtonGenerator(
			theme.MediaSkipNextIcon().Name(),
		),
	)

	if len(mediaSkipNextButtons) != 1 {
		t.Error("only one skip next button expected.")
		return
	}

	pButton = mediaSkipNextButtons[0]

	return
}

func ExtractSkipPreviousButton(t *testing.T) (pButton *widget.Button) {
	mediaSkipPreviousButtons := slices.Collect(
		toolbarButtonGenerator(
			theme.MediaSkipPreviousIcon().Name(),
		),
	)

	if len(mediaSkipPreviousButtons) != 1 {
		t.Error("only one skip previous button expected.")
		return
	}

	pButton = mediaSkipPreviousButtons[0]

	return
}
