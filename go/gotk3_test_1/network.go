// network.go

// Network Settings and Protocol Rules.

//=============================================================================|

package main

//=============================================================================|

const MESSAGE_AUX_FIELDS_SIZE = 30

const MESSAGE_SIZE_MAX = MESSAGE_TEXT_SIZE_MAX + MESSAGE_AUX_FIELDS_SIZE
const MESSAGE_SIZE_MIN = MESSAGE_AUX_FIELDS_SIZE
const MESSAGE_TEXT_SIZE_MAX = 65535

const MESSAGE_POSTFIX uint16 = MESSAGE_PREFIX
const MESSAGE_PREFIX uint16 = 63903 // Binary Square in Big Endian.

const NETWORK_PROTOCOL = "tcp4"

//=============================================================================|
