package main

import (
	"fmt"
	"testing"
)

type categoryTest struct {
	questionNumber   int
	expectedCategory string
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
		actual := categoryGetter(tt.questionNumber)
		if actual != tt.expectedCategory {
			t.Errorf("categoryGetter(%d): expected %v, actual %v", tt.questionNumber, tt.expectedCategory, actual)
		}
	}
}

func TestResultChecker(t *testing.T) {
	resultChecker("Usually", "Fear of Conflict")
	// if err != nil {
	// 	fmt.Printf("Result checker failed %v.\n", err)
	// 	return
	// }
	fmt.Print(totals)
	if totals["Fear of Conflict"] != 3 {
		t.Error("Result checker failed")
	}
}
