------------------------
Block C.R.C. Check Tool.
Description.
------------------------

This is a Test Tool to create and verify CRC32 CheckSums of a File.

In 'CheckSum File Creation' Mode, this Tools reads the Data File by Blocks of 
the specified Size and then writes CheckSums to the CheckSums File. For each 
Block of the Data File, two CheckSums are calculated and written: 

1. CheckSum of the current Block;
2. CheckSum of the "sliding" ("running") Block.

Sliding Block is a Concatenation of all previously read Blocks (before the 
current Block) into a big single Block. The first sliding CheckSum equals to 
the CheckSum of the first current Block. The last sliding CheckSum equals to 
the CheckSum of the whole Data File. Sliding Block CheckSums help in inspecting 
Error Types, they may be very useful in discovering the Source of Errors. Not 
in 100% Cases, but very often, the Place and Reason of Corruption may be 
learned by the Inspection of Error List.

In 'CheckSum Verification' Mode, this Tools reads the CheckSum File (as well 
as the Data File) and verifies the Integrity of both Data File and CheckSums 
File.

Unfortunately, this Tool can only discover Errors and the Place where they 
occur. This Tool can not recover corrupted Data, while this Operation is very 
Time consuming due to the Nature of Hash Sums.

If the Path to CheckSums File is set to empty, the Tool then copies it from the 
Data File Path Parameter, with the 'bcrc' Extension appended.

To get a full List of Command Line Parameters, run the Tool with '-h' Argument.
------------------------
