# Problem
Build a program that will rate the Ecobee thermostat owners by the quality of insulation of their home.  Each user's R-value is to be rated relative to the others in the same region (city/province/county) they live.

# Solution
To analyze this data produced by Ecobee thermostats the program has been written in Golang, making it fast as well cross-platform compatible out of the box.  Only Go's standard library is used so no 3rd party package installs or cross-platform compiling is required.  With a modern version of Go installed the program can be seamlessly compiled and run on the target machine.

## Run Instructions
With Go installed and configured on your machine [(instructions here)](https://golang.org/doc/install), download and extract the program and cd into it.  Run it with the following commands.  
Linux:
```
go run . < input.txt
```
Windows:
```
go build
./thermostat.exe < input.txt
```
`input.txt` can be replaced with any file containing valid input data.

## Design
The problem notes following these principles:

__YAGNI__ - You Aren't Gonna Need It - states to not implement presumed future capability now because it likely won't be needed.  This is because implementing code that is dead on arrival unnecessarily adds complexity to a program.  An example of breaking the YAGNI principle in this program would be functionality to convert an Owner object back to the input string representation.  While a complete, general-purpose API would provide such functionality it's not needed here.

__DRY__ Don't Repeat Yourself.  Functions exist in programming for this exact reason.  Laying out code in well designed modules following the principles *low coupling* and *high cohesion* allow for programs to easily avoid messes like duplicated code.

## Testing
Unit tests have been implemented using Go's `testing` package.
Run the unit tests with the command:
`go test`

Unit tests have total coverage over `owner.go` and `owner_collection.go`.  This covers everything except the main module code which interfaces between stdin/stdout and these modules.  Testing main.go is best left to functional and system level tests.

Manual testing may also be performed with the provided input files.  `input.txt` contains valid input data while `input_invalid.txt` will cause an error.  With knowledge of the source code input data can be altered to test the various edge cases.

## Performance
### Time
In terms of time complexity, the program loads the dataset in `O(N^2)` time.  This is because the dataset is sorted after loading the input data.  There is a performance trade-off here with running queries, which benefit from the sorted dataset and each run in `O(N)` time.

### Space
The program is `O(N)` in space complexity.  Though each entry is loaded into a struct the use of pointers is leveredged by storing the entire set as a slice (array interface) of struct pointers.  This means when sorting only the memory addresses of each struct are copied avoiding moving the entire struct.  Though not expressed in Big-O notation the memory savings here are *considerable*.
