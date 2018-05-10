// 7.h.

class mystring
{
	int capacity; // Terminating NULL is not counted.
	char* content;
	int length; // Terminating NULL is not counted.

public:

	mystring(char*);
	~mystring();

	mystring& operator= (mystring& right);

	int get_capacity();
	char* get_content();
	int get_length();
	void show();
};
