package main

import "testing"

func TestPredict(t *testing.T) {
    testCases := []struct {
        History History
        Expected int
    } {
        {History{0, 3, 6, 9, 12, 15}, 18},
        {History{1, 3, 6, 10, 15, 21}, 28},
        {History{10, 13, 16, 21, 30, 45}, 68},
        {History{0, -2, -4}, -6},
    }
    for _, testCase := range testCases {
        result := Predict(testCase.History)
        if result != testCase.Expected {
            t.Errorf(
                "Predict failed for %v, got: %v, expected: %v",
                testCase.History, result, testCase.Expected,
            )
        }
    }
}
