// Source: https://github.com/bleu48/GoGo
// 電通大で行われたコンピュータ囲碁講習会をGolangで追う
package main

import (
	// "bufio"
	"fmt"
	// "log"
	// "math"
	"math/rand"
	// "os"
	// "sort"
	// "strconv"
	// "strings"
	// "sync"
	// "unicode"
	// "unsafe"
)

// countScore - 得点計算。
func countScore(turnColor int) int {
	var mk = [4]int{}
	var kind = [3]int{0, 0, 0}
	var score, blackArea, whiteArea, blackSum, whiteSum int

	for y := 0; y < BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			z := getZ(x+1, y+1)
			c := board[z]
			kind[c]++
			if c != 0 {
				continue
			}
			mk[1] = 0
			mk[2] = 0
			for i := 0; i < 4; i++ {
				mk[board[z+dir4[i]]]++
			}
			if mk[1] != 0 && mk[2] == 0 {
				blackArea++
			}
			if mk[2] != 0 && mk[1] == 0 {
				whiteArea++
			}
		}
	}
	blackSum = kind[1] + blackArea
	whiteSum = kind[2] + whiteArea
	score = blackSum - whiteSum
	win := 0
	if float32(score)-komi > 0 {
		win = 1
	}
	fmt.Printf("blackSum=%2d, (stones=%2d, area=%2d)\n", blackSum, kind[1], blackArea)
	fmt.Printf("whiteSum=%2d, (stones=%2d, area=%2d)\n", whiteSum, kind[2], whiteArea)
	fmt.Printf("score=%d, win=%d\n", score, win)
	return win
}

// playoutV5 - 最後まで石を打ちます。得点を返します。
func playoutV5(turnColor int) int {
	color := turnColor
	previousZ := 0
	loopMax := BoardSize*BoardSize + 200

	for loop := 0; loop < loopMax; loop++ {
		var empty = [BoardMax]int{}
		var emptyNum, r, z int
		for y := 0; y <= BoardSize; y++ {
			for x := 0; x < BoardSize; x++ {
				z = getZ(x+1, y+1)
				if board[z] != 0 {
					continue
				}
				empty[emptyNum] = z
				emptyNum++
			}
		}
		r = 0
		for {
			if emptyNum == 0 {
				z = 0
			} else {
				r = rand.Intn(emptyNum)
				z = empty[r]
			}
			err := putStoneV4(z, color, FillEyeErr)
			if err == 0 {
				break
			}
			empty[r] = empty[emptyNum-1]
			emptyNum--
		}
		if z == 0 && previousZ == 0 {
			break
		}
		previousZ = z
		PrintBoardV3()
		fmt.Printf("loop=%d,z=%d,c=%d,emptyNum=%d,koZ=%d\n",
			loop, get81(z), color, emptyNum, get81(koZ))
		color = flipColor(color)
	}
	return countScore(turnColor)
}
