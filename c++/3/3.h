// 3.h.

class myclass
{
	int a;
	
public:
	myclass();
	myclass(int i);
	
	int get_a();
	void show_a();
	
	friend int spy_reader(myclass obj);
	friend void spy_writer(myclass *obj, int i);
};
