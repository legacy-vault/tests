// mainwindow.h.

#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QAbstractSocket>
#include <QByteArray>
#include <QIODevice>
#include <QKeyEvent>
#include <QLineEdit>
#include <QMainWindow>
#include <QPushButton>
#include <QSpinBox>
#include <QString>
#include <QTcpSocket>
#include <QTextEdit>
#include <QtGlobal>

#ifndef QT_NO_NETWORKPROXY
#include <QAuthenticator>
#include <QNetworkProxy>
#endif

namespace Ui
{
    class MainWindow;
}

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    // Constructor.
    explicit MainWindow(QWidget *parent = Q_NULLPTR);

    // Destructor.
    ~MainWindow();

    // Methods.
    void journalRecordAdd(QString &text);

public slots:
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

private slots:
    void buttonConnectClicked();
    void buttonDisconnectClicked();
    void buttonSendClicked();

signals:
    void clientAbort();
    void clientConfigure(
            QString hostName,
            quint16 hostPort,
            QIODevice::OpenMode hostMode,
            QAbstractSocket::NetworkLayerProtocol hostProtocol);
    void clientConnect();
    void clientDisconnect();
    void messageIsPrepared(QByteArray text);

private:
    // Journal.
    QTextEdit *teJournal;

    // Buttons.
    QPushButton *pbConnect;
    QPushButton *pbDisconnect;
    QPushButton *pbSend;

    // Inputs.
    QLineEdit *leHost;
    QLineEdit *leMessage;
    QSpinBox *sbPort;

    // UI.
    Ui::MainWindow *ui;

    void keyPressEvent(QKeyEvent *event);
};

#endif // MAINWINDOW_H
