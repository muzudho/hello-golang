package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	// BoardSize - 何路盤。
	BoardSize = 9
	// Width - 枠込み。
	Width = (BoardSize + 2)
	// BoardMax - 枠込み盤の配列サイズ。
	BoardMax = (Width * Width)
)

var usiKomaKanji = [20]string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九",
	"❿", "⓫", "⓬", "⓭", "⓮", "⓯", "⓰", "⓱", "⓲", "⓳"}

/*
// gogo01.go 用
var board = [BoardMax]int{
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 1, 1, 0, 1, 0, 0, 0, 0, 3,
	3, 2, 2, 1, 1, 0, 1, 2, 0, 0, 3,
	3, 2, 0, 2, 1, 2, 2, 1, 1, 0, 3,
	3, 0, 2, 2, 2, 1, 1, 1, 0, 0, 3,
	3, 0, 0, 0, 2, 1, 2, 1, 0, 0, 3,
	3, 0, 0, 2, 0, 2, 2, 1, 0, 0, 3,
	3, 0, 0, 0, 0, 2, 1, 1, 0, 0, 3,
	3, 0, 0, 0, 0, 2, 2, 1, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 2, 1, 0, 0, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
}
*/

/*
// gogo02.go 用。
var board = [BoardMax]int{
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 0, 0, 0, 0, 0, 2, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 2, 1, 2, 2, 2, 3,
	3, 0, 0, 0, 0, 2, 1, 1, 1, 1, 3,
	3, 0, 0, 0, 0, 0, 2, 1, 2, 2, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 1, 2, 0, 0, 0, 0, 0, 0, 3,
	3, 1, 2, 0, 2, 0, 0, 0, 0, 0, 3,
	3, 0, 1, 2, 0, 2, 2, 1, 1, 0, 3,
	3, 0, 0, 0, 0, 2, 1, 0, 2, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
}
*/

// gogo03.go 用。
var board = [BoardMax]int{
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
}

var dir4 = [4]int{1, Width, -1, -Width}
var koZ int

func getZ(x int, y int) int {
	return y*Width + x
}

func get81(z int) int {
	y := z / Width
	x := z - y*Width
	if z == 0 {
		return 0
	}
	return x*10 + y
}

func flipColor(col int) int {
	return 3 - col
}

var checkBoard = [BoardMax]int{}

func countLibertySub(tz int, color int, pLiberty *int, pStone *int) {
	checkBoard[tz] = 1
	*pStone++
	for i := 0; i < 4; i++ {
		z := tz + dir4[i]
		if checkBoard[z] != 0 {
			continue
		}
		if board[z] == 0 {
			checkBoard[z] = 1
			*pLiberty++
		}
		if board[z] == color {
			countLibertySub(z, color, pLiberty, pStone)
		}
	}

}

func countLiberty(tz int, pLiberty *int, pStone *int) {
	*pLiberty = 0
	*pStone = 0
	for i := 0; i < BoardMax; i++ {
		checkBoard[i] = 0
	}
	countLibertySub(tz, board[tz], pLiberty, pStone)
}

func takeStone(tz int, color int) {
	board[tz] = 0
	for i := 0; i < 4; i++ {
		z := tz + dir4[i]
		if board[z] == color {
			takeStone(z, color)
		}
	}
}

func main() {
	// GoGoV1()
	// GoGoV2()
	GoGoV3()
}

// GoGoV1 - バージョン１。
func GoGoV1() {
	PrintBoardV1()
}

// GoGoV2 - バージョン２。
func GoGoV2() {
	PrintBoardV2()
	err := putStoneV2(getZ(7, 5), 2)
	fmt.Printf("err=%d\n", err)
	PrintBoardV2()
}

// GoGoV3 - バージョン３。
func GoGoV3() {
	color := 1
	rand.Seed(time.Now().UnixNano())
	for {
		z := playOneMove(color)
		fmt.Printf("moves=%4d, color=%d, z=%d\n", moves, color, get81(z))
		PrintBoardV3()

		record[moves] = z
		moves++
		if moves == 1000 {
			fmt.Printf("max moves!\n")
			break
		}
		if z == 0 && moves >= 2 && record[moves-2] == 0 {
			fmt.Printf("two pass\n")
			break
		}
		color = flipColor(color)
	}
}
