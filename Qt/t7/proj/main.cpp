#include "mainwindow.h"
#include <QApplication>
#include <QWebEngineView>
#include <QWidget>
#include <QString>
#include <QScrollArea>
#include <QObject>
#include <QUrl>
#include <QSizePolicy>
#include <QCursor>

int main(int argc, char *argv[])
{
    QApplication app(argc, argv);
    MainWindow win;

    QWidget *browser;
    QString browserName;
    QScrollArea *scrollArea;
    QString scrollAreaName;
    QWebEngineView *view;
    QSizePolicy sizePolicy;

    sizePolicy.setHorizontalPolicy(QSizePolicy::Expanding);
    sizePolicy.setVerticalPolicy(QSizePolicy::Expanding);

    scrollAreaName = QString("scrollArea");
    scrollArea = win.findChild<QScrollArea *>(scrollAreaName);
    browserName = QString("widget");
    browser = win.findChild<QWidget *>(browserName);
    //view = new QWebEngineView(browser);
    view = new QWebEngineView(scrollArea);
    //scrollArea->setWidget(browser);
    //scrollArea->setWidget(view);

    scrollArea->setWidgetResizable(true);
    view->setCursor(Qt::WhatsThisCursor);
    view->setSizePolicy(sizePolicy);
    browser->setSizePolicy(sizePolicy);

    win.show();
    view->load(QUrl("https://ya.ru"));
    view->show();

    QObject::connect
            (&win,
             &MainWindow::changeUrl,
             view,
             // We can not use simple '&QWebEngineView::load' here.
             static_cast<void (QWebEngineView::*)(const QUrl &)>(&QWebEngineView::load)
             );

    return app.exec();
}
