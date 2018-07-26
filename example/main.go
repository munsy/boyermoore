package main

import (
	"fmt"

	"github.com/munsy/boyermoore"
)

func main() {
	bm := boyermoore.NewBoyerMoore(TestSearchString)
	fmt.Println(bm.ToSearch, bm.Search("TRUTH"))
}
