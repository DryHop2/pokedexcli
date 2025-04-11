package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Pikachu Thunderbolt  ",
			expected: []string{"pikachu", "thunderbolt"},
		},
		{
			input:    "   Bulbasaur    Razor Leaf",
			expected: []string{"bulbasaur", "razor", "leaf"},
		},
		{
			input:    "\tCharizard FLAMETHROWER\n",
			expected: []string{"charizard", "flamethrower"},
		},
		{
			input:    "Snorlax   Rest",
			expected: []string{"snorlax", "rest"},
		},
		{
			input:    "    Mewtwo      Shadow    Ball   ",
			expected: []string{"mewtwo", "shadow", "ball"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) returned %d words, want %d", c.input, len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("cleanInput(%q)[%d] = %q, want %q", c.input, i, actual[i], c.expected[i])
			}
		}
	}
}
