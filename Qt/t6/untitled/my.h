#ifndef MY_H
#define MY_H

#include <QObject>

// F.D.
class MyJunk;

class MyObject : public QObject
{
    Q_OBJECT
public:
    explicit MyObject(int id, QObject *parent = Q_NULLPTR);
    virtual ~MyObject();
    MyJunk *junk;
    int id;
};


class MyJunk : public QObject
{
    Q_OBJECT
public:
    MyJunk(QObject *parent = Q_NULLPTR);
    MyJunk(int id, QObject *parent = Q_NULLPTR);
    virtual ~MyJunk();
    int id;
};

class Catcher : public QObject
{
    Q_OBJECT

public:
    explicit Catcher(int id, QObject *parent = Q_NULLPTR);
    virtual ~Catcher();
    int id;
public slots:
    void catcher(QObject *obj);
};

#endif // MY_H
