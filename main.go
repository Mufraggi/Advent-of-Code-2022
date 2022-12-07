package main

import (
	"bufio"
	"os"
)

func main() {
	//fmt.Println(int('w'))
	//process("BfzSzRSVVMVNRMDDNBSNSnfBmbrglGQbmNpQggFjpgpbQlQb")
	f, err := openFile("./files/d3.txt")
	if err != nil {
		panic("file not found")
	}
	groupBy3(f)
	//res := parseAndGroup(f)
	//fmt.Println(res)

}

func openFile(fileName string) (*os.File, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return f, nil
}
func parseAndGroup(f *os.File) int32 {
	var res int32
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		res += process(row)
	}
	return res
}

func process(row string) int32 {
	part1, part2 := splitRow(row)
	res := commpare(part1, part2)
	if len(res) > 0 {
		return calScore(res[0])
	}
	return 0

}

func splitRow(row string) (map[int32]int32, map[int32]int32) {
	mid := len(row) / 2
	return toSet(row[:mid]), toSet(row[mid:])
}

func toSet(str string) map[int32]int32 {
	m := make(map[int32]int32)
	for _, c := range str {
		m[c] = c
	}
	return m
}

func commpare(m1 map[int32]int32, m2 map[int32]int32) []int32 {
	var l []int32
	for _, s := range m1 {
		_, res := m2[s]
		if res {
			l = append(l, s)
		}
	}
	return l
}

func calScore(c int32) int32 {
	if c >= 65 && c <= 90 {
		return c - 38
	} else if c >= 97 && c <= 122 {
		return c - 96
	}
	return 0
}
