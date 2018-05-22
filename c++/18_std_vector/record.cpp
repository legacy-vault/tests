// record.cpp.

#include "record.h"
#include <iomanip>
#include <iostream>
#include <string>

// Constructors.
Record::Record()
{
	name = std::string();
	age = 0;
}

Record::Record(std::string newName)
{
	name = newName;
	age = 0;
}

Record::Record(int newAge)
{
	name = std::string();
	age = newAge;
}

Record::Record(std::string newName, int newAge)
{
	name = newName;
	age = newAge;
}

Record::Record(int newAge, std::string newName)
{
	age = newAge;
	name = newName;
}

// Destructor.
Record::~Record()
{
	name.clear();
	age = 0;
}

// Copy-Constructor.
Record::Record(const Record& obj)
{
	name = obj.name;
	age = obj.age;
}

// Move-Constructor.
Record::Record(Record&& obj)
{
	*this = std::move(obj);
}

// Copying Operator '='.
Record& Record::operator= (const Record& obj)
{
	name = obj.name;
	age = obj.age;
	
	return *this;
}

// Moving Operator '='.
Record& Record::operator= (Record&& obj)
{
	name = obj.name;
	age = obj.age;
	
	return *this;
}

// Methods.

// Prints the Contents to the standart Output.
void Record::print()
{
	std::cout << "Age: [" << std::setw(3) << age << "], ";
	std::cout << "Name: [" << name << "].\r\n";
}

// Comparators.

// 1. By Age.
bool Record::isLessByAge(const Record& obj)
{
	if ((this->age) < (obj.age))
	{
		return true;
	}
	else
	{
		return false;
	}
}

// 2. By Name.
bool Record::isLessByName(const Record& obj)
{
	int result = (this->name).compare(obj.name);
	
	if (result < 0)
	{
		return true;
	}
	else
	{
		return false;
	}
}
