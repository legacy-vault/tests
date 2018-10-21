#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QMainWindow>

namespace Ui
{
    class MainWindow;
}

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    explicit MainWindow(QWidget *parent = 0);
    ~MainWindow();

public slots:
    void StartManager();
    void WorkerFinished(int id);

private:
    Ui::MainWindow *ui;
    int workerIDLast; // ID of Worker who last received the Job.
    int workersCount; // Number of Workers who are busy.
};

class Worker : public QObject
{
    Q_OBJECT

public:
    explicit Worker(int i) { this->id = i; }
    ~Worker(){}

public slots:
    void start();

signals:
    void finished(int id);

private:
    int id;

};

#endif // MAINWINDOW_H
