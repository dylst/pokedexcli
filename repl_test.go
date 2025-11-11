package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct{
		input string 
		expected []string
	}{
		{
			input: "	hello  world	",
			expected: []string{"hello", "world"},
		},
		{
			input: " how	are you",
			expected: []string{"how", "are", "you"},
		},
		{
			input: "charizard is top tier",
			expected: []string{"charizard", "is", "top", "tier"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("fail")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("fail")
			}
		}
	}
}