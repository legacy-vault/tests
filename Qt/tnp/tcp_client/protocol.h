// protocol.h.

#ifndef PROTOCOL_H
#define PROTOCOL_H

class Protocol
{

public:
    // Constructor.
    explicit Protocol();

    // Destructor.
    ~Protocol();

    // Constants.
    static const char PROTOCOL_VERSION_MAJOR = 1;
    static const char PROTOCOL_VERSION_MINOR = 0;

    static const int CLIENT_CONNECTION_TIMEOUT = 120; // Seconds.

    static const int MESSAGE_FIELD_SIZE_LEN = 3;
    // 1.   Field 'Size' is read into uint64, so this const should be <= 8.
    // 2.   Moreover, 'QTcpSocket::bytesAvailable' Function returns only int64,
    //      so this const should be <= 7.
    // 3.   While our Message Size is limited to 65535 Bytes, const Value of 2
    //      Bytes may be enough.
    // 4.   Finally, this Size is fixed in Protocol and implemented both in
    //      Client and Server, so let it be 3, having the first Byte reserved.

    static const int MESSAGE_TEXT_BYTES_COUNT_MAX = 65535; // uint16.
    static const int MESSAGE_IDLE_TIMEOUT = 120; // Seconds.

    static const char SERVER_REPLY_ERROR = 0;      // BIN (BE): 00000000.
    static const char SERVER_REPLY_SUCCESS = 255;  // BIN (BE): 11111111.
    static const char SERVER_REPLY_TIMEOUT = 85;   // BIN (BE): 01010101.
};

#endif // PROTOCOL_H
