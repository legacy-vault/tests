// 11.cpp.

#include "11.h"

A::A(int iP)
{
	idParent = iP;
	std::cout << "Constructor. ID=" << idParent << "\r\n";
}

A::~A()
{
	std::cout << "Destructor. ID=" << idParent << "\r\n";
}


B::B(int iC, int iP)
	: A(iP)
{
	idChild = iC;
	std::cout << "Constructor. ID=" << idChild << "\r\n";
}
	
B::~B()
{
	std::cout << "Destructor. ID=" << idChild << "\r\n";
}

int main()
{
	A obj_1(111);
	B obj_2(222, 333);
	A *obj_3;
	B *obj_4;
	
	obj_3 = new A(444);
	delete obj_3;
	
	obj_4 = new B(555, 666);
	delete obj_4;
}
