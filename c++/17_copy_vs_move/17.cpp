// 17.cpp.

#include "17.h"

#include <iostream>

C::C()
{
	std::cout << "Constructor.\r\n";
	
	this->id = 999;
}

C::~C()
{
	std::cout << "Destructor.\r\n";
}

C::C(const C& obj)
{
	std::cout << "Copy-Constructor.\r\n";
	
	this->id = obj.id;
}

C::C(C&& obj)
{
	std::cout << "Move-Constructor.\r\n";
	
	this->id = obj.id;
}

C& C::operator=(const C& obj)
{
	std::cout << "Copying Operator '='.\r\n";
	
	this->id = obj.id;
	
	return *this;
}

C& C::operator=(C&& obj)
{
	std::cout << "Moving Operator '='.\r\n";
	
	this->id = obj.id;
	
	return *this;
}

C RValueGenerator(C obj)
{
	return obj;
}

int main()
{
	C obj_1;						// Starts the Constructor.
	obj_1.id = 1;
	
	std::cout << "\r\n";
	
	// 1.
	std::cout << "[1]\r\n";
	C obj_2 = obj_1;				// Starts the Copy-Constructor.
	std::cout << "obj_2.id = " << obj_2.id << ".\r\n";
	std::cout << "obj_1.id = " << obj_1.id << ".\r\n";
	
	std::cout << "\r\n";
	
	// 2.
	std::cout << "[2]\r\n";
	obj_2 = obj_1;					// Copying Operator '='.
	obj_2.id = 2;
	
	std::cout << "\r\n";
	
	// 3.
	std::cout << "[3]\r\n";
	C obj_3 = std::move(obj_2);		// Starts the Move-Constructor.
	std::cout << "obj_3.id = " << obj_3.id << ".\r\n";
	std::cout << "obj_2.id = " << obj_2.id << ".\r\n";
	
	std::cout << "\r\n";
	
	// 4.
	std::cout << "[4]\r\n";
	C obj_4;
	obj_4 = RValueGenerator(C());	// Starts the Move-Operator=.
	std::cout << "obj_4.id = " << obj_4.id << ".\r\n";
	
	std::cout << "\r\n";
	
	// 5.
	std::cout << "[5]\r\n";
	C obj_5;
	obj_5 = std::move(obj_2);	// Starts the Move-Operator=.
	std::cout << "obj_5.id = " << obj_5.id << ".\r\n";
	std::cout << "obj_2.id = " << obj_2.id << ".\r\n";
	
	std::cout << "\r\n";
	
	return 0;
}
