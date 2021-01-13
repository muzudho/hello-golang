// 電通大で行われたコンピュータ囲碁講習会をGolangで追う
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func playoutV9(turnColor int) int {
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
		if flagTestPlayout != 0 {
			record[moves] = z
			moves++
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

func primitiveMonteCalroV9(color int) int {
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
				win := -playoutV9(flipColor(color))
				winSum += win
				koZ = koZCopy2
				copy(board[:], boardCopy2[:])
			}
			winRate = float64(winSum) / float64(tryNum)
			if winRate > bestValue {
				bestValue = winRate
				bestZ = z
				// fmt.Printf("bestZ=%d,color=%d,v=%5.3f,tryNum=%d\n", get81(bestZ), color, bestValue, tryNum)
			}
			koZ = koZCopy
			copy(board[:], boardCopy[:])
		}
	}
	return bestZ
}

func searchUctV9(color int, nodeN int) int {
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
		win = -playoutV9(flipColor(color))
	} else {
		if c.Next == NodeEmpty {
			c.Next = createNode()
		}
		win = -searchUctV9(flipColor(color), c.Next)
	}
	c.Rate = (c.Rate*float64(c.Games) + float64(win)) / float64(c.Games+1)
	c.Games++
	pN.ChildGameSum++
	return win
}

func getBestUctV9(color int) int {
	max := -999
	nodeNum = 0
	uctLoop := 1000 // 少な目
	var bestI = -1
	next := createNode()
	for i := 0; i < uctLoop; i++ {
		var boardCopy = [BoardMax]int{}
		koZCopy := koZ
		copy(boardCopy[:], board[:])

		searchUctV9(color, next)

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
		// fmt.Printf("%2d:z=%2d,rate=%.4f,games=%3d\n", i, get81(c.Z), c.Rate, c.Games)
	}
	bestZ := pN.Children[bestI].Z
	fmt.Printf("bestZ=%d,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
		get81(bestZ), pN.Children[bestI].Rate, max, allPlayouts, nodeNum)
	return bestZ
}

func getComputerMove(color int, fUCT int) int {
	var z int
	st := time.Now()
	allPlayouts = 0
	if fUCT != 0 {
		z = getBestUctV9(color)
	} else {
		z = primitiveMonteCalroV9(color)
	}
	t := time.Since(st).Seconds()
	fmt.Printf("%.1f sec, %.0f playoutV9/sec, play_z=%2d,moves=%d,color=%d,playouts=%d\n",
		t, float64(allPlayouts)/t, get81(z), moves, color, allPlayouts)
	return z
}

func printSgf() {
	fmt.Printf("(;GM[1]SZ[%d]KM[%.1f]PB[]PW[]\n", BoardSize, komi)
	for i := 0; i < moves; i++ {
		z := record[i]
		y := z / Width
		x := z - y*Width
		var sStone = [2]string{"B", "W"}
		fmt.Printf(";%s", sStone[i&1])
		if z == 0 {
			fmt.Printf("[]")
		} else {
			fmt.Printf("[%c%c]", x+'a'-1, y+'a'-1)
		}
		if ((i + 1) % 10) == 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf(")\n")
}
func selfplay() {
	color := 1

	for {
		fUCT := 1
		if color == 1 {
			fUCT = 0
		}
		z := getComputerMove(color, fUCT)
		addMovesV8(z, color)
		if z == 0 && moves > 1 && record[moves-2] == 0 {
			break
		}
		if moves > 300 {
			break
		} // too long
		color = flipColor(color)
	}

	printSgf()
}

func testPlayout() {
	flagTestPlayout = 1
	playoutV9(1)
	PrintBoardV8()
	printSgf()
}
