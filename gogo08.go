// 電通大で行われたコンピュータ囲碁講習会をGolangで追う
package main

import (
	// "bufio"
	"fmt"
	"math"

	// "log"

	"math/rand"
	"os"
	// "sort"
	// "strconv"
	// "strings"
	// "sync"
	// "unicode"
	// "unsafe"
)

// PrintBoardV8 - 盤の描画。
func PrintBoardV8() {
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
		fmt.Printf("|")
		if y == 4 {
			fmt.Printf("  koZ=%d,moves=%d", get81(koZ), moves)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("  +")
	for x := 0; x < BoardSize; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
}

func playoutV8(turnColor int) int {
	color := turnColor
	previousZ := 0
	loopMax := BoardSize*BoardSize + 200

	allPlayouts++
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

// UCT
const (
	Childrenize = BoardSize*BoardSize + 1
	NodeMax     = 10000
	NodeEmpty   = -1
	IllegalZ    = -1
)

// Child - 子。
type Child struct {
	Z     int
	Games int
	Rate  float64
	Next  int
}

// Node - ノード。
type Node struct {
	ChildNum     int
	Children     [Childrenize]Child
	ChildGameSum int
}

var node = [NodeMax]Node{}
var nodeNum = 0

func addChild(pN *Node, z int) {
	n := pN.ChildNum
	pN.Children[n].Z = z
	pN.Children[n].Games = 0
	pN.Children[n].Rate = 0.0
	pN.Children[n].Next = NodeEmpty
	pN.ChildNum++
}

func createNode() int {
	if nodeNum == NodeMax {
		fmt.Printf("node over Err\n")
		os.Exit(0)
	}
	pN := &node[nodeNum]
	pN.ChildNum = 0
	pN.ChildGameSum = 0
	for y := 0; y <= BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			z := getZ(x+1, y+1)
			if board[z] != 0 {
				continue
			}
			addChild(pN, z)
		}
	}
	addChild(pN, 0)
	nodeNum++
	return nodeNum - 1
}

func selectBestUcb(nodeN int) int {
	pN := &node[nodeN]
	selectI := -1
	maxUcb := -999.0
	ucb := 0.0
	for i := 0; i < pN.ChildNum; i++ {
		c := &pN.Children[i]
		if c.Z == IllegalZ {
			continue
		}
		if c.Games == 0 {
			ucb = 10000.0 + 32768.0*rand.Float64()
		} else {
			ucb = c.Rate + 1.0*math.Sqrt(math.Log(float64(pN.ChildGameSum))/float64(c.Games))
		}
		if ucb > maxUcb {
			maxUcb = ucb
			selectI = i
		}
	}
	if selectI == -1 {
		fmt.Printf("Err! select\n")
		os.Exit(0)
	}
	return selectI
}

func searchUctV8(color int, nodeN int) int {
	pN := &node[nodeN]
	var c *Child
	var win int
	for {
		selectI := selectBestUcb(nodeN)
		c = &pN.Children[selectI]
		z := c.Z
		err := putStoneV4(z, color, FillEyeErr)
		if err == 0 {
			break
		}
		c.Z = IllegalZ
		// fmt.Printf("ILLEGAL:z=%2d\n", get81(z))
	}
	if c.Games <= 0 {
		win = -playoutV8(flipColor(color))
	} else {
		if c.Next == NodeEmpty {
			c.Next = createNode()
		}
		win = -searchUctV8(flipColor(color), c.Next)
	}
	c.Rate = (c.Rate*float64(c.Games) + float64(win)) / float64(c.Games+1)
	c.Games++
	pN.ChildGameSum++
	return win
}

func getBestUctV8(color int) int {
	max := -999
	nodeNum = 0
	uctLoop := 10000
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
		fmt.Printf("%2d:z=%2d,rate=%.4f,games=%3d\n", i, get81(c.Z), c.Rate, c.Games)
	}
	bestZ := pN.Children[bestI].Z
	fmt.Printf("bestZ=%d,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
		get81(bestZ), pN.Children[bestI].Rate, max, allPlayouts, nodeNum)
	return bestZ
}

func addMoves(z int, color int) {
	err := putStoneV4(z, color, FillEyeOk)
	if err != 0 {
		fmt.Printf("Err!\n")
		os.Exit(0)
	}
	record[moves] = z
	moves++
	PrintBoardV8()
}
