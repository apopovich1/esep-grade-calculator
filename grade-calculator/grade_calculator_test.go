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

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 79, Assignment)
	gradeCalculator.AddGrade("exam 1", 60, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 75, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 65, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 68, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 30, Assignment)
	gradeCalculator.AddGrade("exam 1", 95, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 40, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestAssignmentString(t *testing.T){
	expected_value := "assignment"
	actual_value := Assignment.String()
	if expected_value != actual_value{
		t.Errorf("Expected Assignment String to return '%s', got '%s' instead", expected_value, actual_value)
	}
}

//Technically the two tests below are not required for the 100% coverage, but I included them to have more thorough testing
func TestExamString(t *testing.T){
	expected_value := "exam"
	actual_value := Exam.String()
	if expected_value != actual_value{
		t.Errorf("Expected Exam String to return '%s', got '%s' instead", expected_value, actual_value)
	}
}

func TestEssayString(t *testing.T){
	expected_value := "essay"
	actual_value := Essay.String()
	if expected_value != actual_value{
		t.Errorf("Expected Essay String to return '%s', got '%s' instead", expected_value, actual_value)
	}
}

func TestPass(t *testing.T){
	expected_value := "Pass"
	GradeCalculator :=NewGradeCalculator()
	GradeCalculator.AddGrade("assignment", 70, Assignment)
	GradeCalculator.AddGrade("exam",70,Exam)
	GradeCalculator.AddGrade("essay",70, Essay)
	_ = GradeCalculator.GetFinalGrade()
	actual_value := GradeCalculator.isPassing
	if expected_value != actual_value{
		t.Errorf("Expected '%s' at 70, got '%s' instead", expected_value, actual_value)
	}
}

func TestFail(t *testing.T){
	expected_value := "Fail"
	GradeCalculator :=NewGradeCalculator()
	GradeCalculator.AddGrade("assignment", 60, Assignment)
	GradeCalculator.AddGrade("exam",40,Exam)
	GradeCalculator.AddGrade("essay",30, Essay)
	_ = GradeCalculator.GetFinalGrade()
	actual_value := GradeCalculator.isPassing
	if expected_value != actual_value{
		t.Errorf("Expected '%s' below 70, got '%s' instead", expected_value, actual_value)
	}
}

func TestEmpty(t *testing.T){
	GradeCalculator := NewGradeCalculator()
	expected_value := "Fail"
	_ = GradeCalculator.GetFinalGrade()
	actual_value := GradeCalculator.isPassing
	if expected_value != actual_value{
		t.Errorf("Expected '%s' when empty, got '%s' instead", expected_value, actual_value)
	}
}