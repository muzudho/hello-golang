// 電通大で行われたコンピュータ囲碁講習会をGolangで追う
package main

import (
	"fmt"
)

var usiKomaKanji = [20]string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九",
	"❿", "⓫", "⓬", "⓭", "⓮", "⓯", "⓰", "⓱", "⓲", "⓳"}

// PrintBoardV1 - 盤の描画。
func PrintBoardV1() {
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
