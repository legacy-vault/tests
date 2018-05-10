// client.h.

#ifndef CLIENT_H
#define CLIENT_H

#include <QByteArray>
#include <QIODevice>
#include <QObject>
#include <QString>
#include <QTcpSocket>
#include <QTextEdit>
#include <QtGlobal>
#include <QTimer>

#ifndef QT_NO_NETWORKPROXY
#include <QAuthenticator>
#include <QNetworkProxy>
#endif

class Client : public QObject
{
    Q_OBJECT

public:
    // Constructor.
    explicit Client(QObject *parent = Q_NULLPTR);

    // Destructor.
    virtual ~Client();

    // Constants.
    //

    // Methods.
    QString getAddress();
    bool isTryingToConnect();

public slots:
    void abort();
    void configure(
        QString hostName,
        quint16 hostPort,
        QIODevice::OpenMode hostMode,
        QTcpSocket::NetworkLayerProtocol hostProtocol);
    void connect();
    void disconnect();
    void messageSend(QByteArray text);

private slots:
    void socketConnectTimeout();
    void socketError(QTcpSocket::SocketError error);
    void socketHostIsFound();
    void socketIsConnected();
    void socketIsDisconnected();
#ifndef QT_NO_NETWORKPROXY
    void socketProxyAuthenticationIsRequired(
            const QNetworkProxy &proxy,
            QAuthenticator *authenticator);
#endif
    void socketReplyIsAvailable();
    void socketStateIsChanged(QTcpSocket::SocketState state);

signals:
    void clientError(QTcpSocket::SocketError error, QString text);
    void clientHostIsFound();
    void clientIsAborted();
    void clientIsConnected();
    void clientIsDisconnected();
#ifndef QT_NO_NETWORKPROXY
    void clientProxyAuthenticationIsRequired(
            const QNetworkProxy &proxy,
            QAuthenticator *authenticator);
#endif
    void clientStateIsChanged(QTcpSocket::SocketState state);
    void sendDone();
    void sendFailed();
    void serverReplyError();
    void serverReplyTimeout();
    void serverReplyUnexpected();

private:
    // Flags.
    bool connectionIsInProgress; // Connection Attempt is being done.
    bool waitingForReply;

    // Configuration of Client.
    QIODevice::OpenMode hostMode;
    QString hostName;
    quint16 hostPort;
    QTcpSocket::NetworkLayerProtocol hostProtocol;

    // Network TCP Socket.
    QTcpSocket socket;

    // Timers.
    QTimer timerConnect;
    QTimer timerReply;
    int timerConnectTimeoutInterval;
};

#endif //CLIENT_H
