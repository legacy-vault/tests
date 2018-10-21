package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

//------------------------------------------------------------------------------
// Variables
//------------------------------------------------------------------------------

var gtkWindow *gtk.Window
var gtkLabel *gtk.Label

//------------------------------------------------------------------------------

func main() {

	gtk_init()

	// GTK main Loop (blocks until 'gtk.MainQuit()' is run.
	gtk.Main()
}

//------------------------------------------------------------------------------

func gtk_init() {

	var err error

	// Initialize GTK without parsing any Command Line Arguments.
	gtk.Init(nil)

	gtkWindow, err = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create a Window:", err) //
	}

	gtkWindow.SetTitle("GTK 3 Window")

	gtkWindow.Connect("destroy", gtk_handler_destroy)

	// Label Widget.
	gtkLabel, err = gtk.LabelNew("It works!")
	if err != nil {
		log.Fatal("Unable to create a Label:", err) //
	}

	// Label -> Window.
	gtkWindow.Add(gtkLabel)

	// Window Size.
	gtkWindow.SetDefaultSize(640, 480)

	// Show all Widgets in Window.
	gtkWindow.ShowAll()
}

//------------------------------------------------------------------------------

func gtk_handler_destroy() {

	// GTK 'destroy' Event Handler.

	gtk.MainQuit()
}

//------------------------------------------------------------------------------
