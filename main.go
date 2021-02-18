package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
)

func callback(s string) {
	fmt.Printf("s[%s]\n", s)
}

func main() {
	a := app.New()
	w := a.NewWindow("Examples")
	d := dialog.NewEntryDialog("title", "message", callback, w)
	d.Show()
	w.Resize(fyne.Size{Width: 300, Height: 300})
	w.ShowAndRun()
}
