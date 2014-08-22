package validator

import (
	"testing"
)

type testPair struct {
	acronym     string
	productName []string
	out         bool
}

var testPairsIsValid = []testPair{
	{"GIZTK3", []string{"GISi", "Zombie", "Tracker", "3000"}, true},
	{"GZT3", []string{"GISi", "Zombie", "Tracker", "3000"}, true},
	{"GISEE0", []string{"GISi", "Zombie", "Tracker", "3000"}, true},
	{"GZT3k", []string{"GISi", "Zombie", "Tracker", "3000"}, false},
	{"GZT", []string{"GISi", "Zombie", "Tracker", "3000"}, false},
	{"GZTK", []string{"GISi", "Zombie", "Tracker", "3000"}, false},
	{"BLAH", []string{"GISi", "Zombie", "Tracker", "3000"}, false},
	{"GOOSE", []string{"Google", "Search", "Engine"}, true},
	{"GEANIE", []string{"Google", "Search", "Engine"}, true},
	{"OOSN", []string{"Google", "Search", "Engine"}, true},
	{"LEACHE", []string{"Google", "Search", "Engine"}, true},
	// {"OOGLEE", []string{"Google", "Search", "Engine"}, true},
	// {"GOOSEEE", []string{"Google", "Search", "Engine"}, true},
	// {"GGSEE", []string{"Google", "Search", "Engine"}, true},
	// {"GGSE", []string{"Google", "Search", "Engine"}, true},
	// {"SGE", []string{"Google", "Search", "Engine"}, false},
	// {"GEANIEOOSN", []string{"Google", "Search", "Engine"}, false},
	// {"OGGLEE", []string{"Google", "Search", "Engine"}, false},
	// {"GGG", []string{"Google", "Search", "Engine"}, false},
	// {"GGSEA", []string{"Google", "Search", "Engine"}, false},
}

func TestIsValid(t *testing.T) {
	for _, pair := range testPairsIsValid {
		result := IsValid(pair.acronym, pair.productName)
		if result != pair.out {
			t.Errorf("Expected %v to be %v, but got %v", pair.acronym, pair.out, result)
		}
	}
}
