// main.h.

#ifndef MAIN_H
#define MAIN_H

#include <vector>

// Forward Declarations.
class Record;

int main();
void recordsShow(std::vector<Record>);
void recordsSortByAge(std::vector<Record> &);
void recordsSortByName(std::vector<Record> &);

#endif // MAIN_H
