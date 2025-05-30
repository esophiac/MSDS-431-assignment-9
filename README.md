# Assignment Week 9 - Create a Go Package
For this assignment, we created a package to read a slice of integers and/or floats and compute a trimmed mean. Then, we create a repository to use the repository that demonstrates the utility of the first repository for floats and ints. This test was done using samples of at least 100 integers and 100 floats with symmetric trimming results and with 0.05 observations taken from both ends of the distribution. 

## Background
This assignment uses the Mersenne Twister from the third-party package at [https://github.com/seehuhn/mt19937](https://github.com/seehuhn/mt19937).

## Roles of Programs and Data
This program was completed in Go, but there is also a demonstration of the trimmed mean in R. The demonstration in R shows that the bootstrap sampled trimmed mean is the same thing as the trimmed mean of the original sample. Therefore, the bootstrap sampled mean in Go is the same thing as the original sample mean.

### trimmed Package
The package that was used to compute the trimmed mean can be found at [https://github.com/esophiac/trimmed](https://github.com/esophiac/trimmed).

### Programs
These are the programs in the repository.
- go.mod: defines the module's properties
- go.sum: record of the libraries the project depends on
- main_test.go: tests and benchmarks the fuctions in the main.go file
- main.go: demonstrating the trimmed package
- README.md: the readme file for the repository
- trimmed_mean.R : demonstrating how the trimmed mean works in R.

## Application
An executable for this project was created using Windows. To create your own executable, run go build in the same directory as the go program. For more information, see the Gopher documentation on creating an executable [here](https://go.dev/doc/tutorial/compile-install).

## Use of AI
AI was not used for this assignment.