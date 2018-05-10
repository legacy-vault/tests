// server.h.

#ifndef SERVER_H
#define SERVER_H

#include <QAbstractSocket>
#include <QByteArray>
#include <QDateTime>
#include <QHostAddress>
#include <QMutex>
#include <QObject>
#include <QString>
#include <QTcpServer>
#include <QTcpSocket>
#include <QtGlobal>

// Forward Declarations.
class ServerWorkerThread;
class ServerWorker;

class Server : public QObject
{
    Q_OBJECT

public:
    // Constructor.
    explicit Server(QObject *parent = Q_NULLPTR);

    // Destructor.
    virtual ~Server();

    // Constants.
    static const int WORKER_TIMER_TICK = 500; // Milliseconds.
    static const int THREADS_LIST_CLEAN_INTERVAL= 60; // Seconds
    // N.B.: Lower Values decrease the 'Ping', but increase the CPU Load.

    // Methods.
    quint64 getWorkersCount();

public slots:
    void configure(quint32 srvIPAddress, quint16 srvPort);
    void start();
    void stop();

private slots:
    void tcpAcceptError(QAbstractSocket::SocketError socketError);
    void tcpNewConnection();
    void workerCloseConnection(QTcpSocket *socket);
    void workerNewMessage(QString msg);
    void workerFinished(int id);
    void workerSendReplyToSocket(QTcpSocket *socket, QByteArray reply);

signals:
    void acceptError(QAbstractSocket::SocketError socketError);
    void connectionsIndicatorUpdate();
    void listeningFailed();
    void listeningStarted();
    void newConnection();
    void newMessage(QString msg); // to GUI.

private:
    // Network TCP Server.
    QTcpServer server;

    // Configuration.
    QHostAddress address;
    quint16 port;

    // Workers Data.
    quint64 workerIDLast; // ID of Worker who last received the Job.
    quint64 workersCount; // Number of Workers who are busy.
    QMutex workersDataMutex; // Protects: 'workerIDLast', 'workersCount'.

    // Threads List.
    QList<ServerWorkerThread *> threads;
    QMutex threadsListMutex; // Protects: 'threads', 'threadsListLastCleanTime'.
    QDateTime threadsListLastCleanTime;

    // Methods.
    void threadListClean();
    void threadRegister(ServerWorkerThread *thread);
};

#endif // SERVER_H
