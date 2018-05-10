// server_worker.h.

#ifndef SERVER_WORKER_H
#define SERVER_WORKER_H

#include <QByteArray>
#include <QObject>
#include <QString>
#include <QTcpSocket>
#include <QtGlobal>

class ServerWorker : public QObject
{
    Q_OBJECT

public:
    // Constructor.
    explicit ServerWorker(int id, QObject *parent = Q_NULLPTR);

    // Destructor.
    virtual ~ServerWorker();

    // Constants.
    static const char STATE_CLIENT_DATA_AVAILABLE = 1;  // 001.
    static const char STATE_CLIENT_CONNECTION_LOST = 2; // 010.
    static const char STATE_CLIENT_IDLE_TIMEOUT = 4;    // 100.

    // Methods.
    int getId();

    // Fields.
    //

public slots:
    void start(QTcpSocket *connection);

signals:
    void closeConnection(QTcpSocket *socket);
    void finished(int id);
    void newMessage(QString msg);
    void sendReplyToSocket(QTcpSocket *socket, QByteArray reply);

private:
    // Fields.
    int id;

    // Methods.
    char waitForMessageData(QTcpSocket *conn, qint64 msgSize);
    char waitForNData(QTcpSocket *connection, qint64 n);
    char waitForSizeData(QTcpSocket *connection);
};

#endif // SERVER_WORKER_H
