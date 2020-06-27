package main

import (
	"fmt"

	"github.com/gmhcode/hello/morestrings"

	"github.com/google/go-cmp/cmp"
)

func hello() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
