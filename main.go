package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Name struct {
	nb int
}

func main() {
	f, err := openFile("./files/d4.txt")
	if err != nil {
		panic("file not found")
	}
	scanner := bufio.NewScanner(f)
	res := 0
	for scanner.Scan() {
		row := scanner.Text()
		if splitSectionAndProcess(row) {
			fmt.Println(splitSectionAndProcess(row))
			res += 1
		}
	}
	println(res)
}

func openFile(fileName string) (*os.File, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func containte(min1 int, max1 int, min2 int, max2 int) bool {
	if min1 <= min2 && min2 <= max1 && min1 <= max2 && max2 <= max1 {
		return true
	}
	return false
}

func containteV2(min1 int, max1 int, min2 int, max2 int) bool {
	if min1 <= min2 && min2 <= max1 || min1 <= max2 && max2 <= max1 {
		return true
	}
	return false
}

func splitSectionAndProcess(row string) bool {
	split := strings.Split(row, ",")
	min1, max1 := getMinAx(split[0])
	min2, max2 := getMinAx(split[1])
	res1 := containte(min1, max1, min2, max2)
	res2 := containte(min2, max2, min1, max1)
	return res1 || res2
}

func getMinAx(split string) (int, int) {
	sec1 := strings.Split(split, "-")
	min, _ := strconv.Atoi(sec1[0])
	max, _ := strconv.Atoi(sec1[1])
	return min, max
}
