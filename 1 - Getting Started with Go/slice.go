package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	sli := make([]int, 0)
	var txt string
	var ap int
	for {
		fmt.Scan(&txt)
		if txt == "X" {
			break
		}
		ap, _ = strconv.Atoi(txt)
		sli = append(sli, ap)
		sort.SliceStable(sli, func(i, j int) bool { return sli[i] < sli[j] })
		fmt.Println(sli)
	}
	sort.SliceStable(sli, func(i, j int) bool { return sli[i] < sli[j] })
}
