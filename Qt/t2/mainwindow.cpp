#include <QString>
#include "mainwindow.h"
#include "ui_mainwindow.h"

MainWindow::MainWindow(QWidget *parent) :
    QMainWindow(parent),
    ui(new Ui::MainWindow)
{
    ui->setupUi(this);
}

MainWindow::~MainWindow()
{
    delete ui;
}

void MainWindow::on_pushButtonStart_clicked()
{
    QLineEdit *lineEdit;
    QString lineEditName;

    QPushButton *pushButtonStart;
    QString pushButtonStartName;

    QPushButton *pushButtonStop;
    QString pushButtonStopName;

    // Find Objects.
    lineEditName = QString("lineEdit");
    lineEdit = this->findChild<QLineEdit *>(lineEditName);

    pushButtonStartName = QString("pushButtonStart");
    pushButtonStart = this->findChild<QPushButton *>(pushButtonStartName);

    pushButtonStopName = QString("pushButtonStop");
    pushButtonStop = this->findChild<QPushButton *>(pushButtonStopName);

    // Change 'Enabled' Status of Objects.
    pushButtonStop->setDisabled(false);
    lineEdit->setDisabled(true);
    pushButtonStart->setDisabled(true);
}

void MainWindow::on_pushButtonStop_clicked()
{
    QLineEdit *lineEdit;
    QString lineEditName;

    QPushButton *pushButtonStart;
    QString pushButtonStartName;

    QPushButton *pushButtonStop;
    QString pushButtonStopName;

    // Find Objects.
    lineEditName = QString("lineEdit");
    lineEdit = this->findChild<QLineEdit *>(lineEditName);

    pushButtonStartName = QString("pushButtonStart");
    pushButtonStart = this->findChild<QPushButton *>(pushButtonStartName);

    pushButtonStopName = QString("pushButtonStop");
    pushButtonStop = this->findChild<QPushButton *>(pushButtonStopName);

    // Change Availability of Objects.
    pushButtonStart->setDisabled(false);
    lineEdit->setDisabled(false);
    pushButtonStop->setDisabled(true);
}
