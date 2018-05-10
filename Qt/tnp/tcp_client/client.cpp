// client.cpp.

#include "protocol.h"
#include "client.h"

#include <QByteArray>
#include <QIODevice>
#include <QObject>
#include <QString>
#include <QTcpSocket>
#include <QtGlobal>
#include <QTimer>

// Constructor.
Client::Client(QObject *parent)
    : QObject(parent)
{
    hostName = QString("");
    hostPort = 0;
    hostMode = QIODevice::ReadWrite;
    hostProtocol = QTcpSocket::AnyIPProtocol;

    // Timers' Intervals.
    timerConnectTimeoutInterval = Protocol::CLIENT_CONNECTION_TIMEOUT * 1000;

    connectionIsInProgress = false;

    // Connect Socket's Signals with Slots.

    QObject::connect(&timerConnect, &QTimer::timeout,
                     this, &Client::socketConnectTimeout);

    QObject::connect(
        &socket,
        // We can not use simple '&QTcpSocket::error' here.
        QOverload<QTcpSocket::SocketError>::of(&QTcpSocket::error),
        this,
        &Client::socketError);

    QObject::connect(&socket, &QTcpSocket::hostFound,
                     this, &Client::socketHostIsFound);

    QObject::connect(&socket, &QTcpSocket::connected,
                     this, &Client::socketIsConnected);

    QObject::connect(&socket, &QTcpSocket::disconnected,
                     this, &Client::socketIsDisconnected);

#ifndef QT_NO_NETWORKPROXY
    QObject::connect(&socket, &QTcpSocket::proxyAuthenticationRequired,
                     this, &Client::socketProxyAuthenticationIsRequired);
#endif

    QObject::connect(&socket, &QTcpSocket::readyRead,
                     this, &Client::socketReplyIsAvailable);

    QObject::connect(&socket, &QTcpSocket::stateChanged,
                     this, &Client::socketStateIsChanged);
}

// Destructor.
Client::~Client()
{
    ;
}

// Gets Client's Address String.
QString Client::getAddress()
{
    QString address;

    address = hostName + ":" + QString::number(hostPort);

    return address;
}

// Tells whether the Client is now trying to connect or not.
bool Client::isTryingToConnect()
{
    return connectionIsInProgress;
}

// Aborts the Connection Attempt.
void Client::abort()
{
    // Stop the Connection Attempt.
    socket.abort();

    // Stop the Timer.
    timerConnect.stop();

    // Connection has been aborted and is now stopped.
    connectionIsInProgress = false;

    // Signal to others.
    emit clientIsAborted();
}

// Configures the Client.
void Client::configure(
                QString host,
                quint16 port,
                QIODevice::OpenMode mode,
                QTcpSocket::NetworkLayerProtocol protocol)
{
    hostName = host;
    hostPort = port;
    hostMode = mode;
    hostProtocol = protocol;
}

// Starts the Connection Attempt.
void Client::connect()
{
    // Start Timeout Timer.
    timerConnect.setInterval(timerConnectTimeoutInterval);
    timerConnect.start();

    // Try to connect.
    connectionIsInProgress = true;
    socket.connectToHost(
                hostName,
                hostPort,
                hostMode,
                hostProtocol);

    // Signaling is done when the Socket is connected.
}

// Closes the Connection.
void Client::disconnect()
{
    // Close Connection.
    socket.close();

    // Signaling is done when the Socket is closed.
}

// Sends a Message to the Host.
void Client::messageSend(QByteArray text)
{
    qint64 bytesWritten;
    int i;
    int iMax;
    int j;
    QByteArray msg;
    int msgLength;
    int textLength;
    quint64 textLength64;

    textLength = text.size();
    textLength64 = quint64(textLength);

    msgLength = textLength + Protocol::MESSAGE_FIELD_SIZE_LEN;
    msg.resize(msgLength);

    // Write 'Size' Field. Format is "Big Endian".
    j = Protocol::MESSAGE_FIELD_SIZE_LEN - 1; // Reverse Counter;
    for (i = 0; i < Protocol::MESSAGE_FIELD_SIZE_LEN; i++)
    {
        msg[i] = quint8(textLength64 >> (j*8));
        j--;
    }

    // Write 'Message' Field.
    i = Protocol::MESSAGE_FIELD_SIZE_LEN;
    iMax = Protocol::MESSAGE_FIELD_SIZE_LEN + textLength - 1;
    j = 0;
    while (i <= iMax)
    {
        msg[i] = text[j];
        i++;
        j++;
    }

    // Write to Socket.
    bytesWritten = this->socket.write(msg);
    if (bytesWritten != msgLength)
    {
        emit sendFailed();
        this->disconnect();
        return;
    }

    // Now we are waiting for the 'readyRead' Signal from Socket.
    waitingForReply = true;
}

// Connection Attempt has taken too much Time.
void Client::socketConnectTimeout()
{
    // Stop the Connection Attempt.
    socket.abort();

    // Stop the Timer.
    timerConnect.stop();

    // Connection has failed and is now stopped.
    connectionIsInProgress = false;

    // Signal to others.
    emit clientIsAborted();
}

// Connectivity Error.
void Client::socketError(QTcpSocket::SocketError error)
{
    QString socketErrorText;
    bool timerIsActive;

    // Abort Connection.
    socket.abort();

    // Stop the Timer, if it is running.
    timerIsActive = timerConnect.isActive();
    if (timerIsActive)
    {
        timerConnect.stop();
    }

    // If an Error occurred during Connection Attempt,
    // we must disable its Flag.
    if (connectionIsInProgress == true)
    {
        connectionIsInProgress = false;
    }

    // Signal to others.
    socketErrorText = socket.errorString();
    emit clientError(error, socketErrorText);
}

// Host is Found.
void Client::socketHostIsFound()
{
    emit clientHostIsFound();
}

// Connection has been established.
void Client::socketIsConnected()
{
    // Disable Conection Timer.
    timerConnect.stop();

    // Connection is done.
    connectionIsInProgress = false;

    // Signal to others.
    emit clientIsConnected();
}

// Connection has been closed.
void Client::socketIsDisconnected()
{
    emit clientIsDisconnected();
}

// Connections Requires Proxy Authentication.
#ifndef QT_NO_NETWORKPROXY
void Client::socketProxyAuthenticationIsRequired(
        const QNetworkProxy &proxy,
        QAuthenticator *authenticator)
{
    emit clientProxyAuthenticationIsRequired(proxy, authenticator);
}
#endif

// Reply from Host is available for Reading.
void Client::socketReplyIsAvailable()
{
    quint8 byte;
    QByteArray reply;
    quint8 replyError;
    quint8 replyGood;
    quint8 replyTimeout;

    // Prepare Data.
    replyGood = quint8(Protocol::SERVER_REPLY_SUCCESS);
    replyError = quint8(Protocol::SERVER_REPLY_ERROR);
    replyTimeout = quint8(Protocol::SERVER_REPLY_TIMEOUT);

    // Read Host's Reply. We need only 1 Byte.
    reply = socket.read(1);
    if (reply.size() != 1)
    {
        if (waitingForReply == true)
        {
            // We were waiting for Reply, but we failed to read it.
            // Disable waiting and mark the Message as unsent.
            waitingForReply = false;
            emit sendFailed();
            socket.close();
            return;
        }
        else
        {
            // We were not waiting for any Reply.
            // So, a Message from the Server may be:
            //      * about the Timeout,
            //      * about Server's Error,
            //      * or a Message of a Type we are not aware of.
            // We can ignore all of them, but we must close the Connection
            // to show that something strange is happening.
            socket.close();
            return;
        }
    }

    // Reply is received.
    byte = reply[0];

    if (waitingForReply != true)
    {
        // We were not waiting for any Reply.
        // So, a Message from the Server may be:
        //      * about the Timeout,
        //      * about Server's Error,
        //      * or a Message of a Type we are not aware of.

        if (byte == replyTimeout)
        {
            // We have been idle for too long.
            emit serverReplyTimeout();
            socket.close();
            return;
        }
        else if (byte == replyError)
        {
            // Server's internal Error.
            emit serverReplyError();
            socket.close();
            return;
        }
        else
        {
            // Message of an unknown Type.
            // We must close the Connection to show that something strange
            // is happening.
            emit serverReplyUnexpected();
            socket.close();
            return;
        }
    }
    else
    {
        // We are waiting for Reply, and it is now received.
        // It can be:
        //      * about the successfull Message Delivery,
        //      * about Server's Error,
        //      * or a Message of a Type we are not aware of.

        if (byte == replyGood)
        {
            // Successfull Message Delivery.
            waitingForReply = false;
            emit sendDone();
            return;
        }
        else if (byte == replyError)
        {
            // Server's internal Error.
            waitingForReply = false;
            emit sendFailed();
            socket.close();
            return;
        }
        else
        {
            // Message of an unknown Type.
            // We must close the Connection to show that something strange
            // is happening.
            waitingForReply = false;
            emit serverReplyUnexpected();
            socket.close();
            return;
        }
    }
}

// Connection State has changed.
void Client::socketStateIsChanged(QTcpSocket::SocketState state)
{
    emit clientStateIsChanged(state);
}
