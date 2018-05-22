// 16_a.cpp.

#include <algorithm>
#include <chrono>
#include <iostream>
#include <iterator>
#include <ratio>
#include <vector>
 
int main ()
{
	int i;
	int j;
	int j_max;
	int el_a;
	int el_b;
	int t;
	int tests_count;
	std::chrono::time_point<std::chrono::high_resolution_clock> time_1;
	std::chrono::time_point<std::chrono::high_resolution_clock> time_2;
	std::chrono::duration<unsigned long long int, std::milli> time_delta;
	std::vector<int> v;
	int v_size;
	
	v = std::vector<int> {7, 5, 16, 8, 43, 15, 86, 30, 15, 82, 16, 9, 93, 71, 3};
	v_size = v.size();
	j_max = v_size - 1;
	tests_count = 1000 * 1000 * 1;
	
	// Show.
	for (i = 0; i < v_size; i++)
	{
		std::cout << v[i] << " ";
	}
	std::cout << "\r\n";
	
	time_1 = std::chrono::high_resolution_clock::now();
	for (t = 1; t <= tests_count; t++)
	{
		v = std::vector<int> {7, 5, 16, 8, 43, 15, 86, 30, 15, 82, 16, 9, 93, 71, 3};
		
		// Sort (Bubble).
		for (i = 0; i < v_size; i++)
		{
			for (j = 0; j <= j_max; j++)
			{
				el_a = v[j];
				el_b = v[j+1];
				
				if (el_a > el_b)
				{
					v[j] = el_b;
					v[j+1] = el_a;
				}
			}
			
			if (std::is_sorted(std::begin(v), std::end(v)))
			{
				break;
			}
		} // End of Sort.
	} // End of Tests Loop.
	time_2 = std::chrono::high_resolution_clock::now();
	
	// Show.
	for (i = 0; i < v_size; i++)
	{
		std::cout << v[i] << " ";
	}
	std::cout << "\r\n";
	
	time_delta = std::chrono::duration_cast<std::chrono::milliseconds>(time_2 - time_1);
	std::cout << time_delta.count() << " ms.\r\n";
	
	return 0;
}
