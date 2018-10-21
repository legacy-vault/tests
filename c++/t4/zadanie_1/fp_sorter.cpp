// fp_sorter.cpp.

#include <algorithm>
#include <cstdio>
#include <fstream>
#include <iostream>
#include <string>
#include <sys/stat.h>
#include <vector>

#include "fp_sorter.h"

int main
(
    int argc,
    char* argv[]
)
{
	char *file_dst_path;
	char *file_src_path;
	int result;

	if (argc != 3)
	{
		std::cout << "Usage:" << std::endl <<
			"\t" << "fp_sorter <source_file> <destination_file>" << std::endl;

		return 0;
	}

	file_src_path = argv[1];
	file_dst_path = argv[2];

	result = fp_sort(file_src_path, file_dst_path);

	if (result != error_none)
	{
		std::cout << "Error! Error Code = [" << result << "]." << std::endl;
	}

	return result;
}

// Checks whether there is any Data (in any Segment) available to be sorted.
bool data_is_available
(
    std::vector<Segment> &segments
)
{
    bool have_data = false;
    unsigned int i;
    unsigned int seg_count = segments.size();
    bool segment_has_data;

    for (i = 0; i < seg_count; i++)
    {
        segment_has_data = segments[i].hasUnprocessedItems;
        have_data = have_data || segment_has_data;
    }

    return have_data;
}

int fp_sort
(
    const char *src_path,
    const char *dst_path
)
{
	// Buffer for String Representation of 'double' Value.
	// -0.1234567890123456e+308 and terminating "\0". 24+1 Symbols.
	char buffer[25]; // Stores maximum 16 Digits after Point.

	std::string buffer_in;
	char ch;
	const char *cmd;
	std::string cmd_str;
	std::vector<double> data_segment;
	unsigned int data_segments_count;
	unsigned int data_segment_idx; // Iterator.
	std::string file_name;
	std::ofstream file_dst;
	std::ifstream file_src;
	long int file_size;
	unsigned int i;
	double item;
	double minimum_item;
	double record;
	unsigned int record_num;
	int result_fclose;
	int result_fopen;
	int result_save;
	int result_system;
	Segment seg;
	std::vector<Segment> segments;
	bool segment_has_data;
	unsigned int segment_with_least_item_idx;
	bool shouldContinue;
	
	// Report.
	std::cout << "Preparing Segments..." << std::flush;

	// Create temporary Folder.
	cmd_str = std::string("mkdir \"") + folder_tmp + std::string("\"");
	cmd = cmd_str.c_str();
	result_system = system(cmd);
	if (result_system != 0)
	{
		std::cout << msg_error_system << std::endl;

		return error_system;
	}

	// Open Source File.
	file_src.open(src_path, std::ios_base::in);
	result_fopen = file_src.is_open();
	if (result_fopen != 1)
	{
		std::cout << msg_error_fopen << std::endl;

		return error_fopen;
	}

	data_segment_idx = 0;
	file_name = folder_tmp + folder_sep + file_segment_prefix +
        std::to_string(data_segment_idx) + file_segment_postfix;
	record_num = 0;

	// Read Records from File.
	while (std::getline(file_src, buffer_in))
	{
        ch = buffer_in.back();
        if (ch == '\r')
        {
            buffer_in.pop_back();
        }
        record = std::stod(buffer_in);
        data_segment.push_back(record);

        // Next Record.
        record_num++;

        // R.A.M. is full?
        if (record_num == records_count_in_ram)
        {
            // Sort Data.
            std::sort(data_segment.begin(), data_segment.end());

            // Save to a temporary File.
            result_save = segment_save(
                data_segment,
                file_name,
                data_segment.size());

            if (result_save != 0)
            {
                std::cout << msg_error_save << std::endl;

                return error_save;
            }

            // Reset Records Counter.
            record_num = 0;
            data_segment_idx++;
            file_name = folder_tmp + folder_sep + file_segment_prefix +
                std::to_string(data_segment_idx) + file_segment_postfix;
            data_segment.clear();
        }
	}

	if (file_src.fail())
    {
        if (file_src.eof())
        {
            if (record_num != records_count_in_ram)
            {
                // We have got a partially filled Segment,
                // which has not been saved to File.
                data_segment_idx++;
                data_segments_count = data_segment_idx;

                // Sort Data.
                std::sort(data_segment.begin(), data_segment.end());

                // Save to a temporary File.
                result_save = segment_save(
                    data_segment,
                    file_name,
                    record_num);

                if (result_save != 0)
                {
                    std::cout << msg_error_save << std::endl;

                    return error_save;
                }
            }
        }
        else
        {
            // Not EOF Error.
            std::cout << msg_error_fread << std::endl;

            return error_fread;
        }
    }

	// Close File.
	file_src.close();
	result_fclose = file_src.is_open();
	if (result_fclose != 0)
	{
		std::cout << msg_error_fclose << std::endl;

		return error_fclose;
	}
	
	// Report.
	std::cout << "Done." << std::endl;

	// Concatenate Data from Segments...
	
	// Report.
	std::cout << "Processing " << data_segments_count << 
		" Segments..." << std::flush;

	// Initialize Data Segments.
	data_segment_idx = 0;
	while (data_segment_idx < data_segments_count)
	{
        // Set Segment's Parameters.

        // 1. File Name & Descriptor.
        seg.file_name = folder_tmp + folder_sep + file_segment_prefix +
                std::to_string(data_segment_idx) + file_segment_postfix;

        // Here we use good old C-Style Way to open Files.
        // C++ Streams can not be used as they can not be copied!
        // Good old and simple C is often much better than the C++ Monster.
        seg.file = fopen(seg.file_name.c_str(), "r");
        if (seg.file == NULL)
        {
            std::cout << msg_error_fopen << std::endl;

            return error_fopen;
        }

        // 2. Items Count.
        file_size = get_file_size(seg.file_name);
        if (file_size % sizeof(double) != 0)
        {
            std::cout << msg_error_segdata << std::endl;

            return error_segdata;
        }
        seg.items_count = file_size / sizeof(double);
        seg.current_item_index = 0;
        seg.current_item = 0;
        if (seg.items_count != 0)
        {
            seg.hasUnprocessedItems = true;
        }
        else
        {
            seg.hasUnprocessedItems = false;

            std::cout << msg_error_segdata << std::endl;

            return error_segdata;
        }

        // Save Segment's Parameters.
        segments.push_back(seg);

        // 3. Current Item.
        segments[data_segment_idx].current_item = get_segment_item(
            segments,
            data_segment_idx,
            segments[data_segment_idx].current_item_index);


        // Next Segment.
        data_segment_idx++;
	}

	// Open Output File.
	file_dst.open(dst_path, std::ios_base::out | std::ios_base::trunc);
	result_fopen = file_dst.is_open();
	if (result_fopen != 1)
	{
		std::cout << msg_error_fopen << std::endl;

		return error_fopen;
	}

	// Process Data Segments.
	shouldContinue = data_is_available(segments);
	while (shouldContinue)
	{
        // Find which Segment has the least Element.

        // 1. Find first Segment with Data available.
        i = 0;
        segment_with_least_item_idx = 0;
        segment_has_data = segments[i].hasUnprocessedItems;
        while (segment_has_data == false)
        {
            i++;    // We know that at least one of the Segments has Data,
                    // so, we do not check the Range of 'i' Index.
            segment_has_data = segments[i].hasUnprocessedItems;
        }
        segment_with_least_item_idx = i;
        minimum_item = segments[segment_with_least_item_idx].current_item;

        // 2. Search rest Segments for minimum Element.
        for (i = 0; i < data_segments_count; i++)
        {
            segment_has_data = segments[i].hasUnprocessedItems;
            if (segment_has_data)
            {
                item = segments[i].current_item;
                if (item < minimum_item)
                {
                    minimum_item = item;
                    segment_with_least_item_idx = i;
                }
            }
        }

        // Save minimum Element to Output File.
        sprintf(buffer, "%.16e", minimum_item);
        file_dst << buffer << numbers_separator;

        // Shift Element in Segment used.
        shift_segment_item(segments, segment_with_least_item_idx);

        // Next Loop.
        shouldContinue = data_is_available(segments);
	}

	// Close Output File.
	file_dst.close();
	result_fclose = file_dst.is_open();
	if (result_fclose != 0)
	{
		std::cout << msg_error_fclose << std::endl;

		return error_fclose;
	}

	// Finalize Data Segments.
	data_segment_idx = 0;
	while (data_segment_idx < data_segments_count)
	{
        seg = segments[data_segment_idx];

        // Close File.
        result_fclose = fclose(seg.file);
        if (result_fclose != 0)
        {
            std::cout << msg_error_fclose << std::endl;

            return error_fclose;
        }

        // Next Segment.
        data_segment_idx++;
	}

	// Delete temporary Folder.
	cmd_str = std::string("rm -r \".") + 
		folder_sep + folder_tmp + 
		std::string("\"");
	
	cmd = cmd_str.c_str();
	result_system = system(cmd);
	if (result_system != 0)
	{
		std::cout << msg_error_system << std::endl;

		return error_system;
	}
	
	// Report.
	std::cout << "Done." << std::endl;

	return error_none;
}

// Returns File's Size.
long int get_file_size
(
    std::string file_name
)
{
    struct stat stat_buf;
    int rc;

    rc = stat(file_name.c_str(), &stat_buf);
    if (rc == 0)
    {
        return stat_buf.st_size;
    }
    else
    {
        return -1;
    }
}

// Returns an Item from Segment.
// On Error: stops the Program.
double get_segment_item
(
    std::vector<Segment> &segments,
    unsigned int segment_idx,
    unsigned int item_idx
)
{
    FILE *file;
    double item;
    unsigned int items_count;
    long int offset;
    int result;
    unsigned int segments_count;

    // Check Input Parameters...

    // 1. Segment Index.
    segments_count = segments.size();
    if (segment_idx >= segments_count)
    {
        // Error Report.
        std::cout << "Segment Index: " << segment_idx << std::endl;
        std::cout << msg_error_index << std::endl;

        exit(error_index);
    }

    // 2. Item Index.
    items_count = segments[segment_idx].items_count;
    if (item_idx >= items_count)
    {
        // Error Report.
        std::cout << "Item Index: " << item_idx << std::endl;
        std::cout << msg_error_index << std::endl;

        exit(error_index);
    }

    // Get File.
    file = segments[segment_idx].file;

    // Seek Record in File.
    offset = item_idx * sizeof(double);
    result = fseek(file, offset, SEEK_SET);
    if (result != 0)
    {
        // Error Report.
        std::cout << "Segment Index: " << segment_idx << std::endl;
        std::cout << "Item Index: " << item_idx << std::endl;
        std::cout << msg_error_fseek << std::endl;

        exit(error_fseek);
    }

    // Read Record from File.
    item = 0;
    result = fread(&item, sizeof(double), 1, file);
    if (result != 1)
    {
        // Error Report.
        std::cout << "Segment Index: " << segment_idx << std::endl;
        std::cout << "Item Index: " << item_idx << std::endl;
        std::cout << msg_error_fread << std::endl;

        exit(error_fread);
    }

    return item;
}

// Saves a Segment of Data to a binary File.
// The 'count' Parameter tells how many Elements of the Array must be saved.
int segment_save
(
    std::vector<double> &vec,
    std::string &file_name,
    unsigned int count
)
{
    std::ofstream file;
    int i;
    int i_max;
    std::ios_base::openmode mode;
    double record;
    int result;

    // Open File.
    mode = std::ios_base::out | std::ios_base::trunc | std::ios_base::binary;
    file.open(file_name, mode);
	result = file.is_open();
	if (result != 1)
	{
		std::cout << msg_error_fopen << std::endl;

		return error_fopen;
	}

	// Write Records to File.
	i_max = count;
	for (i = 0; i < i_max; i++)
	{
        record = vec.at(i);
        file.write( reinterpret_cast<char*>( &record ), sizeof(record) );
	}

	// Close File.
	file.close();
	result = file.is_open();
	if (result != 0)
	{
		std::cout << msg_error_fclose << std::endl;

		return error_fclose;
	}

	return error_none;
}

// Moves "Cursor" to next Item. If next item does not exist, raises a Flag.
void shift_segment_item
(
    std::vector<Segment> &segments,
    unsigned int segment_idx
)
{
    unsigned int current_item_index;
    bool hasUnprocessedItems;
    unsigned int items_count;
    unsigned int segments_count;

    // Check Input Parameters...

    // 1. Segment Index.
    segments_count = segments.size();
    if (segment_idx >= segments_count)
    {
        // Error Report.
        std::cout << "Segment Index: " << segment_idx << std::endl;
        std::cout << msg_error_index << std::endl;

        exit(error_index);
    }

    // No Items to Process.
    hasUnprocessedItems = segments[segment_idx].hasUnprocessedItems;
    if (hasUnprocessedItems == false)
    {
        return;
    }

    // Increase Current Item Counter.
    items_count = segments[segment_idx].items_count;
    segments[segment_idx].current_item_index++;
    current_item_index = segments[segment_idx].current_item_index;

    // Last Item was processed.
    if (current_item_index >= items_count)
    {
        segments[segment_idx].hasUnprocessedItems = false;
        segments[segment_idx].current_item = 0;

        return;
    }

    // Update current Item.
    segments[segment_idx].current_item = get_segment_item(
        segments,
        segment_idx,
        current_item_index);
}
