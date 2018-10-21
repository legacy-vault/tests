// 6.h.

class myclass
{
	char *text;
	
public:
	
	myclass(char *);				// Constructor.
	myclass(const myclass &obj);	// Copy Constructor.
	~myclass();						// Destructor.
	
	char *get_text();
	void show_text();
};

void my_show(myclass);
void my_show_2(myclass&);
