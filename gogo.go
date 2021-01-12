package main

import (
	"fmt"
)

// "github.com\muzudho\hello-golang"
// "github.com/muzudho/hello-golang/gogo01.go"
// "github.com/muzudho/hello-golang"

const (
	// BoardSize - 何路盤。
	BoardSize = 9
	// Width - 枠込み。
	Width = (BoardSize + 2)
	// BoardMax - 枠込み盤の配列サイズ。
	BoardMax = (Width * Width)
)

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

// gogo02.go 用。
var board = [BoardMax]int{
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, //    1 2 3 4 5 6 7 8 9
	3, 0, 0, 0, 0, 0, 2, 0, 0, 0, 3, // 1 ������������������
	3, 0, 0, 0, 0, 2, 1, 2, 2, 2, 3, // 2 ������������������
	3, 0, 0, 0, 0, 2, 1, 1, 1, 1, 3, // 3 ������������������
	3, 0, 0, 0, 0, 0, 2, 1, 2, 2, 3, // 4 ������������������
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, // 5 ������������������
	3, 0, 1, 2, 0, 0, 0, 0, 0, 0, 3, // 6 ������������������
	3, 1, 2, 0, 2, 0, 0, 0, 0, 0, 3, // 7 ������������������
	3, 0, 1, 2, 0, 2, 2, 1, 1, 0, 3, // 8 ������������������
	3, 0, 0, 0, 0, 2, 1, 0, 2, 1, 3, // 9 ������������������
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
}

func main() {
	GoGoV1()
	// GoGoV2()
}

// GoGoV1 - バージョン１。
func GoGoV1() {
	PrintBoardV1()
}

// GoGoV2 - バージョン２。
func GoGoV2() {
	PrintBoardV2()
	err := put_stone(get_z(7, 5), 2)
	fmt.Printf("err=%d\n", err)
	PrintBoardV2()
}
