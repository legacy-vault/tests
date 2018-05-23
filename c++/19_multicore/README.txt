This is an Example of simple multi-core Data Cruncher.

It prepares the random Data (using CPU's single Core),
then inspects the Data divided into Segments using multiple Threads. 
Each Thread Worker searches for duplicate Neighbour Elements in his Data 
Segment and gathers some simple Statistics (maximum & minimum Elements).

In Linux-based O.S. these Threads are separated into different CPU Cores using 
the 'CPU Affinity' Feature of the Operating System. On Operating Systems other 
than Linux-based, the Behaviour depends on the C++ Language Compiler.

Author: McArcher.
Date: 2018-05-23.
