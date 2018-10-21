// 12.cpp.

#include "12.h"
#include <iostream>

int main()
{
	A obj_1;
	obj_1.aa = 11;
	
	B *obj_2;
	obj_2 = new B;
	obj_2->bb = 22;
	obj_2->ptrToA = &obj_1;
	
	std::cout << obj_1.aa << " " << obj_2->bb << " ";
	std::cout << obj_2->ptrToA << " " << (obj_2->ptrToA)->aa << "\r\n";
	
	delete obj_2;
	
	std::cout << obj_1.aa << "\r\n";
}
