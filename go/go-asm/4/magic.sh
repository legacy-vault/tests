#!/bin/bash

# Build separate Source Files into separate Binary Executables.
go build a.go
go build b.go

# Create Dumps of each Executable.
go tool objdump a > a.dump
go tool objdump b > b.dump
