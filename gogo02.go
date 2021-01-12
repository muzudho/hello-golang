// Source: https://github.com/bleu48/GoGo
// 電通大で行われたコンピュータ囲碁講習会をGolangで追う
package main

import (
	// "bufio"
	"fmt"
	// "log"
	// "math"
	// "math/rand"
	// "os"
	// "sort"
	// "strconv"
	// "strings"
	// "sync"
	// "time"
	// "unicode"
	// "unsafe"
)

// putStoneV2 - 石を置きます。
func putStoneV2(tz int, color int) int {
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
	// if wall+mycolSafe==4 {return 3}
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

// PrintBoardV2 - 盤の描画。
func PrintBoardV2() {
	// "● " - Visual Studio Code の 全角半角崩れ対応。
	// "○ " - Visual Studio Code の 全角半角崩れ対応。
	var str = [4]string{"・", "● ", "○ ", "＃"}
	fmt.Printf("\n   ")
	for x := 0; x < BoardSize; x++ {
		fmt.Printf("%2d", x+1)
	}
	fmt.Printf("\n  +")
	for x := 0; x < BoardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
	for y := 0; y < BoardSize; y++ {
		fmt.Printf("%s|", usiKomaKanji[y+1])
		for x := 0; x < BoardSize; x++ {
			fmt.Printf("%s", str[board[x+1+Width*(y+1)]])
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("  +")
	for x := 0; x < BoardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
}
