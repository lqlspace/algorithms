package search

import (
	"testing"
)


func TestCreateRandNumsJSON(t *testing.T) {
	CreateRandNumsJSON("AAAA", 1000)
}


func TestCreateRandNumsJSONStream(t *testing.T) {
	CreateRandNumsJSONStream("cccc", 10000)
}
