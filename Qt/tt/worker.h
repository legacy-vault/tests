#ifndef WORKER_H
#define WORKER_H

#include <QObject>
#include <QString>
#include <QThread>

class Worker : public QObject
{
    Q_OBJECT

public:
    Worker(int id);
    ~Worker();

public slots:
    void doWork(const QString &parameter);

signals:
    void resultReady(const QString &result);

private:
    int id;
};

class Controller : public QObject
{
    Q_OBJECT

public:
    Controller(int id);
    ~Controller();

public slots:
    void handleResults(const QString &);

signals:
    void operate(const QString &);

private:
    QThread workerThread;

};

#endif // WORKER_H
