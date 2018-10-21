// gtk.go

package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

//------------------------------------------------------------------------------
// Variables
//------------------------------------------------------------------------------

var gtkBuilder *gtk.Builder

var gtkBuilderPath string

//------------------------------------------------------------------------------

func gtk_init() {

	var err error

	gtkBuilderPath = "gui/gui.gresource"

	gtk.Init(nil)

	gtkBuilder, err = gtk.BuilderNewFromResource(gtkBuilderPath)
	gtk_check_creation("Builder", err)
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
