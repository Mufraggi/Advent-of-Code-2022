package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := openFile("./files/d1.txt")
	if err != nil {
		panic("file not found")
	}
	elfs := parseAndGroup(f)
	elfsCalsTotal := sumCalForElf(elfs)
	sort.Ints(elfsCalsTotal)
	fmt.Println(elfsCalsTotal[len(elfsCalsTotal)-1])
	resTop3 := []int{
		elfsCalsTotal[len(elfsCalsTotal)-1],
		elfsCalsTotal[len(elfsCalsTotal)-2],
		elfsCalsTotal[len(elfsCalsTotal)-3],
	}
	fmt.Println(sum(resTop3))
}

func openFile(fileName string) (*os.File, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func parseAndGroup(f *os.File) [][]int {
	elfs := [][]int{}
	var elfCaloris []int
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		row := scanner.Text()
		if len(row) == 0 {
			elfs = append(elfs, elfCaloris)
			elfCaloris = []int{}
		} else {
			cal, err := strconv.Atoi(row)
			if err != nil {
				panic("cal not number")
			}
			elfCaloris = append(elfCaloris, cal)
		}
	}
	return elfs
}

func sumCalForElf(elfs [][]int) []int {
	elfsCalsTotal := []int{}
	for _, elf := range elfs {
		elfsCalsTotal = append(elfsCalsTotal, sum(elf))
	}
	return elfsCalsTotal
}

func sum(cals []int) int {
	res := 0
	for _, cal := range cals {
		res += cal
	}
	return res
}
