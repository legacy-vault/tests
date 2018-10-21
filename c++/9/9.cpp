// 9.cpp.

#include <iostream>

class MyClass
{
	
public:
	
	int i;
	static int j;
};

int MyClass::j;

int main()
{
	MyClass obj_1;
	MyClass obj_2;
	
	obj_1.i = 1;
	obj_2.i = 2;
	
	obj_1.j = 3;
	
	std::cout << obj_1.j << "\r\n";
	
	MyClass::j = 4;
	
	std::cout << obj_2.j << "\r\n";
	
	return 0;
}
