1. Install the 'Graphviz' Tool.
-------------------------------

2. Build the Program.
---------------------
	go build

3. Run the Program for the Profilers to create some Data.
---------------------------------------------------------
	1.exe -profiler_cpu=cpu.prof -profiler_mem=mem.prof

4. Analyze the Profiling Data.
------------------------------
	go tool pprof cpu.prof
	>help
	>top
	>web
	>svg
