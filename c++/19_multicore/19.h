// 19.h.

#ifndef FILE_19_H
#define FILE_19_H

#include <cstddef>	/* For std::size_t */
#include <cstdint>
#include <vector>

typedef uint16_t Uint16;
typedef unsigned long long int Uint64;
typedef std::vector<Uint16> Data;

struct Task
{
	std::size_t cpusCount;
	int cpuNumPlanned;
	Data *data;
	char id;
	Uint64 indexFirst;
	Uint64 indexLast;
	bool segmentIsNotFirst;	//	If true, Worker will not compare the
							//	first Element of Data with the Element
							//	which is before the first Data Element.
};

struct Result
{
	bool error;
	bool duplicateIsFound;
	std::vector<Uint64> indices;
	Uint16 elementMax;
	Uint16 elementMin;
	Uint64 elementOccurrence[UINT16_MAX + 1];	// How many Times each Element 
												// is seen in Data Segment.
};

struct Total
{
    std::vector<Uint64> indices;
	Uint16 elementMax;
	Uint16 elementMin;
	Uint64 elementOccurrence[UINT16_MAX + 1];	// How many Times each Element 
												// is seen in all Data.
};

const int appExitCodeError = 1;
const int appExitCodeSuccess = 0;
const Uint64 dataCount = 1000 * 1000 * 1000;
const int threadsCount = 4;
const bool verbose = false;
const char *outputFileName = "element_occurrence.csv";

int main();
void generateData(Data&, Uint64);
void worker(Task task);

#endif // FILE_19_H
