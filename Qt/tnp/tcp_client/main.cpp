// main.cpp.

#include "main.h"
#include "client.h"
#include "protocol.h"
#include "mainwindow.h"

#include <QApplication>
#include <QObject>
#include <QTcpSocket>

Client client;

int main(int argc, char *argv[])
{
    QApplication app (argc, argv);
    MainWindow win;
    char protoVerMajor;
    char protoVerMinor;

    // Check Protocol Version.
    protoVerMajor = Protocol::PROTOCOL_VERSION_MAJOR;
    protoVerMinor = Protocol::PROTOCOL_VERSION_MINOR;
    if ((protoVerMajor != 1) || (protoVerMinor != 0))
    {
        return(1);
    }

    // Connect Signals with Slots.

    // 1. Signals from GUI to Client.
    QObject::connect(&win, &MainWindow::clientAbort,
                     &client, &Client::abort);

    QObject::connect(&win, &MainWindow::clientConfigure,
                     &client, &Client::configure);

    QObject::connect(&win, &MainWindow::clientConnect,
                     &client, &Client::connect);

    QObject::connect(&win, &MainWindow::clientDisconnect,
                     &client, &Client::disconnect);

    QObject::connect(&win, &MainWindow::messageIsPrepared,
                     &client, &Client::messageSend);

    // 2. Signals from Client to GUI.
    QObject::connect(&client, &Client::clientError,
                     &win, &MainWindow::clientError);

    QObject::connect(&client, &Client::clientHostIsFound,
                     &win, &MainWindow::clientHostIsFound);

    QObject::connect(&client, &Client::clientIsAborted,
                     &win, &MainWindow::clientIsAborted);

    QObject::connect(&client, &Client::clientIsConnected,
                     &win, &MainWindow::clientIsConnected);

    QObject::connect(&client, &Client::clientIsDisconnected,
                     &win, &MainWindow::clientIsDisconnected);

#ifndef QT_NO_NETWORKPROXY
    QObject::connect(&client, &Client::clientProxyAuthenticationIsRequired,
                     &win, &MainWindow::clientProxyAuthenticationIsRequired);
#endif

    QObject::connect(&client, &Client::clientStateIsChanged,
                     &win, &MainWindow::clientStateIsChanged);

    QObject::connect(&client, &Client::sendDone,
                     &win, &MainWindow::sendDone);

    QObject::connect(&client, &Client::sendFailed,
                     &win, &MainWindow::sendFailed);

    QObject::connect(&client, &Client::serverReplyError,
                     &win, &MainWindow::serverReplyError);

    QObject::connect(&client, &Client::serverReplyTimeout,
                     &win, &MainWindow::serverReplyTimeout);

    QObject::connect(&client, &Client::serverReplyUnexpected,
                     &win, &MainWindow::serverReplyUnexpected);

    win.show();

    return app.exec();
}
