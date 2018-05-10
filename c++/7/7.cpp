// 7.cpp.

/*
	This is a simple Example of Operator Overload.
	
	It uses old and unsafe 'strlen' and 'strcpy' Functions instead of their 
	modern and safe Replacements.
*/

#include <cstring>
#include <iostream>
#include "7.h"

using namespace std;

// Constructor.
mystring::mystring(char *src)
{
	int tmp_len; // Number of Bytes allocated (with NULL Terminator).
	int tmp_num; // Number of useful Characters.
	char *tmp_content;
	char *tmp_copy_result;
	
	cout << "Constructor." << "\r\n";
	// Set the Indicator of non-initialized State.
	// In case of Error, these will be Markers of failed Initialization.
	this->capacity = 0;
	this->length = 0;
	
	// Allocate Memory.
	tmp_num = strlen(src);
	tmp_len = tmp_num + 1;
	tmp_content = new char[tmp_len];
	if (!tmp_content)
	{
		cout << "Memory Allocation Error." << "\r\n";
		exit(1); // No Error Handeler yet.
	}
	
	// Copy Data.
	tmp_copy_result = strcpy(tmp_content, src);
	if (tmp_copy_result != tmp_content)
	{
		cout << "Copy Error." << "\r\n";
		exit(1); // No Error Handeler yet.
	}
	
	// If Allocation and Copy were successful, 
	// Bind Memory with Object.
	this->content = tmp_content;
	this->capacity = tmp_len - 1; // Terminating NULL we do not count.
	this->length = this->capacity;
}

// Destructor.
mystring::~mystring()
{
	cout << "Destructor." << "\r\n";
	
	delete [] this->content;
	this->length = 0;
	this->capacity = 0;
}

// Operator '='.
mystring& mystring::operator= (mystring& right)
{
	cout << "Operator '='." << "\r\n";
	
	char *right_content;
	int right_len;
	int tmp_len;
	char *tmp_content;
	char *tmp_copy_result;
	
	// Cache Values.
	right_content = right.get_content();
	right_len = right.length;
	
	// Compare Length of new String with Capacity of existing String.
	if (right_len > this->capacity)
	{
		// We need to allocate new Memory and delete the used One.
		
		// Allocate new Memory.
		tmp_len = right_len + 1; // NULL-terminated.
		tmp_content = new char[tmp_len];
		if (!tmp_content)
		{
			cout << "Memory Allocation Error." << "\r\n";
			exit(1); // No Error Handeler yet.
		}
		
		// Copy Data.
		tmp_copy_result = strcpy(tmp_content, right_content);
		if (tmp_copy_result != tmp_content)
		{
			cout << "Copy Error." << "\r\n";
			exit(1); // No Error Handeler yet.
		}
		
		// Delete old Memory.
		delete this->content;
		
		// No Errors occurred.
		
		// Update Pointer to Content.
		this->content = tmp_content;
		
		// Update Length & Capacity.
		this->capacity = right_len;
		this->length = right_len;
	}
	else
	{
		// Use Capacity of existing String.
		
		// If Copy fails, we will know about it.
		this->length = 0;
		
		// Copy Data.
		tmp_copy_result = strcpy(this->content, right_content);
		if (tmp_copy_result != this->content)
		{
			cout << "Copy Error." << "\r\n";
			exit(1); // No Error Handeler yet.
		}
		
		// Garbage Cleaning.
		// ...
		
		// Update Length.
		this->length = right_len;
	}
	
	// Return, to allow Operator Chaining.
	return *this;
}

// Returns the Capacity of a String.
int mystring::get_capacity()
{
	return this->capacity;
}

// Returns the Content of a String.
char *mystring::get_content()
{
	return this->content;
}

// Returns the Length of a String.
int mystring::get_length()
{
	return this->length;
}

// Print Variable's Internals.
void mystring::show()
{
	cout << "[" << this->get_length() << "] ";
	cout << "[" << this->get_capacity() << "] ";
	cout << "[" << this->get_content() << "]";
	cout << "\r\n";
}

int main()
{
	mystring x("32 Symbols are here. 1234567890.");
	mystring y("33 Symbols are here. 1234567890+.");
	mystring z("13 Symbols...");
	mystring empty("");
	
	x.show();
	y.show();
	z.show();
	empty.show();
	
	// Use Operator '='.
	
	// 1.
	x = y;
	x.show();
	y.show();
	
	// 2.
	x = z;
	x.show();
	z.show();
	
	// 3.
	x = y = z;
	x.show();
	y.show();
	z.show();
	
	return 0;
}
