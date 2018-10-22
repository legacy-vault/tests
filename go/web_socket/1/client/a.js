// a.js.

var ws;
var inputWsRequest;
var inputWsResponse;

// Handles 'load' Event of Body Tag.
function body_load_handler()
{
	objectsInit();
	webSocketInit();
}

function objectsInit()
{
	inputWsRequest = document.getElementById("ws_request");
	inputWsResponse = document.getElementById("ws_response");
}

function webSocketInit()
{
	wsAddress = "ws://localhost:2000/test";
	ws = new WebSocket(wsAddress);
	ws.onopen = wsOpenHandler;
	ws.onclose = wsCloseHandler;
	ws.onmessage = wsMessageHandler;
	ws.onerror = wsErrorHandler;
}

function wsOpenHandler()
{
	inputWsResponse.value += "WebSocket Connection has been established.\r\n";
}

function wsCloseHandler(event)
{
	inputWsResponse.value += "WebSocket Connection has been closed. " + 
	"Code: " + event.code + ", Reason: " + event.reason + ".\r\n";
}

function wsMessageHandler(event)
{
	inputWsResponse.value += "Received Message: [" + event.data + "].\r\n";
}

function wsErrorHandler(error)
{
	inputWsResponse.value += "Error: [" + error.message + "].\r\n";
}

function ws_btn_click()
{
	msg = inputWsRequest.value;
	inputWsResponse.value += "Sent Message: [" + msg + "].\r\n";
	ws.send(msg);
}
