// fp_generator.h.

#ifndef FP_GENERATOR_H
#define FP_GENERATOR_H

// Maximum File Size (Bytes).
const unsigned int file_size_max = 1 * 1000 * 1000 * 1000; // 1 GB.
//const unsigned int file_size_max = 1 * 1000 * 1000; // 1 MB.

const char *numbers_separator = "\r\n";
const unsigned int numbers_separator_length = 2;

// File Path.
const char file_out_path[] = "data_src.txt";

// Error Codes.
const char error_fclose = 2;
const char error_fopen = 1;
const char error_none = 0;

// Error Messages.
const char msg_error_fclose[] = "Unable to open output File";
const char msg_error_fopen[] = "Unable to open output File";

// Functions.
int fp_generate(const char *dst_path, unsigned int file_size_limit);
int main();

#endif // FP_GENERATOR_H
