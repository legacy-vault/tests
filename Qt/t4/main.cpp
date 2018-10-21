#include <iostream>
#include "mainwindow.h"
#include <QApplication>

int main(int argc, char *argv[])
{
    QApplication app(argc, argv);
    MainWindow w;
    w.show();

    MyCounter mca, mcb;
    QObject::connect(&mca, &MyCounter::valueChanged,
                     &mcb, &MyCounter::setValue);
    mca.setValue(12);
    std::cout << mca.value() << " " << mcb.value() << "\r\n";

    mcb.setValue(48);
    std::cout << mca.value() << " " << mcb.value() << "\r\n";

    return app.exec();
}
