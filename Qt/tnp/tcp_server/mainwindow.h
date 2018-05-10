// mainwindow.h.

#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QAbstractSocket>
#include <QKeyEvent>
#include <QLCDNumber>
#include <QMainWindow>
#include <QPushButton>
#include <QSpinBox>
#include <QString>
#include <QTextEdit>
#include <QtGlobal>
//#include <QWidget>

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
    void serverAcceptError(QAbstractSocket::SocketError socketError);
    void serverConnectionsIndicatorUpdate();
    void serverListeningFailed();
    void serverListeningStarted();
    void serverNewConnection();
    void serverNewMessage(QString msg);

private slots:
    void buttonStartClicked();
    void buttonStopClicked();

signals:
    void serverConfigure(quint32 srvIPAddress, quint16 srvPort);
    void serverStart();
    void serverStop();

private:
    // Methods.
    void keyPressEvent(QKeyEvent *event);

    // Journal.
    QTextEdit *teJournal;

    // Buttons.
    QPushButton *pbStart;
    QPushButton *pbStop;

    // Inputs.
    QSpinBox *sbIPA1;
    QSpinBox *sbIPA2;
    QSpinBox *sbIPA3;
    QSpinBox *sbIPA4;
    QSpinBox *sbPort;

    // Indicators.
    QLCDNumber *lcdConnections;

    // UI.
    Ui::MainWindow *ui;
};

#endif // MAINWINDOW_H
