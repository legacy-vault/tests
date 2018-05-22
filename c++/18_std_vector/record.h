// record.h.

#ifndef RECORD_H
#define RECORD_H

#include <string>

class Record
{
public:
	// Constructors.
	explicit Record();
	Record(std::string);
	Record(int);
	Record(std::string, int);
	Record(int, std::string);
	
	// Destructor.
	~Record();
	
	// Copy-Constructor.
	Record(const Record&);
	
	// Move-Constructor.
	Record(Record&&);
	
	// Copying Operator '='.
	Record& operator= (const Record&);
	
	// Moving Operator '='.
	Record& operator= (Record&&);
	
	// Fields.
	std::string name;
	int age;
	
	// Methods.
	void print();
	
	// Comparators.
	bool isLessByAge(const Record&);
	bool isLessByName(const Record&);
};

#endif // RECORD_H
