// 19.cpp.

#include "19.h"
#include <cstddef>		/* For std::size_t */
#include <fstream>
#include <future>
#include <iostream>
#include <random>
#include <sstream>
#include <thread>
#include <vector>

#ifdef __linux__
	#include <sched.h>
#endif

std::promise<Result> pipes[threadsCount];

int main()
{
	std::size_t cpusCount;
	std::size_t cpuIter;
	Data data;
	bool error;
	std::ofstream fileOut;
	Uint64 i;
	Uint64 ii;
	Uint64 index;
	int j;
	int jMax;
	int k;
	int kMax;
	Uint64 n;
	std::future<Result> results_tmp[threadsCount];
	Result results[threadsCount];
	Task tasks[threadsCount];
	std::thread threads[threadsCount];
	Total total; // Total Results from all Threads.

	// Prepare Data.
	cpusCount = std::thread::hardware_concurrency();
	std::cout << "CPU Count = [" << cpusCount << "].\r\n";
	std::cout << "Generating Data. Please, wait..." << std::flush;
	generateData(data, dataCount);
	if (verbose)
	{
		for (Uint64 i=0; i<dataCount; i++)
		{
			std::cout << data[i] << " ";
		}
		std::cout << "\r\n\r\n";
	}
	std::cout << "Done.\r\n";

	
	// Divide Task into Parts.
	std::cout << "Creating Tasks..." << std::flush;
	if (dataCount < (2 * threadsCount))
	{
		std::cout << "Data is not not enough!\r\n";
		return appExitCodeError;
	}
	n = dataCount / threadsCount;
	if ((dataCount % threadsCount) != 0)
	{
		n++;
	}
	i = 0;
	cpuIter = 0;
	jMax = threadsCount - 1;
	for (j=0; j<threadsCount; j++)
	{
		tasks[j].id = j+1;
		tasks[j].indexFirst = i;

		// All Segments except the last one always have equal Size.
		// The last Segment may be a bit shorter if all the Data can not be
		// equally divided by all Threads. E.g. 7 = 4 + 3, where 4 != 3.
		if (j != jMax)
		{
			tasks[j].indexLast = i + n - 1;
		}
		else
		{
			tasks[j].indexLast = i + (dataCount - j*n) - 1;
		}

		tasks[j].data = &data;
		if (j != 0)
		{
			tasks[j].segmentIsNotFirst = true;
		}
		else
		{
			tasks[j].segmentIsNotFirst = false;
		}
		
		tasks[j].cpuNumPlanned = cpuIter;
		tasks[j].cpusCount = cpusCount;

		// Next.
		i = i + n;
		cpuIter++;
		if (cpuIter >= cpusCount)
		{
			cpuIter = 0;
		}
	}
	std::cout << "Done.\r\n";

	// Start Workers.
	std::cout << "Starting Workers..." << std::flush;
	for (j=0; j<threadsCount; j++)
	{
		threads[j] = std::thread(worker, tasks[j]);
		threads[j].detach();
	}
	std::cout << "Done.\r\n";

	// Wait for Results.
	for (j=0; j<threadsCount; j++)
	{
		results_tmp[j] = pipes[j].get_future();
	}
	for (j=0; j<threadsCount; j++)
	{
		results[j] = results_tmp[j].get();
	}

	// Show Results of all Workers.
	for (j=0; j<threadsCount; j++)
	{
		ii = results[j].indices.size();
		
		std::cout << "Worker #" << (j+1) << " Result: ";
		std::cout << "[" << results[j].error << "] ";
		std::cout << "[" << results[j].duplicateIsFound << "] ";
		std::cout << "[" << results[j].elementMin << "] ";
		std::cout << "[" << results[j].elementMax << "] ";
		std::cout << "[" << ii << "] ";
		if (verbose)
		{
			std::cout << "<< ";
			for (i=0; i<ii; i++)
			{
				std::cout << results[j].indices[i] << " ";
			}
			std::cout << ">>";
		}
		std::cout << "\r\n";
	}

	// Check Error in Results.
	error = false;
	for (j=0; j<threadsCount; j++)
	{
		if (results[j].error == true)
		{
			error = true;
			break;
		}
	}
	if (error)
	{
		return appExitCodeError;
	}
	
	// Concatenate all the Results.
	
	// 1. Initialize Statistics.
	for (i=0; i<=UINT16_MAX; i++)
	{
		total.elementOccurrence[i] = 0;
	}
	total.elementMax = results[0].elementMax;
	total.elementMin = results[0].elementMin;
	
	// 2. Concatenate Results.
	for (j=0; j<threadsCount; j++)
	{
		// Duplicates List.
		if (results[j].duplicateIsFound == true)
		{
			kMax = results[j].indices.size() - 1;
			for (k=0; k<=kMax; k++)
			{
				index = results[j].indices[k];
				total.indices.push_back(index);
			}
		}
		
		// Statistics.
		if (results[j].elementMax > total.elementMax)
		{
			total.elementMax = results[j].elementMax;
		}
		if (results[j].elementMin < total.elementMin)
		{
			total.elementMin = results[j].elementMin;
		}
		for (i=0; i<=UINT16_MAX; i++)
		{
			total.elementOccurrence[i] += results[j].elementOccurrence[i];
		}
	}

	// Show all Results.
	std::cout << "Overall Results:" << "\r\n";
	std::cout << "Elements Count = [" << dataCount << "].\r\n";
	std::cout << "Minimum Element = [" << total.elementMin << "].\r\n";
	std::cout << "Maximum Element = [" << total.elementMax << "].\r\n";
	std::cout << "Duplicates Count = [" << total.indices.size() << "].\r\n";
	if (verbose)
	{
		std::cout << "Duplicate Elements:\r\n";
		std::cout << "<< ";
		jMax = total.indices.size() - 1;
		for (j=0; j<=jMax; j++)
		{
			std::cout << total.indices[j] << " ";
		}
		std::cout << ">>";
	}
	std::cout << "\r\n";
	
	// Write Occurrence Statistics to File.
	fileOut.open(outputFileName);
	for (i=0; i<=UINT16_MAX; i++)
	{
		fileOut << i << " " << total.elementOccurrence[i] << "\r\n";
	}
	fileOut.close();

	return appExitCodeSuccess;
}

void generateData(Data &data, Uint64 size)
{
	Uint64 i;
	Uint16 element;

	std::random_device rndDevice;
	std::mt19937 rndEngine(rndDevice());
	std::uniform_int_distribution<Uint16> rndDistribution;

	for (i = 0; i < size; i++)
	{
		element = rndDistribution(rndEngine);
		data.push_back(element);
	}
}

void worker(Task task)
{
	int cpuNumPlanned;
	Data *data;
	Uint16 elementCur;
	Uint16 elementMax;
	Uint16 elementMin;
	Uint16 elementPrev;
	Uint64 i;
	char id;
	Uint64 iFirst;
	Uint64 iLast;
	Result result;
	bool segmentIsNotFirst;
	std::stringstream streamBuffer;
	
#ifdef __linux__
	int affRes;
	cpu_set_t cpuset;
	pthread_t thread; // Self Thread.
#endif

	// Read Task.
	cpuNumPlanned = task.cpuNumPlanned;
	id = task.id;
	iFirst = task.indexFirst; // Is at least 0.
	iLast = task.indexLast;
	data = task.data;
	segmentIsNotFirst = task.segmentIsNotFirst;
	result.error = false;
	result.duplicateIsFound = false;
	
	// Set Thread Affinity. For Linux only.
#ifdef __linux__
	thread = pthread_self();
	CPU_ZERO(&cpuset);
	// Mark only the planned CPU as available.
	CPU_SET(cpuNumPlanned, &cpuset);
	// Mark all CPUs as available.
	/*
	for (j = 0; j < cpusCount; j++)
	{
		CPU_SET(j, &cpuset);
	}
	*/
	affRes = pthread_setaffinity_np(thread, sizeof(cpu_set_t), &cpuset);
	if (affRes != 0)
	{
		streamBuffer << "Error. CPU Affinity Set has failed!" << "\r\n";
	}
#endif
	
	// Report.
	streamBuffer << "Worker has started. ";
	streamBuffer << "thread_id=[" << std::this_thread::get_id() << "] ";
	streamBuffer << "CPU.Plan=[" << cpuNumPlanned << "]";
#ifdef __linux__
	streamBuffer << " CPU.Real=[" << sched_getcpu() << "]";
#endif
	streamBuffer << ".";
	streamBuffer << "\r\n";
	streamBuffer << "id=[" << int(id) << "] ";
	streamBuffer << "data=[" << data << "].";
	streamBuffer << "sinf=[" << segmentIsNotFirst << "] ";
	streamBuffer << "iF=[" << iFirst << "] ";
	streamBuffer << "iL=[" << iLast << "] ";
	streamBuffer << "\r\n";
	std::cout << streamBuffer.str();

	// Check Task's Parameters.
	if (iFirst >= iLast)
	{
		// Error.
		std::cout << "Task Error. Thread Termination.\r\n";
		result.error = true;

		pipes[id-1].set_value(result);

		return;
	}
	
	// Initialize Element Occurrence Statistics.
	for (i=0; i<=UINT16_MAX; i++)
	{
		result.elementOccurrence[i] = 0;
	}

	// To ensure that Border-Elements are also checked, we compare the first
	// Data Element with the last Element of the previous Data Segment.
	// Previous Data Segment exists only if we are using non-first Data Segment.
	if (segmentIsNotFirst)
	{
		if ((*data)[iFirst] == (*data)[iFirst-1])
		{
			// Luck.
			result.duplicateIsFound = true;
			result.indices.push_back(iFirst-1);
		}
	}

	// Preparations.
	iFirst++; // Is > 0.
	i = iFirst;
	elementPrev = (*data)[i-1];
	elementMax = elementPrev;
	elementMin = elementPrev;
	result.elementOccurrence[elementPrev]++;
	
	// Search-Loop.
	while ((i > 0) and (i <= iLast))
	{
		elementCur = (*data)[i];
		
		// Gather Statistics.
		if (elementCur > elementMax)
		{
			elementMax = elementCur;
		}
		if (elementCur < elementMin)
		{
			elementMin = elementCur;
		}
		
		// Occurrence Statistics.
		result.elementOccurrence[elementCur]++;
		
		// Check Duplicate.
		if (elementPrev == elementCur)
		{
			// Luck.
			result.duplicateIsFound = true;
			result.indices.push_back(i-1);
		}

		// Next.
		elementPrev = elementCur;
		i++; // Protection against Overflow is done in 'i>0' Check.
	}
	
	result.elementMax = elementMax;
	result.elementMin = elementMin;

	// Send Result to main Thread.
	pipes[id-1].set_value(result);

	return;
}
