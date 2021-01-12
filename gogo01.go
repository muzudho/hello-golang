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

const (
	// BoardSize - 何路盤。
	BoardSize = 9
	// Width - 枠込み。
	Width = (BoardSize + 2)
	// BoardMax - 枠込み盤の配列サイズ。
	BoardMax = (Width * Width)
)

// 9路盤用
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

var usiKomaKanji = [20]string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九",
	"❿", "⓫", "⓬", "⓭", "⓮", "⓯", "⓰", "⓱", "⓲", "⓳"}

// PrintBoard - 盤の描画。
func PrintBoard() {
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

func main() {
	PrintBoard()
}
