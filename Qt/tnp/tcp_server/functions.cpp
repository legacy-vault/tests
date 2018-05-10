// functions.cpp.

#include "functions.h"

#include <QEventLoop>
#include <QTimer>

// Sleeps for some Time.
void sleep(int time)
{
    // Sleep for 'time' Milliseconds.
    QEventLoop loop;
    QTimer::singleShot(time, &loop, &QEventLoop::quit);
    loop.exec();
}
