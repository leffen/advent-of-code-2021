package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	g := loadData("data.txt")
	fmt.Printf("Score %d\n", g.DoDraw())
	fmt.Printf("Score2 %d\n", g.DoDraw2())
}

func loadData(fileName string) *Game {
	g := &Game{Draws: []int{}, Boards: []*Board{}}

	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		logrus.Fatal(err)
	}

	b := &Board{}
	lnNum := 0
	for _, ln := range strings.Split(string(fileBytes), "\n") {
		if strings.Contains(ln, ",") {
			g.Draws = stringToIntArray(ln, ",")
			continue
		}
		if len(strings.TrimSpace(ln)) == 0 {
			if len(b.cells) > 0 {
				g.AddBoard(b)
			}
			b = &Board{}
			lnNum = 0
			continue
		}
		ln = strings.TrimSpace(ln)
		for idx, nm := range stringToIntArray(ln, " ") {
			b.NewCell(nm, idx, lnNum)
		}
		lnNum++
	}
	if len(b.cells) > 0 {
		g.AddBoard(b)
	}

	return g
}

type Game struct {
	NumDraws int
	Draws    []int
	Boards   []*Board
}

func (g *Game) DoDraw() int {
	rc := []*Board{}

	for idx, num := range g.Draws {
		for _, brd := range g.Boards {
			brd.Check(num)
			if brd.HaveBingo() {
				brd.Show()
				rc = append(rc, brd)
			}
		}
		if len(rc) > 0 {
			g.NumDraws = idx + 1
			return rc[0].GetScore() * num
		}
	}
	return -1
}

func (g *Game) DoDraw2() int {
	g.ResetDraw()
	score := 0
	for _, num := range g.Draws {
		for _, brd := range g.Boards {
			if brd.haveBingo {
				continue
			}
			brd.Check(num)
			if brd.HaveBingo() {
				score = brd.GetScore() * num
				fmt.Printf("Found BINGO in %d with score %d\n", brd.ID, score)
			}
		}
	}
	return score
}
func (g *Game) ResetDraw() {
	for _, brd := range g.Boards {
		for _, c := range brd.cells {
			c.Checked = false
		}
	}
}
func (g *Game) AddBoard(b *Board) {
	//b.Show()
	b.ID = len(g.Boards) + 1
	g.Boards = append(g.Boards, b)
}

type Cell struct {
	Num     int
	Col     int
	Row     int
	Checked bool
}

type Board struct {
	ID        int
	haveBingo bool
	cells     []*Cell
	cmap      map[int]*Cell
	maxCol    int
	maxRow    int
	Score     int
}

func (b *Board) Show() {
	b.UpdateMap()
	fmt.Printf("Booard %d (mc:%d mr: %d nc: %d )----------\n", b.ID, b.maxCol, b.maxRow, len(b.cells))
	for row := 0; row <= b.maxRow; row++ {
		for col := 0; col <= b.maxCol; col++ {
			num := row*(b.maxCol+1) + col
			cell, ok := b.cmap[num]
			if !ok {
				break
			}
			cout := fmt.Sprintf("%d", cell.Num)
			if cell.Checked {
				cout = cout + "*"
			}
			fmt.Printf("%s ", cout)
		}
		fmt.Println("")
	}
}

func (b *Board) Check(num int) {
	if b.haveBingo {
		return
	}
	for _, c := range b.cells {
		if c.Num == num {
			c.Checked = true
			return
		}
	}
}
func (b *Board) UpdateMap() {
	b.cmap = map[int]*Cell{}

	for _, c := range b.cells {
		b.cmap[c.Row*(b.maxCol+1)+c.Col] = c
	}
}

func (b *Board) GetScore() int {
	score := 0
	for _, c := range b.cells {
		if !c.Checked {
			score += c.Num
		}
	}
	b.Score = score
	return score
}

func (b *Board) HaveBingo() bool {
	if b.haveBingo {
		return true
	}

	b.UpdateMap()

	// Horizontal checks
	for row := 0; row <= b.maxRow; row++ {
		numInRow := 0
		for col := 0; col <= b.maxCol; col++ {
			num := row*(b.maxCol+1) + col
			cell, ok := b.cmap[num]
			if !ok {
				break
			}
			if !cell.Checked {
				break
			}
			numInRow++
		}
		if numInRow == b.maxRow+1 {
			//			logrus.Debugf("FOUND horiz numInRow=%d maxRow=%v %v", numInRow, b.maxRow, b)
			b.haveBingo = true
			return true
		}
	}

	// Vertical
	for col := 0; col <= b.maxCol; col++ {
		numInCol := 0
		for row := 0; row <= b.maxRow; row++ {
			num := row*(b.maxCol+1) + col
			cell, ok := b.cmap[num]
			if !ok {
				break
			}
			if !cell.Checked {
				break
			}
			numInCol++
		}
		if numInCol == b.maxCol+1 {
			//			logrus.Debugf("FOUND vert numInRow=%d maxRow=%v %v", numInCol, b.maxCol, b)
			b.haveBingo = true
			return true
		}
	}

	// Diagonal
	numInDiag := 0
	for col := 0; col <= b.maxCol; col++ {
		num := col*(b.maxCol+1) + col
		cell, ok := b.cmap[num]
		if !ok {
			break
		}
		if !cell.Checked {
			break
		}
		numInDiag++
	}
	if numInDiag == b.maxRow+1 {
		b.haveBingo = true
		// logrus.Debugf("FOUND diag numInDiag=%d maxCol=%v %v", numInDiag, b.maxCol, b)
	}
	return numInDiag == b.maxRow+1
}

func (b *Board) NewCell(num, col, row int) Cell {
	c := Cell{Num: num, Col: col, Row: row, Checked: false}
	b.AddCell(&c)
	//	logrus.Debugf("Added cell %v", c)
	return c
}

func (b *Board) AddCell(c *Cell) {
	b.cells = append(b.cells, c)
	if c.Col > b.maxCol {
		b.maxCol = c.Col
	}

	if c.Row > b.maxRow {
		b.maxRow = c.Row
	}

}

func NewBoard(nums []int, colCount int) *Board {
	b := &Board{}
	xl := len(nums)

	rowNum := 0
	for i := 0; i < xl; i++ {
		c := Cell{Num: nums[i], Col: i % colCount, Row: rowNum, Checked: false}
		b.cells = append(b.cells, &c)
		if i > 0 && (i%colCount == 0) {
			rowNum++
		}
	}

	return b
}

func stringToIntArray(ln string, sep string) []int {
	rc := []int{}
	items := strings.Split(ln, sep)
	//logrus.Debugf("ITEMS=%v", items)
	for _, num := range items {
		if strings.TrimSpace(num) == "" {
			continue
		}

		//		logrus.Debugf("ITEM=%v", num)
		rc = append(rc, atoi(num))
	}
	return rc
}

func atoi(data string) int {
	i, err := strconv.Atoi(strings.TrimSpace(data))
	if err != nil {
		logrus.Fatalf("Num=%s err=%s", data, err)
	}

	return i
}
