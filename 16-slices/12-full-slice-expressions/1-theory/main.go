// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true

	sliceable := []int{1, 2, 3, 4, 5}
	// https://stackoverflow.com/a/27938683
	fmt.Println(cap(sliceable[0:2:2]))
}
