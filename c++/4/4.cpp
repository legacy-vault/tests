// 4.cpp.

#include <iostream>
#include "4.h"

using namespace std;

int main()
{
	int i;
	int* ptr;
	
	ptr = new int;
	if (!ptr)
	{
		cout << "Memory Allocation Error.";
	}
	cout << "Ptr:" << ptr << "\r\n";
	
	*ptr = 100;
	
	i = *ptr;
	cout << i << "\r\n";
	
	
	
	return 0;
}
