// server_worker_thread.h.

#ifndef SERVER_WORKER_THREAD_H
#define SERVER_WORKER_THREAD_H

#include <QDateTime>
#include <QObject>
#include <QTcpSocket>
#include <QThread>

// Forward Declarations.
class ServerWorker;

class ServerWorkerThread : public QThread
{
    Q_OBJECT

public:
    //Constructor.
    explicit ServerWorkerThread(
            QTcpSocket *connection,
            ServerWorker *worker,
            QObject *parent = Q_NULLPTR);

    // Destructor.
    ~ServerWorkerThread();

    // Methods.
    QTcpSocket *getConnection();
    QDateTime getTimeStart();
    QDateTime getTimeStop();
    ServerWorker *getWorker();

    // Fields.
    //

public slots:
    void run();
    void workerFinished();

signals:
    void workerStart(QTcpSocket *connection);

private:
    // Fields.
    QTcpSocket *connection;
    QDateTime timeStart;
    QDateTime timeStop;
    ServerWorker *worker;

    // Methods.
    //
};

#endif // SERVER_WORKER_THREAD_H
