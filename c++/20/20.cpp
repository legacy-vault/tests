#include <iostream>

class Base
{
public:
	Base()
	{
		std::cout << "Base Constructor.\r\n";
	}
};

class Sub1 : public Base
{
	;
};

class Sub2 : public Base
{
	;
};

class Multi : public Sub1, public Sub2
{
	;
};

int main()
{
	Multi m;
	
	return 0;
}
