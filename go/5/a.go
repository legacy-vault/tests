// a.go

package main

import (
	"log"

	//"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

//-----------------------------------------------------------------------------|

var builder *gtk.Builder
var errg *glib.Error
var i uint
var window_main *gtk.Window

//-----------------------------------------------------------------------------|

func main() {

	gtk.Init(nil)

	builder = gtk.NewBuilder()
	i, errg = builder.AddFromFile("1.glade")
	log.Println(i, errg) ///

	/*
		// Main Window.
		window_main = gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
		window_main.SetPosition(gtk.WIN_POS_CENTER)
		window_main.SetTitle("Main Window")
		//window_main.SetIconName("gtk-dialog-info")
		pixbuf, err_glib = gdkpixbuf.NewPixbufFromFile("face-smile.png")
		window_main.SetIcon(pixbuf)

		window_main.Connect("destroy", func() { quit() })
	*/

	//window_main.SetSizeRequest(320, 240)
	//window_main.ShowAll()

	gtk.Main()
}

//-----------------------------------------------------------------------------|

func quit() {

	gtk.MainQuit()
}

//-----------------------------------------------------------------------------|

func error_check(err error) {

	if err != nil {
		log.Println(err)
	}
}

//-----------------------------------------------------------------------------|
