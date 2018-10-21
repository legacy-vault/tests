// a.go

package main

import (
	//"os"
	//"path"
	"log"

	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

//-----------------------------------------------------------------------------|

var acg *gtk.AccelGroup
var alig *gtk.Alignment
var cascademenu *gtk.MenuItem
var context_id uint
var err_glib *glib.Error
var frame *gtk.Frame
var lbl *gtk.Label
var menu *gtk.Menu
var menubar *gtk.MenuBar
var menuitem *gtk.ImageMenuItem
var pixbuf *gdkpixbuf.Pixbuf
var statusbar *gtk.Statusbar
var vbox, vbox_2 *gtk.VBox
var vpaned *gtk.VPaned
var window_main *gtk.Window

//-----------------------------------------------------------------------------|

func main() {

	gtk.Init(nil)

	// Main Window.
	window_main = gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window_main.SetPosition(gtk.WIN_POS_CENTER)
	window_main.SetTitle("Main Window")
	//window_main.SetIconName("gtk-dialog-info")
	pixbuf, err_glib = gdkpixbuf.NewPixbufFromFile("face-smile.png")
	window_main.SetIcon(pixbuf)

	window_main.Connect("destroy", func() { quit() })

	// VBox.
	vbox = gtk.NewVBox(false, 1)

	// Menu.
	menubar = gtk.NewMenuBar()
	vbox.PackStart(menubar, false, false, 0)
	// Menu 'Program'.
	cascademenu = gtk.NewMenuItemWithLabel("Program")

	menubar.Append(cascademenu)
	menu = gtk.NewMenu()
	cascademenu.SetSubmenu(menu)
	//menuitem = gtk.NewMenuItemWithLabel("Quit")
	acg = gtk.NewAccelGroup()
	menuitem = gtk.NewImageMenuItemFromStock(gtk.STOCK_CLOSE, acg)
	menuitem.Connect("activate", func() { quit() })
	menu.Append(menuitem)

	vpaned = gtk.NewVPaned()
	vbox.Add(vpaned)

	frame = gtk.NewFrame("Frame")
	vbox_2 = gtk.NewVBox(false, 5)
	frame.Add(vbox_2)
	vpaned.Pack1(frame, false, false)

	// Status Bar.
	statusbar = gtk.NewStatusbar()
	context_id = statusbar.GetContextId("go-gtk")
	statusbar_text_set("abc")
	vbox.PackStart(statusbar, false, false, 0)

	//window_main.Add(vbox)

	alig = gtk.NewAlignment(0, 1, 0, 0)
	lbl = gtk.NewLabel("Test Label!")
	alig.Add(lbl)
	window_main.Add(alig)

	window_main.SetSizeRequest(320, 240)
	window_main.ShowAll()

	gtk.Main()

	/*
		//--------------------------------------------------------
		// GtkVPaned
		//--------------------------------------------------------


		//--------------------------------------------------------
		// GtkFrame
		//--------------------------------------------------------


		frame2 := gtk.NewFrame("Demo")
		framebox2 := gtk.NewVBox(false, 1)
		frame2.Add(framebox2)


		vpaned.Pack2(frame2, false, false)

		//--------------------------------------------------------
		// GtkImage
		//--------------------------------------------------------
		dir, _ := path.Split(os.Args[0])
		imagefile := path.Join(dir, "../../data/go-gtk-logo.png")

		label := gtk.NewLabel("Go Binding for GTK")
		label.ModifyFontEasy("DejaVu Serif 15")
		framebox1.PackStart(label, false, true, 0)

		//--------------------------------------------------------
		// GtkEntry
		//--------------------------------------------------------
		entry := gtk.NewEntry()
		entry.SetText("Hello world")
		framebox1.Add(entry)

		image := gtk.NewImageFromFile(imagefile)
		framebox1.Add(image)

		//--------------------------------------------------------
		// GtkScale
		//--------------------------------------------------------
		scale := gtk.NewHScaleWithRange(0, 100, 1)
		scale.Connect("value-changed", func() {
			println("scale:", int(scale.GetValue()))
		})
		framebox2.Add(scale)

		//--------------------------------------------------------
		// GtkHBox
		//--------------------------------------------------------
		buttons := gtk.NewHBox(false, 1)

		//--------------------------------------------------------
		// GtkButton
		//--------------------------------------------------------
		button := gtk.NewButtonWithLabel("Button with label")
		button.Clicked(func() {
			println("button clicked:", button.GetLabel())
			messagedialog := gtk.NewMessageDialog(
				button.GetTopLevelAsWindow(),
				gtk.DIALOG_MODAL,
				gtk.MESSAGE_INFO,
				gtk.BUTTONS_OK,
				entry.GetText())
			messagedialog.Response(func() {
				println("Dialog OK!")

				//--------------------------------------------------------
				// GtkFileChooserDialog
				//--------------------------------------------------------
				filechooserdialog := gtk.NewFileChooserDialog(
					"Choose File...",
					button.GetTopLevelAsWindow(),
					gtk.FILE_CHOOSER_ACTION_OPEN,
					gtk.STOCK_OK,
					gtk.RESPONSE_ACCEPT)
				filter := gtk.NewFileFilter()
				filter.AddPattern("*.go")
				filechooserdialog.AddFilter(filter)
				filechooserdialog.Response(func() {
					println(filechooserdialog.GetFilename())
					filechooserdialog.Destroy()
				})
				filechooserdialog.Run()
				messagedialog.Destroy()
			})
			messagedialog.Run()
		})
		buttons.Add(button)

		//--------------------------------------------------------
		// GtkFontButton
		//--------------------------------------------------------
		fontbutton := gtk.NewFontButton()
		fontbutton.Connect("font-set", func() {
			println("title:", fontbutton.GetTitle())
			println("fontname:", fontbutton.GetFontName())
			println("use_size:", fontbutton.GetUseSize())
			println("show_size:", fontbutton.GetShowSize())
		})
		buttons.Add(fontbutton)
		framebox2.PackStart(buttons, false, false, 0)

		buttons = gtk.NewHBox(false, 1)

		//--------------------------------------------------------
		// GtkToggleButton
		//--------------------------------------------------------
		togglebutton := gtk.NewToggleButtonWithLabel("ToggleButton with label")
		togglebutton.Connect("toggled", func() {
			if togglebutton.GetActive() {
				togglebutton.SetLabel("ToggleButton ON!")
			} else {
				togglebutton.SetLabel("ToggleButton OFF!")
			}
		})
		buttons.Add(togglebutton)

		//--------------------------------------------------------
		// GtkCheckButton
		//--------------------------------------------------------
		checkbutton := gtk.NewCheckButtonWithLabel("CheckButton with label")
		checkbutton.Connect("toggled", func() {
			if checkbutton.GetActive() {
				checkbutton.SetLabel("CheckButton CHECKED!")
			} else {
				checkbutton.SetLabel("CheckButton UNCHECKED!")
			}
		})
		buttons.Add(checkbutton)

		//--------------------------------------------------------
		// GtkRadioButton
		//--------------------------------------------------------
		buttonbox := gtk.NewVBox(false, 1)
		radiofirst := gtk.NewRadioButtonWithLabel(nil, "Radio1")
		buttonbox.Add(radiofirst)
		buttonbox.Add(gtk.NewRadioButtonWithLabel(radiofirst.GetGroup(), "Radio2"))
		buttonbox.Add(gtk.NewRadioButtonWithLabel(radiofirst.GetGroup(), "Radio3"))
		buttons.Add(buttonbox)
		//radiobutton.SetMode(false);
		radiofirst.SetActive(true)

		framebox2.PackStart(buttons, false, false, 0)

		//--------------------------------------------------------
		// GtkVSeparator
		//--------------------------------------------------------
		vsep := gtk.NewVSeparator()
		framebox2.PackStart(vsep, false, false, 0)

		//--------------------------------------------------------
		// GtkComboBoxEntry
		//--------------------------------------------------------
		combos := gtk.NewHBox(false, 1)
		comboboxentry := gtk.NewComboBoxEntryNewText()
		comboboxentry.AppendText("Monkey")
		comboboxentry.AppendText("Tiger")
		comboboxentry.AppendText("Elephant")
		comboboxentry.Connect("changed", func() {
			println("value:", comboboxentry.GetActiveText())
		})
		combos.Add(comboboxentry)

		//--------------------------------------------------------
		// GtkComboBox
		//--------------------------------------------------------
		combobox := gtk.NewComboBoxNewText()
		combobox.AppendText("Peach")
		combobox.AppendText("Banana")
		combobox.AppendText("Apple")
		combobox.SetActive(1)
		combobox.Connect("changed", func() {
			println("value:", combobox.GetActiveText())
		})
		combos.Add(combobox)

		framebox2.PackStart(combos, false, false, 0)

		//--------------------------------------------------------
		// GtkTextView
		//--------------------------------------------------------
		swin := gtk.NewScrolledWindow(nil, nil)
		swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
		swin.SetShadowType(gtk.SHADOW_IN)
		textview := gtk.NewTextView()
		var start, end gtk.TextIter
		buffer := textview.GetBuffer()
		buffer.GetStartIter(&start)
		buffer.Insert(&start, "Hello ")
		buffer.GetEndIter(&end)
		buffer.Insert(&end, "World!")
		tag := buffer.CreateTag("bold", map[string]string{
			"background": "#FF0000", "weight": "700"})
		buffer.GetStartIter(&start)
		buffer.GetEndIter(&end)
		buffer.ApplyTag(tag, &start, &end)
		swin.Add(textview)
		framebox2.Add(swin)

		buffer.Connect("changed", func() {
			println("changed")
		})

		//--------------------------------------------------------
		// GtkMenuItem
		//--------------------------------------------------------
		cascademenu := gtk.NewMenuItemWithMnemonic("_File")
		menubar.Append(cascademenu)
		submenu := gtk.NewMenu()
		cascademenu.SetSubmenu(submenu)

		menuitem = gtk.NewMenuItemWithMnemonic("E_xit")
		menuitem.Connect("activate", func() {
			gtk.MainQuit()
		})
		submenu.Append(menuitem)

		cascademenu = gtk.NewMenuItemWithMnemonic("_View")
		menubar.Append(cascademenu)
		submenu = gtk.NewMenu()
		cascademenu.SetSubmenu(submenu)

		checkmenuitem := gtk.NewCheckMenuItemWithMnemonic("_Disable")
		checkmenuitem.Connect("activate", func() {
			vpaned.SetSensitive(!checkmenuitem.GetActive())
		})
		submenu.Append(checkmenuitem)

		menuitem = gtk.NewMenuItemWithMnemonic("_Font")
		menuitem.Connect("activate", func() {
			fsd := gtk.NewFontSelectionDialog("Font")
			fsd.SetFontName(fontbutton.GetFontName())
			fsd.Response(func() {
				println(fsd.GetFontName())
				fontbutton.SetFontName(fsd.GetFontName())
				fsd.Destroy()
			})
			fsd.SetTransientFor(window_main)
			fsd.Run()
		})
		submenu.Append(menuitem)

		cascademenu = gtk.NewMenuItemWithMnemonic("_Help")
		menubar.Append(cascademenu)
		submenu = gtk.NewMenu()
		cascademenu.SetSubmenu(submenu)

		menuitem = gtk.NewMenuItemWithMnemonic("_About")
		menuitem.Connect("activate", func() {
			dialog := gtk.NewAboutDialog()
			dialog.SetName("Go-Gtk Demo!")
			dialog.SetProgramName("demo")
			//dialog.SetAuthors(authors())
			dir, _ := path.Split(os.Args[0])
			imagefile := path.Join(dir, "../../data/mattn-logo.png")
			pixbuf, _ := gdkpixbuf.NewPixbufFromFile(imagefile)
			dialog.SetLogo(pixbuf)
			dialog.SetLicense("The library is available under the same terms and conditions as the Go, the BSD style license, and the LGPL (Lesser GNU Public License). The idea is that if you can use Go (and Gtk) in a project, you should also be able to use go-gtk.")
			dialog.SetWrapLicense(true)
			dialog.Run()
			dialog.Destroy()
		})
		submenu.Append(menuitem)
	*/

	//--------------------------------------------------------

	//framebox2.PackStart(statusbar, false, false, 0)
}

//-----------------------------------------------------------------------------|

func statusbar_text_set(text string) {

	statusbar.Push(context_id, text)
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
