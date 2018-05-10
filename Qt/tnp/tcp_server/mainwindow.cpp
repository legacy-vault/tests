// mainwindow.cpp.

#include "server.h"
#include "mainwindow.h"
#include "ui_mainwindow.h"

#include <QAbstractSocket>
#include <QKeyEvent>
#include <QLCDNumber>
#include <QMainWindow>
#include <QPushButton>
#include <QSpinBox>
#include <QString>
#include <QTextEdit>
#include <QtGlobal>

// Constructor.
MainWindow::MainWindow(QWidget *parent) :
    QMainWindow(parent),
    ui(new Ui::MainWindow)
{
    // Names of GUI Elements.
    QString lcdConnectionsName;
    QString pbStartName;
    QString pbStopName;
    QString sbIPA1Name;
    QString sbIPA2Name;
    QString sbIPA3Name;
    QString sbIPA4Name;
    QString sbPortName;
    QString teJournalName;

    // Set up GUI.
    ui->setupUi(this);

    // Find Objects.

    // 1. Journal.
    teJournalName = QString("teJournal");
    teJournal = findChild<QTextEdit *>(teJournalName);

    // 2. Buttons.
    pbStartName = QString("pbStart");
    pbStart = findChild<QPushButton *>(pbStartName);
    pbStopName = QString("pbStop");
    pbStop = findChild<QPushButton *>(pbStopName);

    // 3. Inputs.
    sbIPA1Name = QString("sbIPA_1");
    sbIPA1 = findChild<QSpinBox *>(sbIPA1Name);
    sbIPA2Name = QString("sbIPA_2");
    sbIPA2 = findChild<QSpinBox *>(sbIPA2Name);
    sbIPA3Name = QString("sbIPA_3");
    sbIPA3 = findChild<QSpinBox *>(sbIPA3Name);
    sbIPA4Name = QString("sbIPA_4");
    sbIPA4 = findChild<QSpinBox *>(sbIPA4Name);
    sbPortName = QString("sbPort");
    sbPort = findChild<QSpinBox *>(sbPortName);

    // 4. Indicators.
    lcdConnectionsName = QString("lcdConnections");
    lcdConnections = findChild<QLCDNumber *>(lcdConnectionsName);

    // Connect Signals with Slots.
    QObject::connect(pbStart, &QPushButton::clicked,
                     this, &MainWindow::buttonStartClicked);

    QObject::connect(pbStop, &QPushButton::clicked,
                     this, &MainWindow::buttonStopClicked);
}

// Destructor.
MainWindow::~MainWindow()
{
    delete ui;
}

// Adds a Text Record to the Journal.
void MainWindow::journalRecordAdd(QString &text)
{
    teJournal->append(text);
}

// Server's Accept Error.
void MainWindow::serverAcceptError(QAbstractSocket::SocketError socketError)
{
    QString journalRecord;

    // Add a Record to Journal.
    journalRecord =
            QString("Server has got an Accept Error [") +
            QString::number(socketError) +
            QString("].");

    journalRecordAdd(journalRecord);
}

// Updates Connections Number Indicator.
void MainWindow::serverConnectionsIndicatorUpdate()
{
    extern Server server;

    quint64 workersCount;
    int workersCountInt;

    workersCount = server.getWorkersCount();
    workersCountInt = int(workersCount);
    lcdConnections->display(workersCountInt);
}

// Server has failed to start.
void MainWindow::serverListeningFailed()
{
    QString journalRecord;

    // Add a Record to Journal.
    journalRecord =
            QString("Server has failed to start. Do check the Settings!");
    journalRecordAdd(journalRecord);

    // Change Availability of Objects.
    // 1. Enable.
    pbStart->setDisabled(false);
    sbIPA1->setDisabled(false);
    sbIPA2->setDisabled(false);
    sbIPA3->setDisabled(false);
    sbIPA4->setDisabled(false);
    sbPort->setDisabled(false);
    // 2. Disable.
    pbStop->setDisabled(true);
}

// Server has started successfully.
void MainWindow::serverListeningStarted()
{
    QString dot;
    int ipA1;
    int ipA2;
    int ipA3;
    int ipA4;
    QString journalRecord;
    quint16 port;
    QString serverAddress;

    // Prepare Data.
    dot = QString(".");
    port = quint16(sbPort->value());

    ipA1 = sbIPA1->value();
    ipA2 = sbIPA2->value();
    ipA3 = sbIPA3->value();
    ipA4 = sbIPA4->value();

    serverAddress =
            QString::number(ipA1) + dot +
            QString::number(ipA2) + dot +
            QString::number(ipA3) + dot +
            QString::number(ipA4) +
            QString(":") +
            QString::number(port);

    // Add a Record to Journal.
    journalRecord = QString("Listening at " + serverAddress + "...");
    journalRecordAdd(journalRecord);
}

// New incoming Connection to Server.
void MainWindow::serverNewConnection()
{
    QString journalRecord;

    // Add a Record to Journal.
    journalRecord = QString("Server has got an incoming Connection.");
    journalRecordAdd(journalRecord);
}

// New incoming Message.
void MainWindow::serverNewMessage(QString msg)
{
    QString journalRecord;

    // Add a Record to Journal.
    journalRecord = QString("Incoming Message:\r\n") + msg;
    journalRecordAdd(journalRecord);
}

// 'Start' Button is clicked.
void MainWindow::buttonStartClicked()
{
    // Server.
    quint32 hostIPAddress;
    quint16 port;
    int ipA1;
    int ipA2;
    int ipA3;
    int ipA4;

    // Change Availability of Objects.
    // 1. Enable.
    pbStop->setDisabled(false);
    // 2. Disable.
    pbStart->setDisabled(true);
    sbIPA1->setDisabled(true);
    sbIPA2->setDisabled(true);
    sbIPA3->setDisabled(true);
    sbIPA4->setDisabled(true);
    sbPort->setDisabled(true);

    // Prepare Data.
    port = quint16(sbPort->value());

    ipA1 = sbIPA1->value();
    ipA2 = sbIPA2->value();
    ipA3 = sbIPA3->value();
    ipA4 = sbIPA4->value();

    hostIPAddress =
            (quint32(ipA1) << 3*8) +
            (quint32(ipA2) << 2*8) +
            (quint32(ipA3) << 1*8) +
            (quint32(ipA4) << 0*8);

    // Configure Server.
    emit serverConfigure(hostIPAddress, port);

    // Start Server.
    emit serverStart();
}

// 'Stop' Button is clicked.
void MainWindow::buttonStopClicked()
{
    QString journalRecord;

    // Stop Server.
    emit serverStop();

    // Add a Record to Journal.
    journalRecord = QString("Server is stopped.");
    journalRecordAdd(journalRecord);

    // Change Availability of Objects.
    // 1. Enable.
    pbStart->setDisabled(false);
    sbIPA1->setDisabled(false);
    sbIPA2->setDisabled(false);
    sbIPA3->setDisabled(false);
    sbIPA4->setDisabled(false);
    sbPort->setDisabled(false);
    // 2. Disable.
    pbStop->setDisabled(true);
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
