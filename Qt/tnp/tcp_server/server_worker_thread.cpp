// server_worker_thread.cpp.

#include "server_worker_thread.h"

#include <QDateTime>
#include <QObject>
#include <QTcpSocket>

// Constructor.
ServerWorkerThread::ServerWorkerThread(
        QTcpSocket *conn,
        ServerWorker *w,
        QObject *parent)
    : QThread(parent)
{
    connection = conn;
    worker = w;
    timeStart = QDateTime::currentDateTime();
}

// Destructor.
ServerWorkerThread::~ServerWorkerThread()
{
    ;
}

// Returns 'connection'.
QTcpSocket *ServerWorkerThread::getConnection()
{
    return connection;
}

// Returns 'timeStart'.
QDateTime ServerWorkerThread::getTimeStart()
{
    return timeStart;
}

// Returns 'timeStop'.
QDateTime ServerWorkerThread::getTimeStop()
{
    return timeStop;
}

// Returns 'worker'.
ServerWorker *ServerWorkerThread::getWorker()
{
    return worker;
}

// Automatic Start.
// [Overloaded Method of base Class]
void ServerWorkerThread::run()
{
    QTcpSocket *conn;

    conn = connection;

    emit workerStart(conn);
}

// Signal from Worker about finished Job.
void ServerWorkerThread::workerFinished()
{
    timeStop = QDateTime::currentDateTime();
    this->quit();
}
