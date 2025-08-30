package tui

import (
	"github.com/rivo/tview"
)

const HOMEPAGE = "homepage"

func setupHomepage(pages *tview.Pages) {
	var grid *tview.Grid = tview.NewGrid().
		SetRows(0, 10).
		SetColumns(40, 0)

	grid.AddItem(
		getGeneralStatusComopnent(),
		ROW_ID_0,
		COL_ID_0,
		ROW_SPAN_1,
		COL_SPAN_1,
		1,
		1,
		NOT_FOCUSED,
	)

	grid.AddItem(
		tview.NewBox().SetTitle("test2").SetBorder(true),
		0,
		1,
		1,
		1,
		1,
		1,
		false,
	)

	grid.AddItem(
		tview.NewBox().SetTitle("test3").SetBorder(true),
		1,
		0,
		1,
		2,
		1,
		1,
		false,
	)

	pages.AddPage(HOMEPAGE, grid, true, true)
}

func getGeneralStatusComopnent() *tview.Grid {
	var generalStatusGrid *tview.Grid = tview.NewGrid().SetBorders(true).SetRows(0).SetColumns(0)

	return generalStatusGrid
}
