================================================================================

This is an Example of an extendable Application.

================================================================================

The whole Demonstration consists of three Parts or Steps.

1. A separate Library holding an Interface, compiled as a DLL.
2. A Program which has a standard Implementation of the mentioned Interface.
3. An external Library implementing the Interface, compiled as a DLL.

================================================================================

How to make all this work?

To make all this work well, the Interface must be made as a separate Solution 
and compiled into a DLL. The Interface is the first Thing to write.

After the Interface is compiled into a separate DLL, the main Program must be 
written as a separate Solution and compiled. To use the external DLL Library 
in Visual Studio 2019 it is enough to select 'Dependencies' in the Solution 
Explorer, then click 'Add Reference...' and select the external DLL File. Do 
not forget to add the 'using' Statement.

Third Step is to write and compile the Extension. As in the main Program, the 
Extension must be written as a separate Solution, must reference an external 
DLL holding an Interface, and must be built into a separate DLL.

After all these Steps are done, the DLL File built in the third Step must be 
copied to the Second Solution to the 'plugin' Folder and renamed as 
'storage.dll'.

Now that you have loaded the Plug-In DLL dynamically in the C# Code in the main 
Program, you are able to cast the loaded Object to the Interface Type and use 
it.

================================================================================
