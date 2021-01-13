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

func countScoreV7(turnColor int) int {
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
	if float64(score)-komi > 0 {
		win = 1
	}
	if turnColor == 2 {
		win = -win
	} // gogo07

	// fmt.Printf("blackSum=%2d, (stones=%2d, area=%2d)\n", blackSum, kind[1], blackArea)
	// fmt.Printf("whiteSum=%2d, (stones=%2d, area=%2d)\n", whiteSum, kind[2], whiteArea)
	// fmt.Printf("score=%d, win=%d\n", score, win)
	return win
}

func playoutV7(turnColor int) int {
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
		// PrintBoard()
		// fmt.Printf("loop=%d,z=%d,c=%d,emptyNum=%d,koZ=%d\n",
		// 	loop, get81(z), color, emptyNum, get81(koZ))
		color = flipColor(color)
	}
	return countScoreV7(turnColor)
}

func primitiveMonteCalroV7(color int) int {
	tryNum := 30
	bestZ := 0
	var bestValue, winRate float64
	var boardCopy = [BoardMax]int{}
	koZCopy := koZ
	copy(boardCopy[:], board[:])
	bestValue = -100.0

	for y := 0; y <= BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			z := getZ(x+1, y+1)
			if board[z] != 0 {
				continue
			}
			err := putStoneV4(z, color, FillEyeErr)
			if err != 0 {
				continue
			}

			winSum := 0
			for i := 0; i < tryNum; i++ {
				var boardCopy2 = [BoardMax]int{}
				koZCopy2 := koZ
				copy(boardCopy2[:], board[:])
				win := -playout(flipColor(color))
				winSum += win
				koZ = koZCopy2
				copy(board[:], boardCopy2[:])
			}
			winRate = float64(winSum) / float64(tryNum)
			if winRate > bestValue {
				bestValue = winRate
				bestZ = z
				fmt.Printf("bestZ=%d,color=%d,v=%5.3f,tryNum=%d\n", get81(bestZ), color, bestValue, tryNum)
			}
			koZ = koZCopy
			copy(board[:], boardCopy[:])
		}
	}
	return bestZ
}
