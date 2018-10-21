#include <iostream>

class Base
{
public:
	Base()
	{
		BaseConstructorRunsCount = 0;
		std::cout << "Base Constructor.\r\n";
		// Do Something useful here.
		// ...
		BaseConstructorRunsCount++;
	}
	Base(bool constructorIsEnabled)
	{
		std::cout << "Base Constructor X.\r\n";
		
		if (constructorIsEnabled == false)
		{
			return;
		}
		else
		{
			// Do Something useful here.
			// ...
		}
	}
	
private:
	int BaseConstructorRunsCount;
};

class Sub1 : public Base
{
public:
	Sub1() : Base()
	{
		std::cout << "Sub1 Constructor.\r\n"; //!
	}
};

class Sub2 : public Base
{
public:
	Sub2() : Base()
	{
		std::cout << "Sub2 Constructor.\r\n"; //!
	}
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
