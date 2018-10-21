// 10.cpp.

#include <iostream>

class MyClass
{
public:
	MyClass(int i)
	{
		id = i;
		std::cout << "Constructor. ID=" << id << "\r\n";
	}
	
	~MyClass()
	{
		std::cout << "Destructor. ID=" << id << "\r\n";
	}
	
	int id;
	int x;
	int y;
};

int main()
{
	MyClass a(111);
	MyClass *b;
	
	a.x = 1;
	a.y = 2;
	
	b = new MyClass(222);
	b->x = 3;
	b->y = 4;
	delete b;
}
