package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 50, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 50, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	calculator := NewGradeCalculator()

	calculator.AddGrade("assignment", 70, Assignment)
	calculator.AddGrade("exam", 70, Exam)
	calculator.AddGrade("essay", 70, Essay)

	actual := calculator.GetFinalGrade()
	if actual != "C" {
		t.Errorf("Expected C; got %s", actual)
	}
}

func TestGetGradeD(t *testing.T) {
	calculator := NewGradeCalculator()

	calculator.AddGrade("assignment", 60, Assignment)
	calculator.AddGrade("exam", 60, Exam)
	calculator.AddGrade("essay", 60, Essay)

	actual := calculator.GetFinalGrade()
	if actual != "D" {
		t.Errorf("Expected D; got %s", actual)
	}
}

func TestGradeTypeString(t *testing.T) {
	tests := []struct {
		gradeType GradeType
		expected  string
	}{
		{Assignment, "assignment"},
		{Exam, "exam"},
		{Essay, "essay"},
	}

	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			actual := test.gradeType.String()
			if actual != test.expected {
				t.Errorf("Expected %q; got %q",
					test.expected, actual)
			}
		})
	}
}
