#include "my.h"
#include "mainwindow.h"
#include <QApplication>

MyJunk j1(11);
MyJunk *j2;
MyObject o1(111);
MyObject *o2;

int main(int argc, char *argv[])
{
    QApplication app(argc, argv);
    MainWindow win;
    win.show();

    j2 = new MyJunk(22);
    o1.junk = &j1;
    o2 = new MyObject(222);
    o2->junk = j2;

    Catcher goblin(999);

    QObject::connect(&o1, &MyObject::destroyed,
                     &goblin, &Catcher::catcher);
    QObject::connect(&j1, &MyJunk::destroyed,
                     &goblin, &Catcher::catcher);

    QObject::connect(o2, &MyObject::destroyed,
                     &goblin, &Catcher::catcher);
    QObject::connect(j2, &MyJunk::destroyed,
                     &goblin, &Catcher::catcher);

    // Automatic Deletion.
    j2->deleteLater();
    o2->deleteLater();

    // Manual Deletion.
    //delete j2;
    //delete o2;

    return app.exec();
}
