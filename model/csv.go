package main

import (
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
)

// slice <-> csv
// test diffmatchpatch

const (
	str1 = "hoge\n hoge\n fuga"
	str2 = "fuga\n hoge\n fuga"
)

func main() {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(str1, str2, false)
	for i := 0; i < 2; i ++ {
		fmt.Println(diffs[i])
	}
}
