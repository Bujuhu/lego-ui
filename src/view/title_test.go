package view

import (
	"fmt"
	"testing"
	"fyne.io/fyne/v2/app"
)


func TestSetTitle(t *testing.T) {
	a := app.NewWithID("LEGO-Test")
	w := a.NewWindow("LEGO-Test")
	w.SetContent(Title("Test Title"))
	
	fmt.Printf("w.Content(): %v\n", w.Content())
}