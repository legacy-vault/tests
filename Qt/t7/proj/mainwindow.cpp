#include "mainwindow.h"
#include "ui_mainwindow.h"

#include <QLineEdit>
#include <QUrl>

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

void MainWindow::on_pushButton_clicked()
{
    QLineEdit *leUrl;
    QString urlStr;
    QUrl url;

    leUrl = this->findChild<QLineEdit *>(QString("lineEdit"));
    urlStr = leUrl->text();
    url = QUrl(urlStr);

    emit changeUrl(url);
}
