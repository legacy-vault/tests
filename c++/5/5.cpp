// 5.cpp.

#include <iostream>
#include "5.h"

using namespace std;

int main()
{
	int i;
	
	f(i);
	cout << i << "\r\n";
	
	return 0;
}

void f(int &n)
{
	n = 100;
}
