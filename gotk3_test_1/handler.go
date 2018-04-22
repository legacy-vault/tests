// handler.go

// Handlers of Events.

//=============================================================================|

package main

//=============================================================================|

//import ""

//=============================================================================|

// 'Quit Application' Button Click Handler.
func btn_app_quit_clicked() {

	// Start Asynchronous Handler.
	go btn_app_quit_clicked_async()
}

//=============================================================================|

// 'Quit Application' Button Click Asynchronous Handler.
func btn_app_quit_clicked_async() {

	var err error

	// When we click on 'Close' Button in G.U.I.,
	// the Handler stops the Server if it is running,
	// then it sends a Signal to U.I. Manager (Loop),
	// U.I. Manager closes User Interface and sends a Signal to the
	// 'AppQuit' Manager, who then closes the main Program's Loop.

	// Disable 'Quit' Button.
	btnQuit.SetSensitive(false)

	// If Server is working...
	if serverIsWorking {

		// Stop the Server.
		err = server_stop()
		if err != nil {
			ui_console_msg_add(err.Error())

			// Enable 'Quit' Button.
			btnQuit.SetSensitive(true)

			return
		}

		// Enable 'On' Button.
		btnServerOn.SetSensitive(true)

	}

	// Start the Process of Application's Termination.
	app_quit()
}

//=============================================================================|

// 'Client Message Send' Button Click Handler.
func btn_client_message_send_clicked() {

	// Start Asynchronous Handler.
	go btn_client_message_send_clicked_async()
}

//=============================================================================|

// 'Client Message Send' Button Click Asynchronous Handler.
func btn_client_message_send_clicked_async() {

	var err error

	// Disable 'Send' Button.
	btnClientMessageSend.SetSensitive(false)

	// Prepare and send the Message.
	err = client_message_send()

	if err != nil {
		ui_console_msg_add(err.Error())

		// Enable 'Send' Button.
		btnClientMessageSend.SetSensitive(true)

		return
	}

	// Enable 'Send' Button.
	btnClientMessageSend.SetSensitive(true)
}

//=============================================================================|

// 'Server Off' Button Click Handler.
func btn_server_off_clicked() {

	// Start Asynchronous Handler.
	go btn_server_off_clicked_async()
}

//=============================================================================|

// 'Server Off' Button Click Asynchronous Handler.
func btn_server_off_clicked_async() {

	var err error

	// Disable 'Off' Button.
	btnServerOff.SetSensitive(false)

	// If Server is stopped...
	if !serverIsWorking {
		return
	}

	// If Server is working...

	// Stop the Server.
	err = server_stop()
	if err != nil {
		ui_console_msg_add(err.Error())

		// Enable 'Off' Button.
		btnServerOff.SetSensitive(true)

		return
	}

	// Enable 'On' Button.
	btnServerOn.SetSensitive(true)
}

//=============================================================================|

// 'Server On' Button Click Handler.
func btn_server_on_clicked() {

	// Start Asynchronous Handler.
	go btn_server_on_clicked_async()
}

//=============================================================================|

// 'Server On' Button Click Asynchronous Handler.
func btn_server_on_clicked_async() {

	var err error

	// Disable 'On' Button.
	btnServerOn.SetSensitive(false)

	// Read Server Port.
	err = config_server_port_read()
	if err != nil {
		ui_console_msg_add(err.Error())

		// Enable 'On' Button.
		btnServerOn.SetSensitive(true)

		return
	}

	// If Server is already working...
	if serverIsWorking {
		return
	}

	// If Server is not yet working...

	// Start the Server.
	err = server_start()
	if err != nil {
		ui_console_msg_add(err.Error())

		// Enable 'On' Button.
		btnServerOn.SetSensitive(true)

		return
	}

	// Enable 'Off' Button.
	btnServerOff.SetSensitive(true)
}

//=============================================================================|
