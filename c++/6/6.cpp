// 6.cpp.

#include <iostream>
#include <cstring>
#include "6.h"

using namespace std;

int main()
{
	myclass x("Hello.");
	
	x.show_text();
	
	my_show(x);		// Triggers the Copy.
	my_show_2(x);	// No Copy is created.
	
	return 0;
}

myclass::myclass(char *src)
{
	cout << "Constructor." << "\r\n";
	
	int len = strlen(src) + 1; // NULL-terminated.
	
	text = new char[len];
	if (!text)
	{
		cout << "Memory Allocation Error." << "\r\n";
		exit(1);
	}
	
	strcpy(text, src);
}

myclass::myclass(const myclass &obj)
{
	cout << "Copy Constructor." << "\r\n";
	
	int len = strlen(obj.text) + 1; // NULL-terminated.
	
	text = new char[len];
	if (!text)
	{
		cout << "Memory Allocation Error." << "\r\n";
		exit(1);
	}
	
	strcpy(text, obj.text);
}

myclass::~myclass()
{
	cout << "Destructor." << "\r\n";
	
	delete [] text;
}
	
char *myclass::get_text()
{
	return text;
}

void myclass::show_text()
{
	cout << text << "\r\n";
}

void my_show(myclass obj)
{
	char *t;
	
	t = obj.get_text();
	cout << t << "\r\n";
}

void my_show_2(myclass &obj)
{
	char *t;
	
	t = obj.get_text();
	cout << t << "\r\n";
}
