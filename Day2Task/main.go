package main

import (
	taskfiletwo "daytwo/taskfile"
	"fmt"
)

func main() {
	sliceofnum := []int{23, 25, 46, 58, 10, 47}
	sortedslice := taskfiletwo.SortDescending(sliceofnum)
	fmt.Println(sortedslice)
}
