package main

import (
//"fmt"

//"github.com/munsy/boyermoore"
)

const (
	// Enable/disable debugging
	DEBUG_ENABLED           = false
	BM_DEBUG_ENABLED        = true // false
	BCS_DEBUG_ENABLED       = true
	PC_DEBUG_ENABLED        = true
	BM_BREAK_DEBUG_ENABLED  = true
	BCR_TABLE_DEBUG_ENABLED = true

	// Test search/pattern strings
	TestSearchString = "WE HOLD THESE TRUTHS TO BE SELF-EVIDENT"
)

/*
func DebugPrint(searchString, patternString string, searchIndex, patternIndex int) {
	if DEBUG_ENABLED {
		fmt.Println(" ")
		PrintPosition(searchString, patternString, searchIndex, patternIndex)
		PrintCompare(searchString, patternString, searchIndex, patternIndex)
	}
}

func PrintCompare(searchString, patternString string, searchIndex, patternIndex int) {
	if PC_DEBUG_ENABLED {
		compareString := "Comparing '" + string(patternString[patternIndex])
		compareString += "' (pattern) with '" + string(searchString[searchIndex])
		compareString += "' (search)"
		fmt.Println(compareString)
	}
}

func PrintPosition(searchString, patternString string, currentPosition, patternPosition int) {
	caretBuilder := ""
	patternBuilder := ""
	patternLength := len(patternString) - 1

	fmt.Println(searchString, " | Current Position:", currentPosition, "/", len(searchString) - 1)

	// Build caret string
	for i := 0;  i < currentPosition; i++ {
		caretBuilder += " "
	}

	caretBuilder += "^"

	// Build pattern string
	for i := 0; i < currentPosition + patternLength - patternPosition; i++ {
		if i < len(patternString) - 1 {
			continue
		}
		patternBuilder += " "
	}

	patternBuilder += patternString

	fmt.Println(caretBuilder)
	fmt.Println(patternBuilder)
}

func PrintBadCharacterRule(pattern string) {
	if DEBUG_ENABLED {
		fmt.Println(TestSearchString)
		fmt.Println(pattern)
		fmt.Println("")
		fmt.Println("  --- BCR ---")

		table := makeBadCharacterRuleTable(pattern)

		for k, v := range table {
			fmt.Println("Key:", string(k), "Value:", v)
		}
		fmt.Println("---------------")
		fmt.Println("")
	}
}

func TestString(searchString, patternString string, willPatternBeFound bool) {
	if DEBUG_ENABLED {
		if BCR_TABLE_DEBUG_ENABLED {
			PrintBadCharacterRule(patternString)
		}
	}

	fmt.Println("----------------------------------------------------------")
	fmt.Println("Testing '" + patternString + "' against '" + searchString + "'")

	if bm(searchString, patternString) {
		fmt.Println("Found", patternString)
		if willPatternBeFound {
			fmt.Println("Test passed.")
		} else {
			fmt.Println("\t\t\tTest FAILED.")
		}
	} else {
		fmt.Println(patternString, "not found.")
		if !willPatternBeFound {
			fmt.Println("Test passed.")
		} else {
			fmt.Println("\t\t\tTest FAILED.")
		}
	}
}
*/
