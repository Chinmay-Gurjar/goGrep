package search

import (
	"testing"
	"fmt"
)

func assertResult(expectedCount int, mcount int, searchResults [][]Results){
	if expectedCount != mcount {
		fmt.Errorf("This should be same: %v", searchResults)
	}
}

func printError(err error) {
	if err != nil {
		fmt.Errorf("%v", err)
	}
}

func TestSearch(t *testing.T) {
	pattern := "is"
	filePath := []string{"test/testFile.txt"}
	expectedCount := 4
	searchResults, mcount, err := Search(pattern, filePath)
	printError(err)
	assertResult(expectedCount, mcount, searchResults)

	pattern = ""
	filePath = []string{"test/testFile.txt"}
	expectedCount = 3
	searchResults, mcount, err = Search(pattern, filePath)
	printError(err)
	assertResult(expectedCount, mcount, searchResults)

	pattern = "aa["
	filePath = []string{"test/testFile.txt"}
	expectedCount = 3
	searchResults, mcount, err = Search(pattern, filePath)
	printError(err)
	assertResult(expectedCount, mcount, searchResults)
}
