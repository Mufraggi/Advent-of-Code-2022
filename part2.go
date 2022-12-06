package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Strategy int

const (
	WinStrat Strategy = iota
	DrawStrat
	LooseStrat
)

func part2() {
	f, err := openFile("./files/d2.txt")
	if err != nil {
		panic("file not found")
	}

	p2Score := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		moves := strings.Split(row, " ")
		p2Score = getMoveToPlay(moves[0], moves[1], p2Score)
	}
	fmt.Println(p2Score)
}

func getMoveToPlay(p1 string, p2 string, acc int) int {
	moveStrategy := getMove(p2)
	_, gameresult2, move := play(getP1Move(p1), moveStrategy)
	return acc + int(gameresult2) + int(move)
}

func play(p1 Move, strategy Strategy) (GameResult, GameResult, Move) {
	if strategy == WinStrat {
		if p1 == Rock {
			return Loose, Win, Paper
		} else if p1 == Paper {
			return Loose, Win, Scissors
		} else {
			return Loose, Win, Rock
		}
	} else if strategy == DrawStrat {
		return Draw, Draw, p1
	} else {
		if p1 == Rock {
			return Win, Loose, Scissors
		} else if p1 == Paper {
			return Win, Loose, Rock
		} else {
			return Win, Loose, Paper
		}
	}
}

func getMove(p2 string) Strategy {
	if p2 == "Y" {
		return DrawStrat
	} else if p2 == "X" {
		return LooseStrat
	} else {
		return WinStrat
	}
}
