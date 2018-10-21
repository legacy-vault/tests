#include "worker.h"
#include <iostream>

Worker::Worker(int i)
{
    this->id = i;
}

Worker::~Worker()
{
    ;
}

void Worker::doWork(const QString &parameter)
{
    QString result;
    int i;

    for (i = 1; i <= 5; i++)
    {
        std::cout << "Worker #" << this->id << " is at step " << i << ".\r\n";
    }

    result = parameter + parameter;

    emit resultReady(result);
}

Controller::Controller(int id)
{
    Worker *worker;

    worker = new Worker(id);
    worker->moveToThread(&workerThread);

    QObject::connect(&workerThread, &QThread::finished,
                     worker, &QObject::deleteLater);

    QObject::connect(this, &Controller::operate,
                     worker, &Worker::doWork);

    QObject::connect(worker, &Worker::resultReady,
                     this, &Controller::handleResults);

    workerThread.start();
}

Controller::~Controller()
{
    workerThread.quit();
    workerThread.wait();
}

void Controller::handleResults(const QString &result)
{
    std::cout << "Controller. Result:" << result.toStdString() << ".\r\n";
}
