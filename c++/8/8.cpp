// 8.cpp.

#include <iostream>
#include "8.h"

using namespace std;

template <class type_1, class type_2>
void myfunc(type_1 x, type_2 y)
{
	cout << x << " " << y << "\r\n";
}

int main()
{
	bool b;
	char *ch;
	char ch_1[] = "hello";
	int i;
	double d;
	
	b = true;
	ch = ch_1;
	i = 3;
	d = 16.42;
	
	myfunc(b, ch);
	myfunc(i, d);
	myfunc("abc", i);
	
	return 0;
}
