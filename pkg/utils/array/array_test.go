package array

import "testing"

type TestCase struct {
	Name     string
	Elements []any
	Value    any
	Result   bool
}

func TestContains(t *testing.T) {
	testCases := []TestCase{
		{Name: "Success", Elements: []any{1, 2, 3, 4}, Value: 4, Result: true},
		{Name: "Failed", Elements: []any{1, 2, 3, 4}, Value: 5, Result: false},
		{Name: "Success", Elements: []any{"a", "b", "c"}, Value: "a", Result: true},
		{Name: "Failed", Elements: []any{"a", "b", "c"}, Value: "d", Result: false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if Contains(testCase.Elements, testCase.Value) != testCase.Result {
				t.Errorf("Expected %v, got %v", testCase.Result, testCase.Elements)
			}
		})
	}
}
