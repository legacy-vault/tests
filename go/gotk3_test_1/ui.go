// ui.go

// User Interface Functions and Objects.

//=============================================================================|

package main

//=============================================================================|

import "log"
import "errors"
import "time"
import "github.com/gotk3/gotk3/gdk"
import "github.com/gotk3/gotk3/glib"
import "github.com/gotk3/gotk3/gtk"

//=============================================================================|

// User Interface Configuration File.
const UI_FILE_PATH = "./ui/main.glade"
const UI_FILE_CHECKSUM_CRC64_HEX = "F05607A985FC8793"
const UI_FILE_CHECKSUM_MD5_HEX = "08693E70B8B1209B5BF5B06241145CF7"

// Signals.
const UI_MANAGER_SIGNAL_STOP = true

// Manager Configuration.
const UI_MANAGER_WAIT_SIGNAL_INTERVAL = 2

// Main Window Settings.
const UI_WINDOW_ID = "window_main"
const UI_WINDOW_TITLE = "Test"

// Buttons.
const UI_BTN_CLIENT_MESSAGE_SEND_ID = "input_client_send"
const UI_BTN_SERVER_OFF_ID = "input_server_off"
const UI_BTN_SERVER_ON_ID = "input_server_on"
const UI_BTN_QUIT_ID = "btn_quit"

// Console.
const UI_CONSOLE_ID = "output_console"
const UI_CONSOLE_TIME_FORMAT = "15:04:05"
const UI_CONSOLE_BRACKET_LEFT = "["
const UI_CONSOLE_BRACKET_RIGHT = "]"
const UI_CONSOLE_BRACKET_SPACER = " "

// Images.

// 1. LED OFF.
const UI_IMAGE_LED_OFF_ACTIVE_FILE_PATH = "./img/led_off_active.png"
const UI_IMAGE_LED_OFF_PASSIVE_FILE_PATH = "./img/led_off_passive.png"
const UI_IMAGE_LED_OFF_TRANSIT_FILE_PATH = "./img/led_off_transit.png"

// 1. LED ON.
const UI_IMAGE_LED_ON_ACTIVE_FILE_PATH = "./img/led_on_active.png"
const UI_IMAGE_LED_ON_PASSIVE_FILE_PATH = "./img/led_on_passive.png"
const UI_IMAGE_LED_ON_TRANSIT_FILE_PATH = "./img/led_on_transit.png"

// Indicators.
const UI_INDICATOR_SERVER_OFF_ID = "output_server_off"
const UI_INDICATOR_SERVER_ON_ID = "output_server_on"

// Input Fields.
const UI_INPUT_CLIENT_HOST_ID = "input_client_host"
const UI_INPUT_CLIENT_MESSAGE_ID = "input_client_message"
const UI_INPUT_CLIENT_PORT_ID = "input_client_port"
const UI_INPUT_SERVER_PORT_ID = "input_server_port"

//=============================================================================|

// User Interface Objects.
var btnClientMessageSend *gtk.Button
var btnServerOff *gtk.Button
var btnServerOn *gtk.Button
var btnQuit *gtk.ToolButton

var gtkBuilder *gtk.Builder
var gtkConsole *gtk.TextView
var gtkConsoleBuffer *gtk.TextBuffer
var gtkWindow *gtk.Window

var indicatorServerOff *gtk.Image
var indicatorServerOn *gtk.Image

var inputClientHost *gtk.Entry
var inputClientHostBuffer *gtk.EntryBuffer
var inputClientMessage *gtk.TextView
var inputClientMessageBuffer *gtk.TextBuffer
var inputClientPort *gtk.Entry
var inputClientPortBuffer *gtk.EntryBuffer
var inputServerPort *gtk.Entry

var pixbufLEDOffActive *gdk.Pixbuf
var pixbufLEDOffPassive *gdk.Pixbuf
var pixbufLEDOffTransit *gdk.Pixbuf

var pixbufLEDOnActive *gdk.Pixbuf
var pixbufLEDOnPassive *gdk.Pixbuf
var pixbufLEDOnTransit *gdk.Pixbuf

// Flags.
var uiHasStopped bool

// Channels.
var uiManagerChan chan bool

//=============================================================================|

// Adds the Message to the Console.
func ui_console_msg_add(msg string) {

	var msg_full string
	var textIter *gtk.TextIter

	// Prepare Message.
	msg_full = UI_CONSOLE_BRACKET_LEFT +
		time.Now().Format(UI_CONSOLE_TIME_FORMAT) +
		UI_CONSOLE_BRACKET_RIGHT +
		UI_CONSOLE_BRACKET_SPACER +
		msg + NL

	// Save internally.
	consoleText = consoleText + msg_full

	// Visualize Changes.
	textIter = gtkConsoleBuffer.GetEndIter()
	gtkConsoleBuffer.Insert(textIter, msg_full)

	// Scroll to the Bottom.
	textIter = gtkConsoleBuffer.GetEndIter()
	gtkConsole.ScrollToIter(textIter, 0, true, 0.5, 0.5)

	//! Scrolling does not work! GTK is bugged ?!

	// Verbose Report.
	if verbose {
		log.Println(msg)
	}
}

//=============================================================================|

// Binds U.I. Objects with Actions.
func ui_bind_actions() error {

	// Notes:
	//
	//	1.	When a Handler Function runs, GTK waits (holds) until the Handler
	//		finishes its Job. The graphical Interface freezes all the time while
	//		it waits. So, we need to start Handlers in an asynchronous Way. So,
	//		these Handlers are only the Wrappers for asynchronous Calls to real
	//		Handlers.

	var err error

	// 1. Application Exit (Cross Button).
	_, err = gtkWindow.Connect("destroy", btn_app_quit_clicked)
	if err != nil {
		return err
	}

	// 2. Exit Button on Toolbar.
	_, err = btnQuit.Connect("clicked", btn_app_quit_clicked)
	if err != nil {
		return err
	}

	// 3. 'Server On' Button.
	_, err = btnServerOn.Connect("clicked", btn_server_on_clicked)
	if err != nil {
		return err
	}

	// 3. 'Server Off' Button.
	_, err = btnServerOff.Connect("clicked", btn_server_off_clicked)
	if err != nil {
		return err
	}

	// 4. 'Client Message Send' Button.
	_, err = btnClientMessageSend.Connect("clicked",
		btn_client_message_send_clicked)
	if err != nil {
		return err
	}

	return nil
}

//=============================================================================|

// Binds U.I. Buttons with Objects.
func ui_get_buttons() error {

	var err error
	var object glib.IObject
	var ok bool

	// 1. Quit.
	object, err = gtkBuilder.GetObject(UI_BTN_QUIT_ID)
	if err != nil {
		return err
	}
	btnQuit, ok = object.(*gtk.ToolButton)
	if !ok {
		err = errors.New(MSG_ERR_1)
		return err
	}

	// 2. 'Server On' Button.
	object, err = gtkBuilder.GetObject(UI_BTN_SERVER_ON_ID)
	if err != nil {
		return err
	}
	btnServerOn, ok = object.(*gtk.Button)
	if !ok {
		err = errors.New(MSG_ERR_1)
		return err
	}

	// 3. 'Server Off' Button.
	object, err = gtkBuilder.GetObject(UI_BTN_SERVER_OFF_ID)
	if err != nil {
		return err
	}
	btnServerOff, ok = object.(*gtk.Button)
	if !ok {
		err = errors.New(MSG_ERR_1)
		return err
	}
	btnServerOff.SetSensitive(false)

	// 4. 'Client Message Send' Button.
	object, err = gtkBuilder.GetObject(UI_BTN_CLIENT_MESSAGE_SEND_ID)
	if err != nil {
		return err
	}
	btnClientMessageSend, ok = object.(*gtk.Button)
	if !ok {
		err = errors.New(MSG_ERR_1)
		return err
	}

	return nil
}

//=============================================================================|

// Binds U.I. Indicators with Objects.
func ui_get_indicators() error {

	var err error
	var object glib.IObject
	var ok bool

	// Prepare Images.

	// 1. LED OFF.
	pixbufLEDOffActive, err =
		gdk.PixbufNewFromFile(UI_IMAGE_LED_OFF_ACTIVE_FILE_PATH)
	if err != nil {
		return err
	}
	pixbufLEDOffPassive, err =
		gdk.PixbufNewFromFile(UI_IMAGE_LED_OFF_PASSIVE_FILE_PATH)
	if err != nil {
		return err
	}
	pixbufLEDOffTransit, err =
		gdk.PixbufNewFromFile(UI_IMAGE_LED_OFF_TRANSIT_FILE_PATH)
	if err != nil {
		return err
	}

	// 2. LED ON.
	pixbufLEDOnActive, err =
		gdk.PixbufNewFromFile(UI_IMAGE_LED_ON_ACTIVE_FILE_PATH)
	if err != nil {
		return err
	}
	pixbufLEDOnPassive, err =
		gdk.PixbufNewFromFile(UI_IMAGE_LED_ON_PASSIVE_FILE_PATH)
	if err != nil {
		return err
	}
	pixbufLEDOnTransit, err =
		gdk.PixbufNewFromFile(UI_IMAGE_LED_ON_TRANSIT_FILE_PATH)
	if err != nil {
		return err
	}

	// Prepare Indicators.

	// 1. Server On LED.
	object, err =
		gtkBuilder.GetObject(UI_INDICATOR_SERVER_ON_ID)
	if err != nil {
		return err
	}
	indicatorServerOn, ok = object.(*gtk.Image)
	if !ok {
		err = errors.New(MSG_ERR_1)
		return err
	}
	indicatorServerOn.SetFromPixbuf(pixbufLEDOnPassive)

	// 2. Server Off LED.
	object, err = gtkBuilder.GetObject(UI_INDICATOR_SERVER_OFF_ID)
	if err != nil {
		return err
	}
	indicatorServerOff, ok = object.(*gtk.Image)
	if !ok {
		err = errors.New(MSG_ERR_1)
		return err
	}
	indicatorServerOff.SetFromPixbuf(pixbufLEDOffActive)

	return nil
}

//=============================================================================|

// Binds U.I. Text Inputs with Objects.
func ui_get_inputs() error {

	var err error
	var object glib.IObject
	var ok bool

	// 1. Server Port.
	object, err = gtkBuilder.GetObject(UI_INPUT_SERVER_PORT_ID)
	if err != nil {
		return err
	}
	inputServerPort, ok = object.(*gtk.Entry)
	if !ok {
		err = errors.New(MSG_ERR_1)
		return err
	}

	// 2. Client Host.
	object, err = gtkBuilder.GetObject(UI_INPUT_CLIENT_HOST_ID)
	if err != nil {
		return err
	}
	inputClientHost, ok = object.(*gtk.Entry)
	if !ok {
		err = errors.New(MSG_ERR_1)
		return err
	}
	inputClientHostBuffer, err = inputClientHost.GetBuffer()
	if err != nil {
		return err
	}

	// 3. Client Port.
	object, err = gtkBuilder.GetObject(UI_INPUT_CLIENT_PORT_ID)
	if err != nil {
		return err
	}
	inputClientPort, ok = object.(*gtk.Entry)
	if !ok {
		err = errors.New(MSG_ERR_1)
		return err
	}
	inputClientPortBuffer, err = inputClientPort.GetBuffer()
	if err != nil {
		return err
	}

	// 4. Client Message.
	object, err = gtkBuilder.GetObject(UI_INPUT_CLIENT_MESSAGE_ID)
	if err != nil {
		return err
	}
	inputClientMessage, ok = object.(*gtk.TextView)
	if !ok {
		err = errors.New(MSG_ERR_1)
		return err
	}
	inputClientMessageBuffer, err = inputClientMessage.GetBuffer()
	if err != nil {
		return err
	}

	return nil
}

//=============================================================================|

// Initializes the User Interface.
func ui_init() error {

	var err error
	var object glib.IObject
	var ok bool
	var ui_file_path string

	// Initialize GTK.
	gtk.Init(nil)

	// Create GTK Builder.
	gtkBuilder, err = gtk.BuilderNew()
	if err != nil {
		return err
	}

	//  Load U.I. from File.
	ui_file_path = UI_FILE_PATH
	err = gtkBuilder.AddFromFile(ui_file_path)
	if err != nil {
		return err
	}

	// Get main Window.
	object, err = gtkBuilder.GetObject(UI_WINDOW_ID)
	if err != nil {
		return err
	}
	gtkWindow, ok = object.(*gtk.Window)
	if !ok {
		err = errors.New(MSG_ERR_1)
		return err
	}

	// Get Console.
	object, err = gtkBuilder.GetObject(UI_CONSOLE_ID)
	if err != nil {
		return err
	}
	gtkConsole, ok = object.(*gtk.TextView)
	if !ok {
		err = errors.New(MSG_ERR_1)
		return err
	}
	gtkConsoleBuffer, err = gtkConsole.GetBuffer()
	if err != nil {
		return err
	}

	// Get Buttons.
	err = ui_get_buttons()
	if err != nil {
		return err
	}

	// Get Input Fields.
	err = ui_get_inputs()
	if err != nil {
		return err
	}

	// Get Indicators.
	err = ui_get_indicators()
	if err != nil {
		return err
	}

	// Configure main Window.
	gtkWindow.SetTitle(UI_WINDOW_TITLE)

	// Connect (bind) Window's Objects with Actions.
	err = ui_bind_actions()
	if err != nil {
		return err
	}

	// Add Message to Console.
	ui_console_msg_add(MSG_1)

	// Show main Window.
	gtkWindow.ShowAll()

	return nil
}

//=============================================================================|

// User Interface Manager.
func ui_manager() {

	var loop bool
	var sleepPeriod time.Duration
	var signal bool

	sleepPeriod = time.Second * UI_MANAGER_WAIT_SIGNAL_INTERVAL
	uiHasStopped = false

	// Start GTK main Loop.
	go gtk.Main()

	// Wait for Signal to stop.
	loop = true
	for loop {

		select {

		case signal = <-uiManagerChan:

			// Signal is received.

			// Verbose Report.
			if verbose {
				log.Println(MSG_3)
			}

			if signal == UI_MANAGER_SIGNAL_STOP {
				loop = false
			}

		default:

			// No Signal is received.
			time.Sleep(sleepPeriod)
		}
	}

	gtk.MainQuit()

	log.Println(MSG_8)
	uiHasStopped = true

	// U.I. has stopped. Now we need to stop main Program's Loop.
	appManagerChan <- APP_MANAGER_SIGNAL_STOP
}

//=============================================================================|
