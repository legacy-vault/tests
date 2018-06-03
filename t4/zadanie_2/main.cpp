// main.cpp.

#include <chrono>
#include <iostream>
#include <sstream>
#include <thread>

#include "rw_mutex.h"

void reader_func(int id);
void writer_func();

mcarcher::RWMutex rwm;
int xyz = 0;

int main()
{
	std::cout << "RWMutex Test." << std::endl;
	
	std::thread thr_writer_1(writer_func);
	std::thread thr_writer_2(writer_func);
	
	std::thread thr_reader_1(reader_func, 1);
	std::thread thr_reader_2(reader_func, 2);
	std::thread thr_reader_3(reader_func, 3);
	std::thread thr_reader_4(reader_func, 4);
	std::thread thr_reader_5(reader_func, 5);
	
	thr_writer_1.join();
	thr_writer_2.join();
	
	thr_reader_1.detach();
	thr_reader_2.detach();
	thr_reader_3.detach();
	thr_reader_4.detach();
	thr_reader_5.detach();
}

void writer_func()
{
	std::stringstream buf;
	int i;
	int i_max;
	std::thread::id id;
	
	id = std::this_thread::get_id();
	i_max = 1000;

	for (i = 1; i <= i_max; i++)
	{
		rwm.RWLock();
		
		// Simulate some Work.
		std::this_thread::sleep_for(std::chrono::milliseconds(10));
		xyz++;
		buf << "[Writer] ID=" << id << " xyz=" << xyz << "." << std::endl;
		std::cout << buf.str();
		buf.str(""); // Clear Buffer.
		
		rwm.RWUnlock();
	}
}

void reader_func(int id)
{
	std::stringstream buf;
	
	while (true)
	{
		rwm.RLock();
		
		// Simulate some Work.
		std::this_thread::sleep_for(std::chrono::milliseconds(3));
		
		rwm.RUnlock();
	}
}
