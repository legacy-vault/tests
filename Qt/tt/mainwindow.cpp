#include "mainwindow.h"
#include "ui_mainwindow.h"
#include "worker.h"
#include <iostream>

MainWindow::MainWindow(QWidget *parent) :
    QMainWindow(parent),
    ui(new Ui::MainWindow)
{
    ui->setupUi(this);

    QPushButton *b;
    b = this->findChild<QPushButton *>("pushButton");
    QObject::connect(b, &QPushButton::clicked,
                     this, &MainWindow::StartManager);
}

MainWindow::~MainWindow()
{
    delete ui;
}

void MainWindow::StartManager()
{
    QString s1, s2;

    std::cout << "StartManager.\r\n";

    Controller c1(1);
    Controller c2(2);

    s1 = QString("s1");
    s2 = QString("s2");

    c1.operate(s1);
    c2.operate(s2);
}
