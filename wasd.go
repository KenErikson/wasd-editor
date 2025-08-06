package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	textView := tview.NewTextView()
	textView.SetBackgroundColor(tcell.NewRGBColor(30, 29, 30))
	textView.SetText("AAAAAAAAAAAAAAAA\nBBBBBBBBBBBBBB\nCCCCCC\nA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCCA\nBBBBBBBBBBBBBB\nCCCCCC")
	textView.SetTextColor(tcell.NewRGBColor(118, 189, 217))

	rowNumbers := tview.NewTextView()
	rowNumbers.SetTextAlign(tview.AlignCenter).SetText("~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n~\n")
	rowNumbers.SetBackgroundColor(tcell.NewRGBColor(30, 29, 30))
	rowNumbers.SetTextColor(tcell.NewRGBColor(128, 128, 128))

	statusBar := tview.NewTextView()
	statusBar.SetBackgroundColor(tcell.NewRGBColor(36, 35, 36))
	statusBar.SetTextColor(tcell.NewRGBColor(212, 212, 212))
	statusBar.SetText("Status: Fire").SetTextAlign(tview.AlignRight)

	gridView := tview.NewGrid().SetColumns(2, 0).SetRows(0, 1)
	gridView.SetBorderPadding(0, 0, 0, 0)
	gridView.AddItem(rowNumbers, 0, 0, 1, 1, 0, 0, false)
	gridView.AddItem(textView, 0, 1, 1, 1, 0, 0, true)
	gridView.AddItem(statusBar, 1, 0, 1, 2, 0, 0, false)

	if err := app.SetRoot(gridView, true).SetFocus(gridView).Run(); err != nil {
		panic(err)
	}
}
