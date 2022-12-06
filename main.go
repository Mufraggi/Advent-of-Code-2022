package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type GameResult int64

const (
	Loose GameResult = 0
	Draw  GameResult = 3
	Win   GameResult = 6
)

type Move int

const (
	Rock Move = iota + 1
	Paper
	Scissors
)

func main() {
	part2()
}

/*func main() {
	f, err := openFile("./files/d2.txt")
	if err != nil {
		panic("file not found")
	}
	score1, score2 := iterGameByGame(f)
	fmt.Println(score1)
	fmt.Println(score2)

}*/

func openFile(fileName string) (*os.File, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func iterGameByGame(f *os.File) (int, int) {
	p1Score := 0
	p2Score := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		p1, p2 := parseRow(row)
		res1, res2 := getResultForGame(p1, p2)
		p1Score = addScore(p1Score, res1, p1)
		p2Score = addScore(p2Score, res2, p2)
	}
	return p1Score, p2Score
}

func addScore(acc int, res GameResult, move Move) int {
	fmt.Println(int(res))
	return acc + int(res) + int(move)
}

func parseRow(row string) (Move, Move) {
	moves := strings.Split(row, " ")
	return getP1Move(moves[0]), getP2Move(moves[1])
}

func getP1Move(move string) Move {
	if move == "A" {
		return Rock
	}
	if move == "B" {
		return Paper
	}
	return Scissors
}

func getP2Move(move string) Move {
	if move == "X" {
		return Rock
	}
	if move == "Y" {
		return Paper
	}
	return Scissors
}

func getResultForGame(p1 Move, p2 Move) (GameResult, GameResult) {
	if p1 == Rock && p2 == Scissors {
		return Win, Loose
	} else if p1 == Paper && p2 == Rock {
		return Win, Loose
	} else if p1 == Scissors && p2 == Paper {
		return Win, Loose
	} else if p2 == Rock && p1 == Scissors {
		return Loose, Win
	} else if p2 == Paper && p1 == Rock {
		return Loose, Win
	} else if p2 == Scissors && p1 == Paper {
		return Loose, Win
	} else {
		return Draw, Draw
	}
}
