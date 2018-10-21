// a.go

package main

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

//-----------------------------------------------------------------------------|

func main() {

	var builder *gtk.Builder
	var err error
	var obj glib.IObject
	var ok bool
	var window *gtk.Window

	gtk.Init(nil)

	// Builder.
	builder, err = gtk.BuilderNew()
	errorCheck(err)
	err = builder.AddFromFile("1.glade")
	errorCheck(err)

	// Builder -> MainWindow.
	obj, err = builder.GetObject("MainWindow")
	errorCheck(err)
	log.Println("obj:", obj) ///
	window, ok = obj.(*gtk.Window)
	if !ok {
		log.Fatal("Object is not found.")
	}

	// Configure MainWindow.
	window.SetTitle("Example")
	window.Connect("destroy", func() { appQuit() })
	window.SetDefaultSize(640, 480)

	// Show All.
	window.ShowAll()

	gtk.Main()
}

//-----------------------------------------------------------------------------|

func errorCheck(err error) {

	if err != nil {
		log.Println(err)
	}
}

//-----------------------------------------------------------------------------|

func appQuit() {

	gtk.MainQuit()
	os.Exit(0)
}

//-----------------------------------------------------------------------------|
