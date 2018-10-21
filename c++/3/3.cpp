// 3.cpp.

#include <iostream>
#include "3.h"

using namespace std;

myclass::myclass()
{
	a = 1;
}

myclass::myclass(int i)
{
	a = i;
}

int myclass::get_a()
{
	return a;
}

void myclass::show_a()
{
	int i;
	
	i = this->a;
	
	cout << i << "\r\n";
}

int spy_reader(myclass obj)
{
	int t;
	
	t = obj.a;
	
	return t;
}

void spy_writer(myclass *obj, int i)
{
	obj->a = i;
}

int main()
{
	myclass obj_1;
	myclass obj_2(2);
	myclass obj_3(3);
	myclass obj_4(4);
	
	int t;
	
	t = obj_1.get_a();
	cout << t << "\r\n";
	
	t = obj_2.get_a();
	cout << t << "\r\n";
	
	t = spy_reader(obj_3);
	cout << t << "\r\n";
	
	spy_writer(&obj_4, 444);
	t = obj_4.get_a();
	obj_4.show_a();
	
	return 0;
}
