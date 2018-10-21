// rw_mutex.h.

#ifndef RW_MUTEX_H
#define RW_MUTEX_H

#include <atomic>
#include <condition_variable>
#include <mutex>

namespace mcarcher
{

class RWMutex
{

public:
	
	RWMutex();
	~RWMutex() = default;
	
	// Information Inquiry Methods.
	bool IsRLocked();
	bool IsRWLocked();
	unsigned int ReadersCount();
	
	// Methods for Readers.
	void RLock();
	void RUnlock();
	
	// Methods for Writers.
	void RWLock();
	void RWUnlock();
	
private:
	
	// Set of atomic Variables (Flags and Counters).
	std::atomic<bool> is_rlocked;
	std::atomic<bool> is_rwlocked;
	std::atomic<bool> new_readers_are_allowed;
	std::atomic<unsigned int> readers_count;
	
	// Set of Variables to synchronize the Usage of 'readers_count':
	//	1. Readers notifying the Writer about its changes,
	// 	2. Writer waiting for it to become Zero.
	std::condition_variable readers_cv;
	std::mutex readers_cv_mutex; // -> readers_cv.
	
	// Set of Variables to synchronize the Usage of 'new_readers_are_allowed':
	//	1. Writer notifying all Readers about it set to TRUE.
	// 	2. Readers waiting for it to become TRUE,
	std::condition_variable writer_cv;
	std::mutex writer_cv_mutex; // -> writer_cv.
	
	// A Mutex to allow only a single Writer.
	std::mutex rw_mutex;
};

}

#endif // RW_MUTEX_H
