// 14.cpp.

#include <iostream>
#include <future>
#include <thread>

int worker()
{
	int result;
	std::chrono::milliseconds dur(3000);
	
	std::this_thread::sleep_for(dur);
	result = 123;
	
	return result;
}

int main()
{
	std::future<int> result_tmp;
	int result_real;
	
	// Start Worker.
	result_tmp = std::async(worker);
	std::cout << "Worker has been started." << "\r\n";
	
	// Wait for the Result.
	std::cout << "Waiting for Result from Worker..." << "\r\n";
	result_real = result_tmp.get();
	
	std::cout << result_real << "\r\n";
}
