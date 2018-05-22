// 17.h.

#ifndef FILE_17_H
#define FILE_17_H

class C
{
public:
	C();					// Constructor.
	~C();					// Destructor.
	C(const C&);			// Copy-Constructor.
	C(C&&);					// Move-Constructor.
	C& operator=(const C&);	// Operator '=' that uses Copy.
	C& operator=(C&&);		// Operator '=' that uses Move.
	
	int id;
};

C RValueGenerator(C);

#endif // FILE_17_H
