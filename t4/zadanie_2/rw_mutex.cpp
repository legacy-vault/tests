// rw_mutex.cpp.

#include <atomic>
#include <condition_variable>
#include <iostream>
#include <mutex>

#include "rw_mutex.h"

using namespace mcarcher;

RWMutex::RWMutex()
{
	is_rlocked = false;
	is_rwlocked = false;
	new_readers_are_allowed = true;
	readers_count = 0;
}

bool RWMutex::IsRLocked()
{
	if (is_rlocked == true)
	{
		return true;
	}
	else
	{
		return false;
	}
}

bool RWMutex::IsRWLocked()
{
	return is_rwlocked;
}

unsigned int RWMutex::ReadersCount()
{
	unsigned int rc;
	
	rc = readers_count;
	
	return rc;
}

void RWMutex::RLock()
{
	//========================================================================//
	//
	// Notes:
	//
	// In C++, Condition Variable (C.V.) is acting not as Channels in Go.
	// If the Waiter starts 'wait' Method after the Signal has been emitted, it
	// will miss the Signal as the C.V. Notifications are not saved for 
	// Future Use. In this Aspect, the 'Channels' Mechanism of Go Language is
	// much easier and safer. To get the correct Behaviour, we guard the Code
	// Section with additional Mutex with 'Unique Lock' Wrapper and an 
	// additional Variable which is checked. Here it is the
	// 'new_readers_are_allowed' Variable.
	//
	//========================================================================//
	
	std::unique_lock<std::mutex> ul(writer_cv_mutex);
	
	// Wait for Readers to be allowed to read.
	while (new_readers_are_allowed != true)
	{
		writer_cv.wait(ul);
	}
	
	readers_count++;
	is_rlocked = true;
}

void RWMutex::RUnlock()
{
	std::unique_lock<std::mutex> ul(readers_cv_mutex);
	
	if (readers_count == 0)
	{
		std::cout << "RUnlock has failed. Already unlocked!" << std::endl;
		exit(1);
	}
	
	readers_count--;
	
	// Final Reader has finished?
	if (readers_count == 0)
	{
		is_rlocked = false;
	}
	
	// Notify waiting Writer about us leaving.
	readers_cv.notify_all();
}

void RWMutex::RWLock()
{
	//========================================================================//
	//
	// Notes:
	//
	// In C++, Condition Variable (C.V.) is acting not as Channels in Go.
	// If the Waiter starts 'wait' Method after the Signal has been emitted, it
	// will miss the Signal as the C.V. Notifications are not saved for 
	// Future Use. In this Aspect, the 'Channels' Mechanism of Go Language is
	// much easier and safer. To get the correct Behaviour, we guard the Code
	// Section with additional Mutex with 'Unique Lock' Wrapper and an 
	// additional Variable which is checked. Here it is the 'readers_count' 
	// Variable.
	//
	//========================================================================//
	
	// Block the Mutex for the solo RW-Job.
	rw_mutex.lock();
	is_rwlocked = true;
	
	std::unique_lock<std::mutex> ul(readers_cv_mutex);
	
	// Forbid new Readers.
	new_readers_are_allowed = false;
	
	// Wait for the final Reader to complete its Job.
	while (readers_count != 0)
	{
		readers_cv.wait(ul);
	}
}

void RWMutex::RWUnlock()
{
	std::unique_lock<std::mutex> ul(writer_cv_mutex);
	
	if (is_rwlocked == false)
	{
		std::cout << "RWUnlock has failed. Already unlocked!" << std::endl;
		exit(1);
	}
	
	// Permit new Readers.
	new_readers_are_allowed = true;
	
	// Notify waiting Readers about us leaving.
	writer_cv.notify_all();
	
	// Unlock the RW Mutex. 
	// The Order is important as Unlocking the Mutex will unblock another 
	// waiting Writer who may see an old 'is_rwlocked' Value.
	is_rwlocked = false;
	rw_mutex.unlock();
}
