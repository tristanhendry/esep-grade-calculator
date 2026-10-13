# ESEP Grade Calculator

A Go library that calculates a final letter grade from assignment,
exam, and essay scores.

## Requirements

- Go 1.24.4 or newer

## Setup

Clone your fork and open the project:

    git clone https://github.com/tristanhendry/esep-grade-calculator.git
    cd esep-grade-calculator/grade-calculator

## How grades are calculated

Each category's average contributes to the final score:

- Assignments: 50%
- Exams: 35%
- Essays: 15%

The calculator uses these letter-grade thresholds:

- A: 90 or higher
- B: 80–89
- C: 70–79
- D: 60–69
- F: below 60

A C or better is passing.

Create a calculator with NewGradeCalculator(), add scores using
AddGrade(name, score, gradeType), and call GetFinalGrade().

Supported grade types are Assignment, Exam, and Essay.

## Run tests

From the grade-calculator folder:

    go test -v ./...

## Measure test coverage

    go test -coverprofile=coverage.txt ./...
    go tool cover -func=coverage.txt

To view covered and uncovered code in a browser:

    go tool cover -html=coverage.txt

## Current limitations

Provide at least one score in each category; an empty category causes
division by zero. Scores are not validated. Category averages use
integer division, and the final weighted score is truncated to an integer.
