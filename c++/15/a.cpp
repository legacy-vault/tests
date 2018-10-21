// a.cpp.

#include <iostream>

class C
{
public:
	C();
	~C();
	C(const C &);
	int id;
};

C::C()
{
	std::cout << "Constructor.\r\n";
}

C::~C()
{
	std::cout << "Destructor.\r\n";
}

C::C(const C &copy)
{
	std::cout << "Copy Constructor.\r\n";
}

int main()
{
	C obj_1;				// Starts the Constructor.
	obj_1.id = 1;
	
	C obj_2 = obj_1;		// Starts the Copy Constructor.
	
	std::cout << ".\r\n";
	obj_2 = obj_1;
	
	return 0;
}
