package lettercombinations

import (
	"fmt"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	datalist := LetterCombinations("23")
	fmt.Println("datalist:", datalist)
}
