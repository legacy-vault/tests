// server_worker.cpp.

#include "server.h"
#include "protocol.h"
#include "functions.h"
#include "server_worker.h"
#include "server_worker_thread.h"

#include <QAbstractSocket>
#include <QByteArray>
#include <QDateTime>
#include <QEventLoop>
#include <QObject>
#include <QString>
#include <QTcpSocket>
#include <QtGlobal>
#include <QTimer>

// Constructor.
ServerWorker::ServerWorker(int i, QObject *parent)
    : QObject(parent)
{
    id = i;
}

// Destructor.
ServerWorker::~ServerWorker()
{
    ;
}

// Returns 'id'.
int ServerWorker::getId()
{
    return id;
}

// Worker's Job.
void ServerWorker::start(QTcpSocket *connection)
{
    QByteArray data1; // Field: Size.
    int data1Size;
    int data1SizeExpected;
    QByteArray data2; // Field: Message.
    int data2Size;
    int data2SizeExpected;
    int i;
    int j;
    QString msgQ;
    quint64 msgSize;
    QByteArray replyError;
    QByteArray replyGood;
    QByteArray replyTimeout;
    char result;

    // Prepare Replies.
    replyError = QByteArray(1, Protocol::SERVER_REPLY_ERROR);
    replyGood = QByteArray(1, Protocol::SERVER_REPLY_SUCCESS);
    replyTimeout = QByteArray(1, Protocol::SERVER_REPLY_TIMEOUT);

    // Serve forever, util the Client becomes idle.
    while (true)
    {
        // Wait for Size-Data Arrival.
        result = waitForSizeData(connection);

        if (result == STATE_CLIENT_CONNECTION_LOST)
        {
            // Connection was lost or closed by the Client.
            emit closeConnection(connection);
            emit finished(this->id);
            return;
        }
        else if (result == STATE_CLIENT_IDLE_TIMEOUT)
        {
            // Connection was idle for too long.
            emit sendReplyToSocket(connection, replyTimeout);
            emit closeConnection(connection);
            emit finished(this->id);
            return;
        }
        else if (result != STATE_CLIENT_DATA_AVAILABLE)
        {
            // Something is wrong.
            emit sendReplyToSocket(connection, replyError);
            emit closeConnection(connection);
            emit finished(this->id);
            return;
        }

        // Read Message Size. Format is "Big Endian".
        data1SizeExpected = Protocol::MESSAGE_FIELD_SIZE_LEN;
        data1 = connection->read(data1SizeExpected);
        data1Size = data1.size();

        // Check.
        if (data1Size < data1SizeExpected)
        {
            // Read Error occurred.
            // Inform the Client about the Error and disconnect.
            emit sendReplyToSocket(connection, replyError);
            emit closeConnection(connection);
            emit finished(this->id);
            return;
        }

        // Big Endian 'data1' -> quint64.
        msgSize = 0;
        j = Protocol::MESSAGE_FIELD_SIZE_LEN - 1; // Reverse Counter.
        for (i = 0; i < Protocol::MESSAGE_FIELD_SIZE_LEN; i++)
        {
            msgSize += (quint64(data1[i]) << (j*8));
            j--;
        }

        // Wait for Message-Data Arrival.
        result = waitForMessageData(connection, qint64(msgSize));

        if (result == STATE_CLIENT_CONNECTION_LOST)
        {
            // Connection was lost or closed by the Client.
            emit closeConnection(connection);
            emit finished(this->id);
            return;
        }
        else if (result == STATE_CLIENT_IDLE_TIMEOUT)
        {
            // Connection was idle for too long.
            emit sendReplyToSocket(connection, replyTimeout);
            emit closeConnection(connection);
            emit finished(this->id);
            return;
        }
        else if (result != STATE_CLIENT_DATA_AVAILABLE)
        {
            // Something is wrong.
            emit sendReplyToSocket(connection, replyError);
            emit closeConnection(connection);
            emit finished(this->id);
            return;
        }

        // Available Data is now enough to get Message from it.
        // Read Message.
        data2SizeExpected = int(msgSize);
        data2 = connection->read(data2SizeExpected);
        data2Size = data2.size();

        // Check.
        if (data2Size < data2SizeExpected)
        {
            // Read Error occurred.
            // Inform the Client about the Error and disconnect.
            emit sendReplyToSocket(connection, replyError);
            emit closeConnection(connection);
            emit finished(this->id);
            return;
        }

        // Convert Message to QString.
        msgQ = QString::fromUtf8(data2);

        // Send good Reply to Client.
        emit sendReplyToSocket(connection, replyGood);

        // Signal to GUI.
        emit newMessage(msgQ);

    } // End of Serve Loop.
}

// Waits for incoming Data to become enough to read Message.
// If waiting too long, boolean 'false' is returned.
char ServerWorker::waitForMessageData(QTcpSocket *conn, qint64 msgSize)
{
    char result;

    result = waitForNData(conn, msgSize);

    return result;
}

// Waits for incoming Data to become enough to read N bytes.
// Reaturns the State of the Process.
char ServerWorker::waitForNData(QTcpSocket *connection, qint64 n)
{
    qint64 bytesAvailable;
    bool bytesAreNotEnough;
    qint64 bytesExpected;
    bool connectionIsGood;
    QAbstractSocket::SocketState connectionState;
    bool connectionStateIsGood;
    quint64 loopNum;
    quint64 loopNumMax;
    int timeoutTime; // ms.
    int timerTick; // ms.

    // Prepare Data
    timeoutTime = Protocol::MESSAGE_IDLE_TIMEOUT * 1000;
    timerTick = Server::WORKER_TIMER_TICK;
    loopNumMax = quint64(timeoutTime / timerTick);
    loopNum = 0;
    bytesExpected = n;

    // Check State of Connection.
    bytesAvailable = connection->bytesAvailable();
    bytesAreNotEnough = (bytesAvailable < bytesExpected);
    connectionState = connection->state();
    connectionStateIsGood =
            (connectionState != QAbstractSocket::UnconnectedState) &&
            (connectionState != QAbstractSocket::ClosingState);
    connectionIsGood =
            connectionStateIsGood &&
            (connection->isOpen()) &&
            (connection->isReadable()) &&
            (connection->isValid()) &&
            (connection->isWritable());

    if (!connectionIsGood)
    {
        return STATE_CLIENT_CONNECTION_LOST;
    }

    // Loop.
    while (bytesAreNotEnough)
    {
        // Sleep.
        sleep(timerTick);

        // Check State of Connection.
        bytesAvailable = connection->bytesAvailable();
        bytesAreNotEnough = (bytesAvailable < bytesExpected);
        connectionState = connection->state();
        connectionStateIsGood =
                (connectionState != QAbstractSocket::UnconnectedState) &&
                (connectionState != QAbstractSocket::ClosingState);
        connectionIsGood =
                connectionStateIsGood &&
                (connection->isOpen()) &&
                (connection->isReadable()) &&
                (connection->isValid()) &&
                (connection->isWritable());

        if (!connectionIsGood)
        {
            return STATE_CLIENT_CONNECTION_LOST;
        }

        loopNum++;
        if (loopNum >= loopNumMax)
        {
            return STATE_CLIENT_IDLE_TIMEOUT;
        }
    }

    return STATE_CLIENT_DATA_AVAILABLE;
}

// Waits for incoming Data to become enough to read Message Size.
// If waiting too long, boolean 'false' is returned.
char ServerWorker::waitForSizeData(QTcpSocket *connection)
{
    char result;

    result = waitForNData(connection, Protocol::MESSAGE_FIELD_SIZE_LEN);

    return result;
}
