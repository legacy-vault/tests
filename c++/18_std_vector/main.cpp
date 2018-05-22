// main.cpp.

#include "main.h"
#include "record.h"
#include <algorithm>
#include <iostream>
#include <string>
#include <vector>

int main()
{
	// Initialize the Data.
	Record r1;
	Record r2(10);
	std::string tmpName = std::string("John Smith");
	Record r3(tmpName);
	tmpName = std::string("Jack Sparrow");
	Record r4(tmpName, 20);
	Record r5(30, tmpName);
	tmpName = std::string("Gordon Freeman");
	Record r6(tmpName, 33);
	tmpName = std::string("Harry Potter");
	Record r7(tmpName, 25);

	// Array (dynamic) Initialization.
	std::vector<Record> records;
	records.push_back(r1);
	records.push_back(r2);
	records.push_back(r3);
	records.push_back(r4);
	records.push_back(r5);
	records.push_back(r6);
	records.push_back(r7);
	recordsShow(records);

	// Sort Records by Age (minimum is first).
	recordsSortByAge(records);
	std::cout << "Sorted by Age.\r\n";
	recordsShow(records);

	// Sort Records by Name (small is first).
	recordsSortByName(records);
	std::cout << "Sorted by Name.\r\n";
	recordsShow(records);

	return 0;
}

void recordsShow(std::vector<Record> recs)
{
	std::vector<Record>::iterator item;
	std::vector<Record>::iterator itemLast;

	item = recs.begin();
	itemLast = recs.end();

	while (item != itemLast)
	{
		(*item).print();

		item++;
	}

	std::cout << "\r\n";
}

void recordsSortByAge(std::vector<Record> &recs)
{
	std::vector<Record>::iterator itemFirst;
	std::vector<Record>::iterator itemLast;

	itemFirst = recs.begin();
	itemLast = recs.end();

	struct
	{
        bool operator() (Record a, Record b) const
        {
			bool result = a.isLessByAge(b);

            return result;
        }
    }
    comparator;

	std::sort(itemFirst, itemLast, comparator);
}

void recordsSortByName(std::vector<Record> &recs)
{
    std::vector<Record>::iterator itemFirst;
	std::vector<Record>::iterator itemLast;

	itemFirst = recs.begin();
	itemLast = recs.end();

	struct
	{
        bool operator() (Record a, Record b) const
        {
			bool result = a.isLessByName(b);

            return result;
        }
    }
    comparator;

	std::sort(itemFirst, itemLast, comparator);
}
