#include "mainwindow.h"
#include "ui_mainwindow.h"
#include <iostream>
#include <QThread>

// Constructor.
MainWindow::MainWindow(QWidget *parent) :
    QMainWindow(parent),
    ui(new Ui::MainWindow)
{
    ui->setupUi(this);

    this->workerIDLast = 0;
    this->workersCount = 0;

    // Bind Button.
    QPushButton *b;
    b = this->findChild<QPushButton *>("pushButton");
    QObject::connect(b, &QPushButton::clicked,
                     this, &MainWindow::StartManager);
}

// Destructor.
MainWindow::~MainWindow()
{
    delete ui;
}

// Starts a Manager, who starts an additional Worker (Thread).
void MainWindow::StartManager()
{
    QThread *t;
    Worker *w;
    int workerID;

    std::cout << "StartManager Entry.\r\n";

    (this->workerIDLast)++;
    workerID = this->workerIDLast;
    w = new Worker(workerID);
    t = new QThread;
    w->moveToThread(t);

    QObject::connect(t, &QThread::started,
                     w, &Worker::start);

    QObject::connect(w, &Worker::finished,
                     this, &MainWindow::WorkerFinished);
    QObject::connect(w, &Worker::finished,
                     t, &QThread::quit);
    QObject::connect(w, &Worker::finished,
                     w, &Worker::deleteLater);

    QObject::connect(t, &QThread::finished,
                     t, &QThread::deleteLater);

    t->start();

    (this->workersCount)++;
    std::cout << "MainWindow: Number of busy Workers = " <<
                 (this->workersCount) <<
                 ".\r\n";

    std::cout << "StartManager Exit.\r\n";
}

// Worker has finished his Job.
void MainWindow::WorkerFinished(int id)
{
    (this->workersCount)--;
    std::cout << "MainWindow: Worker[" << id << "] has finished.\r\n";

    std::cout << "MainWindow: Number of busy Workers = " <<
                 (this->workersCount) <<
                 ".\r\n";
}

// Worker's Job.
void Worker::start()
{
    quint64 i, iMax, j, jMax;
    QString s;

    iMax = 1 * 100; // 1 K.
    jMax = 1000 * 1000; // 1 M.
    std::cout << "Worker[" << (this->id) << "]: I started working!\r\n";

    for (i = 1; i <= iMax; i++)
    {
        std::cout << "Worker[" << (this->id) << "]: Stage " << i << ".\r\n";
        for (j = 1; j <= jMax; j++)
        {
            // Do Something...
            s = QString::number(j);
        }
    }

    emit finished(this->id);
}
