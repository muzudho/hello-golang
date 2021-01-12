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

func putStoneV4(tz int, color int, fillEyeErr int) int {
	var around = [4][3]int{}
	var liberty, stone int
	unCol := flipColor(color)
	space := 0
	wall := 0
	mycolSafe := 0
	captureSum := 0
	koMaybe := 0

	if tz == 0 {
		koZ = 0
		return 0
	}
	for i := 0; i < 4; i++ {
		around[i][0] = 0
		around[i][1] = 0
		around[i][2] = 0
		z := tz + dir4[i]
		c := board[z]
		if c == 0 {
			space++
		}
		if c == 3 {
			wall++
		}
		if c == 0 || c == 3 {
			continue
		}
		countLiberty(z, &liberty, &stone)
		around[i][0] = liberty
		around[i][1] = stone
		around[i][2] = c
		if c == unCol && liberty == 1 {
			captureSum += stone
			koMaybe = z
		}
		if c == color && liberty >= 2 {
			mycolSafe++
		}

	}
	if captureSum == 0 && space == 0 && mycolSafe == 0 {
		return 1
	}
	if tz == koZ {
		return 2
	}
	if wall+mycolSafe == 4 && fillEyeErr == FillEyeErr {
		return 3
	}
	if board[tz] != 0 {
		return 4
	}

	for i := 0; i < 4; i++ {
		lib := around[i][0]
		c := around[i][2]
		if c == unCol && lib == 1 && board[tz+dir4[i]] != 0 {
			takeStone(tz+dir4[i], unCol)
		}
	}

	board[tz] = color

	countLiberty(tz, &liberty, &stone)
	if captureSum == 1 && stone == 1 && liberty == 1 {
		koZ = koMaybe
	} else {
		koZ = 0
	}
	return 0
}

// func get_empty_z() int {
// 	var x, y, z int
// 	for {
// 		x = rand.Intn(9) + 1
// 		y = rand.Intn(9) + 1
// 		z = getZ(x, y)
// 		if board[z] == 0 {
// 			break
// 		}
// 	}
// 	return z
// }

// func play_one_move(color int) int {
// 	var z int
// 	for i := 0; i < 100; i++ {
// 		z := get_empty_z()
// 		err := putStoneV4(z, color)
// 		if err == 0 {
// 			return z
// 		}
// 	}
// 	z = 0
// 	putStoneV4(0, color)
// 	return z
// }

// playoutV4 - 最後まで石を打ちます。
func playoutV4(turnColor int) int {
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
	return 0
}
