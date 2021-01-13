// 電通大で行われたコンピュータ囲碁講習会をGolangで追う
package main

import (
	// "bufio"

	"fmt"

	// "log"

	"os"

	// "sort"
	// "strconv"
	// "strings"
	// "sync"
	"time"
	// "unicode"
	// "unsafe"
)

var recordTime [MaxMoves]float64

func getCharZ(z int) string {
	if z == 0 {
		return "pass"
	}

	y := z / Width
	x := z - y*Width
	ax := 'A' + x - 1
	if ax >= 'I' {
		ax++
	}
	return string(ax) + string(BoardSize+1-y+'0')
}

var usiKomaKanjiV9a = [20]string{" 0", " 1", " 2", " 3", " 4", " 5", " 6", " 7", " 8", " 9",
	"❿", "⓫", "⓬", "⓭", "⓮", "⓯", "⓰", "⓱", "⓲", "⓳"}

// PrintBoardV9a - 盤を描画。
func PrintBoardV9a() {
	// var str = [4]string{"・", "●", "○", "＃"}
	var str = [4]string{" .", " *", " o", " #"}
	fmt.Fprintf(os.Stderr, "\n   ")
	for x := 0; x < BoardSize; x++ {
		fmt.Fprintf(os.Stderr, "%2d", x+1)
	}
	fmt.Fprintf(os.Stderr, "\n  +")
	for x := 0; x < BoardSize; x++ {
		fmt.Fprintf(os.Stderr, "--")
	}
	fmt.Fprintf(os.Stderr, "+\n")
	for y := 0; y < BoardSize; y++ {
		fmt.Fprintf(os.Stderr, "%s|", usiKomaKanjiV9a[y+1])
		for x := 0; x < BoardSize; x++ {
			fmt.Fprintf(os.Stderr, "%s", str[board[x+1+Width*(y+1)]])
		}
		fmt.Fprintf(os.Stderr, "|")
		if y == 4 {
			fmt.Fprintf(os.Stderr, "  koZ=%d,moves=%d", get81(koZ), moves)
		}
		fmt.Fprintf(os.Stderr, "\n")
	}
	fmt.Fprintf(os.Stderr, "  +")
	for x := 0; x < BoardSize; x++ {
		fmt.Fprintf(os.Stderr, "--")
	}
	fmt.Fprintf(os.Stderr, "+\n")
}

func getBestUctV9a(color int) int {
	max := -999
	nodeNum = 0
	uctLoop := 10000 // 多め
	var bestI = -1
	next := createNode()
	for i := 0; i < uctLoop; i++ {
		var boardCopy = [BoardMax]int{}
		koZCopy := koZ
		copy(boardCopy[:], board[:])

		searchUctV8(color, next)

		koZ = koZCopy
		copy(board[:], boardCopy[:])
	}
	pN := &node[next]
	for i := 0; i < pN.ChildNum; i++ {
		c := &pN.Children[i]
		if c.Games > max {
			bestI = i
			max = c.Games
		}
		// fmt.Fprintf(os.Stderr,"%2d:z=%2d,rate=%.4f,games=%3d\n", i, get81(c.Z), c.Rate, c.Games)
	}
	bestZ := pN.Children[bestI].Z
	fmt.Fprintf(os.Stderr, "bestZ=%d,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
		get81(bestZ), pN.Children[bestI].Rate, max, allPlayouts, nodeNum)
	return bestZ
}

func initBoard() {
	for i := 0; i < BoardMax; i++ {
		board[i] = 3
	}
	for y := 0; y < BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			board[getZ(x+1, y+1)] = 0
		}
	}
	moves = 0
	koZ = 0
}

func addMoves9a(z int, color int, sec float64) {
	err := putStoneV4(z, color, FillEyeOk)
	if err != 0 {
		fmt.Fprintf(os.Stderr, "Err!\n")
		os.Exit(0)
	}
	record[moves] = z
	recordTime[moves] = sec
	moves++
	PrintBoardV9a()
}

func playComputerMove(color int, fUCT int) int {
	var z int
	st := time.Now()
	allPlayouts = 0
	if fUCT != 0 {
		z = getBestUctV9a(color)
	} else {
		z = primitiveMonteCalroV9(color)
	}
	t := time.Since(st).Seconds()
	fmt.Fprintf(os.Stderr, "%.1f sec, %.0f playoutV9/sec, play_z=%2d,moves=%d,color=%d,playouts=%d\n",
		t, float64(allPlayouts)/t, get81(z), moves, color, allPlayouts)
	addMoves9a(z, color, t)
	return z
}
func undo() {

}
func testPlayoutV9a() {
	flagTestPlayout = 1
	playoutV9(1)
	PrintBoardV9a()
	printSgf()
}
