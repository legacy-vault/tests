// mainwindow.cpp.

#include "client.h"
#include "mainwindow.h"
#include "protocol.h"
#include "ui_mainwindow.h"

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
#include <QWidget>

#ifndef QT_NO_NETWORKPROXY
#include <QAuthenticator>
#include <QNetworkProxy>
#endif

// Constructor.
MainWindow::MainWindow(QWidget *parent) :
    QMainWindow(parent),
    ui(new Ui::MainWindow)
{
    // Names of GUI Elements.
    QString leHostName;
    QString leMessageName;
    QString pbConnectName;
    QString pbDisconnectName;
    QString pbSendName;
    QString sbPortName;
    QString teJournalName;

    // Set up GUI.
    ui->setupUi(this);

    // Find Objects.

    // 1. Buttons.
    pbConnectName = QString("pbConnect");
    pbConnect = findChild<QPushButton *>(pbConnectName);
    pbDisconnectName = QString("pbDisconnect");
    pbDisconnect = findChild<QPushButton *>(pbDisconnectName);
    pbSendName = QString("pbSend");
    pbSend = findChild<QPushButton *>(pbSendName);

    // 2. Inputs.
    leHostName = QString("leHost");
    leHost = findChild<QLineEdit *>(leHostName);
    leMessageName = QString("leMessage");
    leMessage = findChild<QLineEdit *>(leMessageName);
    sbPortName = QString("sbPort");
    sbPort = findChild<QSpinBox *>(sbPortName);

    // 3. Journal.
    teJournalName = QString("teJournal");
    teJournal = findChild<QTextEdit *>(teJournalName);

    // Connect Signals with Slots.
    QObject::connect(pbConnect, &QPushButton::clicked,
                     this, &MainWindow::buttonConnectClicked);

    QObject::connect(pbDisconnect, &QPushButton::clicked,
                     this, &MainWindow::buttonDisconnectClicked);

    QObject::connect(pbSend, &QPushButton::clicked,
                     this, &MainWindow::buttonSendClicked);
}

// Destructor.
MainWindow::~MainWindow()
{
    delete ui;
}

// 'Connect' Button is clicked.
void MainWindow::buttonConnectClicked()
{
    // Client.
    extern Client client;
    QString clientAddress;
    QString host;
    QIODevice::OpenMode mode;
    quint16 port;
    QAbstractSocket::NetworkLayerProtocol protocol;

    // Various temporary Data.
    QString journalRecord;
    QString pbDisconnectText;

    // Change Availability of Objects.
    // 1. Enable.
    pbDisconnect->setDisabled(false);
    // 2. Disable.
    leHost->setDisabled(true);
    pbConnect->setDisabled(true);
    sbPort->setDisabled(true);

    // Prepare Data.
    host = leHost->text();
    port = quint16(sbPort->value());
    mode = QIODevice::ReadWrite;
    protocol = QAbstractSocket::AnyIPProtocol;

    // Configure Client.
    emit clientConfigure(host, port, mode, protocol);

    // Add a Record to Journal.
    clientAddress = client.getAddress();
    journalRecord = QString("Connecting to " +
                            clientAddress + "...");
    journalRecordAdd(journalRecord);

    // Prepare to abort the Connection.
    pbDisconnectText = QString("Abort");
    pbDisconnect->setText(pbDisconnectText);

    // Connect Client.
    emit clientConnect();
}

// 'Disconnect' Button is clicked.
void MainWindow::buttonDisconnectClicked()
{
    // Client.
    extern Client client;
    bool connectionIsInProgress;

    connectionIsInProgress = client.isTryingToConnect();

    // Decide what to do.
    if (connectionIsInProgress)
    {
        // Abort the Connection.
        emit clientAbort();
    }
    else
    {
        // Close the Connection.
        emit clientDisconnect();
    }

    // Change Availability of Objects.
    pbDisconnect->setDisabled(true);
}

// 'Send' Button is clicked.
void MainWindow::buttonSendClicked()
{
    // Message.
    QString messageUTF16;
    QByteArray messageUTF8;
    int messageUTF8Size; // Size in Bytes.

    QString journalRecord;

    // Get Message.
    messageUTF16 = leMessage->text();
    messageUTF8 = messageUTF16.toUtf8();
    messageUTF8Size = messageUTF8.size();

    // Check Message Length.
    if (messageUTF8Size > Protocol::MESSAGE_TEXT_BYTES_COUNT_MAX)
    {
        // Add a Record to Journal.
        journalRecord = QString("The Message is too long!");
        journalRecordAdd(journalRecord);
        return;
    }

    // Disable Button.
    pbSend->setDisabled(true);

    emit messageIsPrepared(messageUTF8);
}

// Connectivity Error.
void MainWindow::clientError(QTcpSocket::SocketError error, QString text)
{
    QString journalRecord;

    // Add a Record to Journal.
    journalRecord =
            QString("Connectivity Error [") +
            QString::number(error)  +
            QString("] : ") +
            text +
            QString(".");
    journalRecordAdd(journalRecord);

    // Change Availability of Objects.
    // 1. Enable.
    pbConnect->setDisabled(false);
    leHost->setDisabled(false);
    sbPort->setDisabled(false);
    // 2. Disable.
    pbSend->setDisabled(true);
    pbDisconnect->setDisabled(true);
}

// Client Host is Found.
void MainWindow::clientHostIsFound()
{
    QString journalRecord;

    // Add a Record to Journal.
    journalRecord = QString("Host is found.");
    journalRecordAdd(journalRecord);
}

// Connection Attempt has been aborted or cancelled.
void MainWindow::clientIsAborted()
{
    // Client.
    extern Client client;
    QString clientAddress;

    QString journalRecord;

    // Add a Record to Journal.
    clientAddress = client.getAddress();
    journalRecord = QString("Connection to " +
                            clientAddress + " has been aborted.");
    journalRecordAdd(journalRecord);

    // Change Availability of Objects.
    // 1. Enable.
    leHost->setDisabled(false);
    pbConnect->setDisabled(false);
    sbPort->setDisabled(false);
    // 2. Disable.
    pbDisconnect->setDisabled(true);
}

// Connection has been established.
void MainWindow::clientIsConnected()
{
    // Client.
    extern Client client;
    QString clientAddress;

    QString btnDisconnectText;
    QString journalRecord;

    // Add a Record to Journal.
    clientAddress = client.getAddress();
    journalRecord = QString("Connection to " +
                            clientAddress + " has been established.");
    journalRecordAdd(journalRecord);

    // Forget about Connection's Abort.
    btnDisconnectText = QString("Disconnect");
    pbDisconnect->setText(btnDisconnectText);

    // Change Availability of Objects.
    pbSend->setDisabled(false);
}

// Connection has been closed.
void MainWindow::clientIsDisconnected()
{
    // Client.
    extern Client client;
    QString clientAddress;

    QString journalRecord;

    // Add a Record to Journal.
    clientAddress = client.getAddress();
    journalRecord = QString("Connection to " +
                            clientAddress + " has been closed.");
    journalRecordAdd(journalRecord);

    // Change Availability of Objects.
    // 1. Enable.
    pbConnect->setDisabled(false);
    leHost->setDisabled(false);
    sbPort->setDisabled(false);
    // 2. Disable.
    pbSend->setDisabled(true);
    pbDisconnect->setDisabled(true);
}

#ifndef QT_NO_NETWORKPROXY
void MainWindow::clientProxyAuthenticationIsRequired(
            const QNetworkProxy &proxy,
            QAuthenticator *authenticator)
{
    QString journalRecord;

    journalRecord =
        QString("Connection requires Proxy Authentication. Host: ") +
        proxy.hostName() +
        QString(", Realm: ") +
        authenticator->realm() +
        QString(".");

    journalRecordAdd(journalRecord);
}
#endif

// Connection State has been changed.
void MainWindow::clientStateIsChanged(QTcpSocket::SocketState state)
{
    QString stateDescription;

    if (state == 0)
    {
        stateDescription = QString("Not connected.");
    }
    else if (state == 1)
    {
        stateDescription = QString("Performing a Host Name Lookup...");
    }
    else if (state == 2)
    {
        stateDescription = QString("Establishing a Connection...");
    }
    else if (state == 3)
    {
        stateDescription = QString("Connection is established.");
    }
    else if (state == 4)
    {
        stateDescription =
            QString("The Socket is bound to an Address and Port.");
    }
    else if (state == 5)
    {
        stateDescription = QString("For internal Use only.");
    }
    else if (state == 6)
    {
        stateDescription = QString("The Socket is about to close.");
    }
    else
    {
        stateDescription = QString("Unrecognized State.");
    }

    teJournal->append(stateDescription);
}

// Message Send Process has succeeded.
void MainWindow::sendDone()
{
    QString text;

    text = QString("Message has been successfully sent.");

    teJournal->append(text);

    // Enable Button.
    pbSend->setDisabled(false);
}

// Message Send Process has failed.
void MainWindow::sendFailed()
{
    QString text;

    text = QString("Message Send Process has failed.");

    teJournal->append(text);
}

// Server tells us about an its internal Error.
void MainWindow::serverReplyError()
{
    QString text;

    text = QString("Internal Server Error.");

    teJournal->append(text);
}

// Server tells us about our Timeout.
void MainWindow::serverReplyTimeout()
{
    QString text;

    text = QString("Idle Connection Timeout.");

    teJournal->append(text);
}


// Server has sent an unexpected Reply.
void MainWindow::serverReplyUnexpected()
{
    QString text;

    text = QString("UnExpected Reply from Server.");

    teJournal->append(text);
}

// Pressed Key Handler.
void MainWindow::keyPressEvent(QKeyEvent *event)
{
    int keyCode;
    QString text;

    keyCode = event->key();

    if (keyCode == Qt::Key_Escape)
    {
        this->close();
    }
}

// Adds a Text Record to the Journal.
void MainWindow::journalRecordAdd(QString &text)
{
    teJournal->append(text);
}
