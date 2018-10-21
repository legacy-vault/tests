#include "my.h"

#include <iostream>

#include <QObject>

MyObject::MyObject(int i, QObject *parent) :
    QObject(parent)
{
    id = i;
}

MyObject::~MyObject()
{
    std::cout << "Destructor. ID=" << id << "\r\n";
}

MyJunk::MyJunk(int i, QObject *parent) :
    QObject(parent)
{
    id = i;
}

MyJunk::~MyJunk()
{
    std::cout << "Destructor. ID=" << id << "\r\n";
}

Catcher::Catcher(int i, QObject *parent) :
    QObject(parent)
{
    id = i;
}

Catcher::~Catcher()
{
    std::cout << "Destructor. ID=" << id << "\r\n";
}

void Catcher::catcher(QObject *obj)
{
    extern MyJunk *j2;
    extern MyObject *o2;

    QObject *object = obj;
    std::cout << "Catcher.\r\n";

    MyJunk *j = j2;
    MyObject *o = o2;
}
