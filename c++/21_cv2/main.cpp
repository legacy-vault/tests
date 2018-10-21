// main.cpp.

#include <atomic>
#include <chrono>
#include <condition_variable>
#include <iostream>
#include <mutex>
#include <thread>

std::condition_variable cv;
std::mutex cv_mutex;

void f_a();
void f_b();

std::atomic<int> counter;

int main()
{
	counter = 0;
	
	std::thread thr_a(f_a);
	std::thread thr_b(f_b);
	
	thr_b.join();
	thr_a.detach();
}

void f_a()
{
	std::unique_lock<std::mutex> ul(cv_mutex);
	
	std::cout << "A started Sleep.\r\n";
	
	std::this_thread::sleep_for(std::chrono::seconds(1));
	
	counter++;
	
	std::cout << "A finished Sleep and is notifying all...\r\n";
	
	cv.notify_all();
	
	std::cout << "A finished notify.\r\n";
}

void f_b()
{
	std::unique_lock<std::mutex> ul(cv_mutex);
	
	std::cout << "B started Sleep.\r\n";
	
	std::this_thread::sleep_for(std::chrono::seconds(3));
	
	std::cout << "B finished Sleep and is waiting...\r\n";
	
	if (counter == 0)
	{
		cv.wait(ul);
	}
	
	std::cout << "B finished waiting.\r\n";
}
