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
var gtkButton *gtk.Button
var gtkBuilder *gtk.Builder

var gtkWindowTitle string
var gtkLabelText string
var gtkButtonLabelText string
var gtkBuilderPath string

//------------------------------------------------------------------------------

func gtk_init() {

	var err error

	/*
		gtkWindowTitle = "GTK 3 Window"
		gtkLabelText = "It works!"
		gtkButtonLabelText = "Click Me!"
	*/
	//gtkBuilderPath = "/home/username/go/src/tests/gtk3/b/gui//gui.gresource"
	gtkBuilderPath = "gui/gui.gresource"

	// Initialize GTK without parsing any Command Line Arguments.
	gtk.Init(nil)

	gtkBuilder, err = gtk.BuilderNewFromResource(gtkBuilderPath)
	gtk_check_creation("Builder", err)

	/*

		// Window.
		gtkWindow, err = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
		gtk_check_creation("Window", err)
		gtkWindow.SetTitle(gtkWindowTitle)
		gtkWindow.Connect("destroy", gtk_handler_destroy)
		gtkWindow.SetDefaultSize(640, 480)

		// Label.
		gtkLabel, err = gtk.LabelNew(gtkLabelText)
		gtk_check_creation("Label", err)
		gtkWindow.Add(gtkLabel)

		// Button.
		gtkButton, err = gtk.ButtonNewWithLabel(gtkButtonLabelText)
		gtk_check_creation("Button", err)
		gtkWindow.

		// Show all Widgets in Window.
		gtkWindow.ShowAll()

	*/
}

//------------------------------------------------------------------------------

func gtk_handler_destroy() {

	// GTK 'destroy' Event Handler.

	gtk.MainQuit()
}

//------------------------------------------------------------------------------

func gtk_check_creation(objName string, err error) {

	// Checks if the Object has been created.
	// Prints fatal (final) Message on Error.

	if err != nil {

		log.Fatal("Error.", objName, "can not be created.", err) //
	}
}

//------------------------------------------------------------------------------
