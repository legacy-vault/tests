// 13.cpp.

#include <chrono>
#include <condition_variable>
#include <iostream>
#include <thread>
#include <mutex>

std::condition_variable cv;
std::mutex cv_m;

void func_1()
{
	std::thread::id id;
	std::unique_lock<std::mutex> ul(cv_m);
	
	id = std::this_thread::get_id();
	std::cout << "[Thread #1, ID=" << id << "] Has Started.\r\n";
	cv.wait(ul);
	
	std::cout << "[Thread #1] Has Ended.\r\n";
}

void func_2()
{
	std::thread::id id;
	std::chrono::milliseconds dur(5000);
	std::unique_lock<std::mutex> ul(cv_m);
	
	id = std::this_thread::get_id();
	std::cout << "[Thread #2, ID=" << id << "] Has Started.\r\n";
	
	std::cout << "[Thread #2] Has started sleepping...\r\n";
	std::this_thread::sleep_for(dur);
	std::cout << "[Thread #2] Woke up!\r\n";
	cv.notify_all();
	
	std::cout << "[Thread #2] Has Ended.\r\n";
}

int main()
{
	std::thread t_1(func_1);
	std::thread t_2(func_2);
	
	t_1.join();
	t_2.detach();
}
