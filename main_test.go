package main

import (
	"testing"
)

type categoryTest struct {
	questionNumber int
	category       string
}

var categoryTests = []categoryTest{
	{1, "Fear of Conflict"}, {2, "Avoidance of Accountability"}, {3, "Lack of Commitment"},
	{4, "Absence of Trust"}, {5, "Inattention to Results"}, {6, "Absence of Trust"},
	{7, "Fear of Conflict"}, {8, "Lack of Commitment"}, {9, "Inattention to Results"},
	{10, "Fear of Conflict"}, {11, "Avoidance of Accountability"}, {12, "Absence of Trust"},
	{13, "Lack of Commitment"}, {14, "Avoidance of Accountability"}, {15, "Inattention to Results"},
}

func TestCategoryGetter(t *testing.T) {
	for _, tt := range categoryTests {
		t.Logf("Question #: %d Category: %v", tt.questionNumber, tt.category)
		actual := categoryGetter(tt.questionNumber)
		if actual != tt.category {
			t.Errorf("categoryGetter(%d): want %v, got %v", tt.questionNumber, tt.category, actual)
		}
	}
}

func TestCategoryGetterFail(t *testing.T) {
	qnum := 16
	result := categoryGetter(qnum)
	if categoryGetter(22) != "Invalid queston number" {
		t.Errorf("categoryGetter(%d): want: Invalid queston number got %v", qnum, result)
	}
}

func TestResultChecker(t *testing.T) {
	resultChecker("Usually", "Fear of Conflict")
	resultChecker("Sometimes", "Absence of Trust")
	resultChecker("Rarely", "Avoidance of Accountability")

	if totals["Fear of Conflict"] != 3 {
		t.Error("Result checker failed, got:", totals["Fear of Conflict"], "want: 3")
	}
	if totals["Absence of Trust"] != 2 {
		t.Error("Result checker failed, got:", totals["Absence of Trust"], "want: 2")
	}
	if totals["Avoidance of Accountability"] != 1 {
		t.Error("Result checker failed, got:", totals["Avoidance of Accountability"], "want: 1")
	}
}

func Test_main(t *testing.T) {
	main()
	if totals["Absence of Trust"] != 2 {
		t.Error("Result checker failed, got:", totals["Absence of Trust"], "want: 2")
	}
}
