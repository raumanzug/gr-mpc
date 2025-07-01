package report

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type BackingData struct {
	accu          string
	pErrorLabel   *widget.Label
	ClearCallback func() // called when report box is cleared.
}

// AddErr adds an error to report box.
//
// Call it in fyne's thread only!
func (pBD *BackingData) AddErr(err error) {
	if err != nil {
		pBD.accu += err.Error() + "\n"
		pBD.pErrorLabel.SetText(pBD.accu)
	}
}

// Clear clears the report box.
//
// Call it in fyne's thread only!
func (pBD *BackingData) Clear() {
	pBD.accu = ""
	pBD.pErrorLabel.SetText(pBD.accu)
	if pBD.ClearCallback != nil {
		pBD.ClearCallback()
	}
}

// IsEmpty checks whether report box is empty.
//
// Call it in fyne's thread only!
func (pBD *BackingData) IsEmpty() bool {
	return pBD.accu == ""
}

// NewBackingData creates MView object for report box view generation.
//
// Call it in fyne's thread only!
func NewBackingData() *BackingData {
	bd := BackingData{
		accu:        "",
		pErrorLabel: widget.NewLabel(""),
	}
	bd.pErrorLabel.Wrapping = fyne.TextWrapWord

	return &bd
}
