# ESEP Grade Calculator

This library is a simple utility to calculate your grades.
There are three categories that are used to compute a final letter grade for each student:

- **Assignments** (50%)
- **Exams** (35%)
- **Essays** (15%)

Final grades are mapped using:
- **A** ≥ 90, **B** ≥ 80, **C** ≥ 70, **D** ≥ 60, **F** for all others

## Tech Stack
-**Language:** Go (Golang)
- **Main Module** 'esep/grade-calculator'
- **Go Version** '1.24.4'

## Install Go (1.24)
Verify using:
'''bash go version
To install: https://go.dev/doc/install
Download the current Go version for the operating system 
Verify using the command above

## Go resources
Learn more about Go: https://go.dev/learn/

## Project files
- **go.mod** Module and Go version
- **grade_calculator.go** Main code
- **grade_calculator_test.go** Unit tests

## Run unit tests:
```
cd grade_calculator/
go test ./...
```

## Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```

## Contributing
Fork the repository on your account and create a branch per change
Open a PR with a clear description

