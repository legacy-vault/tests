// 11.h.

#include <iostream>

class A
{
public:
	explicit A(int iP);
	~A();
	
	int idParent;
};

class B : public A
{
public:
	explicit B(int iC, int iP);
	~B();
	
	int idChild;
};
