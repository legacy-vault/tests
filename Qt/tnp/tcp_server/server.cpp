// server.cpp.

#include "server_worker_thread.h"
#include "server_worker.h"
#include "server.h"

#include <iostream>

#include <QAbstractSocket>
#include <QByteArray>
#include <QDateTime>
#include <QObject>
#include <QString>
#include <QTcpServer>
#include <QTcpSocket>
#include <QtGlobal>

// Constructor.
Server::Server(QObject *parent)
    : QObject(parent)
{
    // Initialize Counters.
    threadsListLastCleanTime = QDateTime::currentDateTime();

    // Connect Socket's Signals with Slots.

    // Accept Error.
    QObject::connect(&(this->server), &QTcpServer::acceptError,
                     this, &Server::tcpAcceptError);

    // New Connection.
    QObject::connect(&(this->server), &QTcpServer::newConnection,
                     this, &Server::tcpNewConnection);
}

// Destructor.
Server::~Server()
{
    ;
}

// Returns active Workers Count.
quint64 Server::getWorkersCount()
{
    return workersCount;
}

// Configures the Server.
void Server::configure(quint32 srvIPAddress, quint16 srvPort)
{
    address.setAddress(srvIPAddress);
    port = srvPort;
}

// Starts listening.
void Server::start()
{
    server.listen(address, port);

    // Check that Everything was good.
    if (server.isListening() == false)
    {
        // Inform GUI.
        emit listeningFailed();
    }
    else
    {
        // Inform GUI.
        emit listeningStarted();
    }
}

// Stops listening.
void Server::stop()
{
    server.close();
}

// Accept Error.
void Server::tcpAcceptError(QAbstractSocket::SocketError socketError)
{
    // Inform GUI.
    emit acceptError(socketError);
}

// New incoming Connection. Multi-Threaded Version.
void Server::tcpNewConnection()
{
    QTcpSocket *connection;
    ServerWorkerThread *thread;
    ServerWorker *worker;
    int workerID;

    // Inform GUI about a new Connection.
    emit newConnection();

    // Get a Pointer to TCP Socket.
    if (server.hasPendingConnections() == false)
    {
        return;
    }
    connection = server.nextPendingConnection();

    // Increase Workers Counters.
    workersDataMutex.lock();
    workerIDLast++;
    workerID = workerIDLast;
    workersCount++;
    workersDataMutex.unlock();

    // Create a Worker.
    // We can not set a Parent while it will be in another Thread.
    worker = new ServerWorker(workerID);

    // Create a Thread.
    thread = new ServerWorkerThread(connection, worker, this);

    // Put a Worker to the new Thread.
    worker->moveToThread(thread);

    // Bind Signals to Slots.

    // 1.   When Thread starts, it runs its 'run' Slot.
    //      Our customized 'run' Slot emits the 'workerStart' Signal.
    //      This Signal starts the Worker.
    QObject::connect(thread, &ServerWorkerThread::workerStart,
                     worker, &ServerWorker::start);

    // 2.   Signaling of finished Work is for stopping the Thread.
    QObject::connect(worker, &ServerWorker::finished,
                     this, &Server::workerFinished);

    QObject::connect(worker, &ServerWorker::finished,
                     thread, &ServerWorkerThread::workerFinished);

    // 3.   Worker tells the Server about new incoming Message.
    //      Server will retranslate it to GUI.
    QObject::connect(worker, &ServerWorker::newMessage,
                     this, &Server::workerNewMessage);

    // 4. Worker asks Server to make an Action on Worker's Behalf.
    QObject::connect(worker, &ServerWorker::sendReplyToSocket,
                     this, &Server::workerSendReplyToSocket);

    QObject::connect(worker, &ServerWorker::closeConnection,
                     this, &Server::workerCloseConnection);

    // Start the Thread and let it start the Worker.
    thread->start();

    emit connectionsIndicatorUpdate();

    // Register a new Thread in the List.
    threadRegister(thread);

    // Console Report.
    std::cout << "Busy Workers Count: " << workersCount << ".\r\n";
}

// Close TCP Socket on behalf of the Worker.
// The Worker can not do this while he is in another Thread.
void Server::workerCloseConnection(QTcpSocket *socket)
{
    socket->close();
}

// New Message from a Worker in Thread.
// This Message must be retranslated to the GUI.
void Server::workerNewMessage(QString msg)
{
    emit newMessage(msg);
}

// Worker has finished his Job.
void Server::workerFinished(int id)
{
    // Decrease a Counter.
    workersDataMutex.lock();
    workersCount--;
    workersDataMutex.unlock();

    emit connectionsIndicatorUpdate();

    // Console Report.
    std::cout << "Worker[" << id << "] has finished.\r\n"; //dbg
    std::cout << "Busy Workers Count: " << workersCount << ".\r\n";
}

// Send Reply to TCP Socket on behalf of the Worker.
// The Worker can not do this while he is in another Thread.
void Server::workerSendReplyToSocket(QTcpSocket *socket, QByteArray reply)
{
    qint64 bytesWritten;
    int replySize;

    replySize = reply.size();
    bytesWritten = socket->write(reply);
    if (bytesWritten != replySize)
    {
        // Emergency Exit.
        socket->close();
    }
}

// Clean the Gargabe from the Threads List.
void Server::threadListClean()
{
    // This Function is called only from the 'threadRegister',
    // which already locks the List-Access Mutex,
    // so here we do not touch any Mutexes.

    QTcpSocket *connection;
    int i;
    int listSize;
    ServerWorkerThread *thread;
    bool threadIsActive;
    ServerWorker *worker;

    listSize = threads.size();
    for (i = 0; i < listSize; i++)
    {
        thread = threads[i];
        threadIsActive = thread->isRunning();
        if (!threadIsActive)
        {
            // Delete unused Connection.
            connection = thread->getConnection();
            delete connection;

            // Delete Worker.
            worker = thread->getWorker();
            worker->deleteLater();

            // Delete Thread.
            thread->deleteLater();

            // Remove Thread Pointer from List.
            threads.removeAt(i);
            listSize--;
            i--;
        }
    }

    threadsListLastCleanTime = QDateTime::currentDateTime();
}

// Registeres a THread in the List.
void Server::threadRegister(ServerWorkerThread *thread)
{
    QDateTime listNPCT; // Next Predicted Cleaning Time of the List.
    QDateTime now;

    // Lock.
    threadsListMutex.lock();

    // Add Element to List.
    threads.append(thread);

    // Decide to clean or not to clean the Garbage.
    listNPCT = threadsListLastCleanTime.addSecs(THREADS_LIST_CLEAN_INTERVAL);
    now = QDateTime::currentDateTime();
    if (now >= listNPCT)
    {
        // It is Time to do the Cleaning.
        threadListClean();
    }

    // UnLock.
    threadsListMutex.unlock();
}
