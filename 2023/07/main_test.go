package main

import "testing"

func TestHandStrength2(t *testing.T) {
	testCases := []struct {
		Hand     string
		Expected HandStrength
	}{
		{"2222", FiveOfKind},
		{"222J", FiveOfKind},
		{"222JJ", FiveOfKind},
		{"22JJJ", FiveOfKind},
		{"22223", FourOfKind},
		{"222J3", FourOfKind},
		{"222J3", FourOfKind},
		{"22JJ3", FourOfKind},
		{"2JJJ3", FourOfKind},
		{"223J3", FullHouse},
		{"22333", FullHouse},
		{"22234", ThreeOfKind},
		{"22J34", ThreeOfKind},
		{"2JJ34", ThreeOfKind},
		{"22334", TwoPair},
		{"234J5", OnePair},
		{"22345", OnePair},
		{"23456", HighCard},
	}
	for _, test := range testCases {
		result := HandStrength2(test.Hand)
		if result != test.Expected {
			t.Errorf(
                "HandStrength2 returned %s for hand %s but expected %s",
                result.String(),
                test.Hand,
                test.Expected.String(),
            )
		}
	}
}
