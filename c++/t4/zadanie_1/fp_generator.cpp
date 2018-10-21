// fp_generator.cpp.

#include <cstdlib>
#include <cstring>
#include <ctime>
#include <fstream>
#include <iostream>
#include <limits>
#include <random>

#include "fp_generator.h"

int main()
{
	int result;

	result = fp_generate(file_out_path, file_size_max);

	if (result != error_none)
	{
		std::cout << "Error! Error Code = [" << result << "]." << std::endl;
	}

	return result;
}

int fp_generate(const char *dst_path, unsigned int file_size_limit)
{
	// Temporary Data.
	std::ofstream file_out;
	int result_fclose;
	int result_fopen;
	int rndSignFlag = rand() % 2;
	unsigned int file_size;
	unsigned int file_size_new;

	// Buffer for String Representation of 'double' Value.
	// -0.1234567890123456e+308 and terminating "\0". 24+1 Symbols.
	char buffer[25]; // Stores maximum 16 Digits after Point.

	// Random Generator.
	std::default_random_engine rndGenerator;
	double rndValue;
	unsigned int rndValueTextSize;
	constexpr double rndValueNaturalMax = std::numeric_limits<double>::max();
	constexpr double rndValueNaturalMin = 0;
	std::uniform_real_distribution<double> rndDistribution
        (
            rndValueNaturalMin,
            rndValueNaturalMax
        );

	// Open File.
	file_out.open(dst_path, std::ios_base::out | std::ios_base::trunc);
	result_fopen = file_out.is_open();
	if (result_fopen != 1)
	{
		std::cout << msg_error_fopen << std::endl;

		return error_fopen;
	}

	// Generate random Numbers of 'double' Type.
	std::cout << "Creating random Numbers in the Range: [" <<
		(-rndValueNaturalMax) << " : " << rndValueNaturalMax << "]." <<
		std::endl << "Please, wait..." << std::flush;

	srand(time(0));
	file_size = 0;

	while (true)
	{
		// While C++ Standard limits our random Numbers,
		// we do it in two Parts:
		// first, create random natural Number,
		// then create random Sign (+ or -).
		rndValue = rndDistribution(rndGenerator);
		rndSignFlag = rand() % 2;
		if (rndSignFlag == 1)
		{
			rndValue = -rndValue;
		}

		// Create a String Representation of 'double' Type.
		// Maximum available 16 Digits after Point are stored.
		rndValueTextSize = sprintf(buffer, "%.16e", rndValue);

		// Write to File.
		file_size_new = file_size + rndValueTextSize + numbers_separator_length;
		if (file_size_new <= file_size_limit)
		{
            file_out << buffer << numbers_separator;
            file_size = file_size_new;
		}
		else
		{
            break;
		}
	}
	std::cout << "Done." << std::endl;

	// Close File.
	file_out.close();
	result_fclose = file_out.is_open();
	if (result_fclose != 0)
	{
		std::cout << msg_error_fclose << std::endl;

		return error_fclose;
	}

	return error_none;
}
