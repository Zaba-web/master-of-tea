package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	ROW_SPAN_0  = 0
	ROW_SPAN_1  = 0
	COL_SPAN_0  = 0
	COL_SPAN_1  = 0
	ROW_ID_0    = 0
	ROW_ID_1    = 1
	ROW_ID_2    = 2
	COL_ID_0    = 0
	COL_ID_1    = 1
	COL_ID_2    = 2
	FOCUSED     = true
	NOT_FOCUSED = false
)

var tuiApp = tview.NewApplication()
var pages = tview.NewPages()

func InitTui() {
	setupEvents()
	setupHomepage(pages)

	tuiApp.SetRoot(pages, true).SetFocus(pages).Run()
}

func setupEvents() {
	tuiApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			tuiApp.Stop()
		}
		return event
	})
}
