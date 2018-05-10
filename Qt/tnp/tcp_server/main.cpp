// main.cpp.

#include "server.h"
#include "mainwindow.h"
#include "protocol.h"
#include "main.h"

#include <QApplication>
#include <QObject>

Server server;

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

    // 1. Signals from GUI to Server.
    QObject::connect(&win, &MainWindow::serverConfigure,
                     &server, &Server::configure);

    QObject::connect(&win, &MainWindow::serverStart,
                     &server, &Server::start);

    QObject::connect(&win, &MainWindow::serverStop,
                     &server, &Server::stop);

    // 2. Signals from Server to GUI.
    QObject::connect(&server, &Server::acceptError,
                     &win, &MainWindow::serverAcceptError);

    QObject::connect(&server, &Server::listeningFailed,
                     &win, &MainWindow::serverListeningFailed);

    QObject::connect(&server, &Server::listeningStarted,
                     &win, &MainWindow::serverListeningStarted);

    QObject::connect(&server, &Server::newConnection,
                     &win, &MainWindow::serverNewConnection);

    QObject::connect(&server, &Server::newMessage,
                     &win, &MainWindow::serverNewMessage);

    QObject::connect(&server, &Server::connectionsIndicatorUpdate,
                     &win, &MainWindow::serverConnectionsIndicatorUpdate);

    win.show();

    return app.exec();
}
