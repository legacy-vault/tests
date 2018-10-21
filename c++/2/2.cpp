// 2.cpp.

#include <iostream>
#include "2.h"

using namespace std;

myclass::myclass()
{
	a = 1;
}

myclass::~myclass()
{
	a = 0;
}

int myclass::get_a()
{
	return a;
}

void myclass::set_a(int i)
{
	a = i;
}

int main()
{
	myclass ob1;
	int t;
	
	cout << ob1.a << "\r\n";
	
	ob1.set_a(8);
	t = ob1.get_a();
	cout << t << "\r\n";
	
	return 0;
}
