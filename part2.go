package main

import (
	"bufio"
	"fmt"
	"os"
)

func groupBy3(f *os.File) {
	scanner := bufio.NewScanner(f)
	var group []string
	res := int32(0)
	for scanner.Scan() {
		row := scanner.Text()
		if len(group) == 3 {
			res += process3(group[0], group[1], group[2])
			group = []string{}
		}
		group = append(group, row)
	}
	fmt.Println(res)
}

func process3(i1 string, i2 string, i3 string) int32 {
	i1Part := toSet(i1)
	i2Part := toSet(i2)
	i3Part := toSet(i3)
	res := int32(0)
	for _, c := range i1Part {
		if ifInMap(c, i2Part) && ifInMap(c, i3Part) {
			res = calScore(c)
		}
	}
	return res
}

func ifInMap(c int32, m map[int32]int32) bool {
	_, b := m[c]
	return b
}
