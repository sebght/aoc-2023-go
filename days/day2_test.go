package days

import (
	"testing"
)

func TestFirstPartDay2OneValidGame(t *testing.T) {
	result := FirstPartDay2([]string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	})
	expectedResult := 1
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, expectedResult)
	}
}

func TestFirstPartDay2OneInvalidGame(t *testing.T) {
	result := FirstPartDay2([]string{
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	})
	expectedResult := 0
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, expectedResult)
	}
}

func TestFirstPartDay2OneMoreComplexInvalidGame(t *testing.T) {
	result := FirstPartDay2([]string{
		"Game 8: 8 green, 4 blue; 17 green, 4 red; 10 blue, 5 green, 9 red; 9 green, 8 red, 3 blue; 9 green, 5 red, 2 blue",
	})
	expectedResult := 0
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, expectedResult)
	}
}

func TestFirstPartDay2OfficalExample(t *testing.T) {
	result := FirstPartDay2([]string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	})
	expectedResult := 8
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, expectedResult)
	}
}
