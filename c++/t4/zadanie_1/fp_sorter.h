// fp_sorter.h.

#ifndef FP_SORTER_H
#define FP_SORTER_H

#include <vector>

// Maximum R.A.M. Usage (Bytes).
const unsigned int ram_usage_max = 100 * 1000 * 1000; // 100 MB.
//const unsigned int ram_usage_max = 100 * 1000; // 100 KB.

constexpr unsigned int records_count_in_ram =
	ram_usage_max / sizeof(double);

const char *numbers_separator = "\r\n";

std::string file_segment_prefix = "seg_";
std::string file_segment_postfix = ".dat";
std::string folder_sep = "/";
std::string folder_tmp = "tmp";

// Error Codes.
const char error_none = 0;
const char error_fopen = 1;
const char error_fclose = 2;
const char error_system = 3;
const char error_save = 4;
const char error_fread = 5;
const char error_segdata = 6;
const char error_fseek = 7;
const char error_fwrite = 8;
const char error_index = 9;

// Error Messages.
const char msg_error_fclose[] = "Unable to open output File";
const char msg_error_fopen[] = "Unable to open output File";
const char msg_error_fread[] = "File Read Error";
const char msg_error_fseek[] = "File Seek Error";
const char msg_error_fwrite[] = "File Write Error";
const char msg_error_index[] = "Index is out of Bounds";
const char msg_error_save[] = "Save Error";
const char msg_error_segdata[] = "Segment Data File Corruption";
const char msg_error_system[] = "System Call Failure";

// Types.
struct Segment
{
    FILE *file;
    std::string file_name;
    unsigned int items_count;
    double current_item;
    unsigned int current_item_index;
    bool hasUnprocessedItems;
};

// Functions.
bool data_is_available
(
    std::vector<Segment> &segments
);

int fp_sort
(
    const char *src_path,
    const char *dst_path
);

long int get_file_size
(
    std::string file_name
);

double get_segment_item
(
    std::vector<Segment> &segments,
    unsigned int segment_idx,
    unsigned int item_idx
);

int main
(
    int argc,
    char* argv[]
);

int segment_save
(
    std::vector<double> &vec,
    std::string &file_name,
    unsigned int count
);

void shift_segment_item
(
    std::vector<Segment> &segments,
    unsigned int segment_idx
);

#endif // FP_SORTER_H
