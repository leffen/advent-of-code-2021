package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	tx := loadData("data.txt", false, false)
	num := tx.NumOverlapPoints()
	fmt.Printf("Num overlaps %d\n", num)

	tx2 := loadData("data.txt", true, false)
	num2 := tx2.NumOverlapPoints()
	fmt.Printf("Num2 overlaps %d\n", num2)

}

func loadData(filename string, includeDiagonal, dbg bool) *TravelMap {
	lines, err := importFile(filename)

	if err != nil {
		logrus.Fatal(err)
	}

	tx := NewTravelMap()
	tx.IncludeDiagonal = includeDiagonal
	for idx, ln := range lines {
		if len(strings.TrimSpace(ln)) == 0 {
			continue
		}
		if dbg {
			fmt.Printf("Adding %s\n", ln)
		}
		tx.AddLine(ln, idx+1)
		if dbg {
			tx.Show()
		}
	}
	return tx
}

func importFile(fileName string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	return strings.Split(string(fileBytes), "\n"), nil
}

type Koord struct {
	x        int
	y        int
	line     string
	overlaps int
}

func NewKoord(koords string) *Koord {
	items := strings.Split(koords, ",")
	k := &Koord{}
	k.x = atoi(items[0])
	k.y = atoi(items[1])
	return k
}

func (k *Koord) ID() string {
	return fmt.Sprintf("%d,%d", k.x, k.y)
}

type TravelMap struct {
	IncludeDiagonal bool
	cells           map[string]*Koord
	minX, minY      int
	maxX, maxY      int
}

func NewTravelMap() *TravelMap {
	return &TravelMap{cells: map[string]*Koord{}, IncludeDiagonal: false, minX: 100000, minY: 100000}
}

func (t *TravelMap) NumOverlapPoints() int {
	rc := 0
	for _, cell := range t.cells {
		if cell.overlaps > 0 {
			rc++
		}
	}
	return rc
}

func (t *TravelMap) Show() {
	fmt.Print("    ")
	for x := t.minX; x <= t.maxX; x++ {
		fmt.Printf("%1d", x%10)
	}
	fmt.Println("")
	for y := t.minY; y <= t.maxY; y++ {
		fmt.Printf("%03d ", y)
		for x := t.minX; x <= t.maxX; x++ {
			fmt.Print(t.getPoint(x, y))
		}
		fmt.Println("")
	}
}
func (t *TravelMap) getPoint(x, y int) string {
	k := &Koord{x: x, y: y, overlaps: 0}
	id := k.ID()
	c, ok := t.cells[id]
	if !ok {
		return "."
	}
	if c.overlaps > 0 {
		return fmt.Sprintf("%d", c.overlaps+1)
	}
	return "1"
}

func (t *TravelMap) AddLine(ln string, lnNum int) {
	//fmt.Printf("AddLine %s\n", ln)
	items := strings.Split(ln, "->")
	fk := NewKoord(strings.TrimSpace(items[0]))
	tk := NewKoord(strings.TrimSpace(items[1]))

	if !t.IncludeDiagonal && fk.x != tk.x && fk.y != tk.y {
		// Do nothing with diagonal lines
		return
	}

	if fk.x == tk.x {
		// Horiz line
		y1 := iif(fk.y < tk.y, fk.y, tk.y)
		y2 := iif(fk.y < tk.y, tk.y, fk.y)
		t.addXLine(fk.x, y1, y2, lnNum)
		return
	}

	if fk.y == tk.y {
		// Vert line
		x1 := iif(fk.x < tk.x, fk.x, tk.x)
		x2 := iif(fk.x < tk.x, tk.x, fk.x)

		t.addYLine(fk.y, x1, x2, lnNum)
		return
	}

	t.addDLine(fk.x, fk.y, tk.x, tk.y, lnNum)
}

func (t *TravelMap) addDLine(x1, y1, x2, y2 int, lnNum int) {
	fmt.Printf("(%d,%d) -> (%d,%d)\n", x1, y1, x2, y2)

	xs := iif(x1 < x2, x1, x2)
	xm := iif(x1 < x2, x2, x1)

	// TODO FIX
	y := iif(x1 < x2, y1, y2)
	dy := 1
	if (x1 >= x2 && y2 > y1) || (x1 < x2 && y1 > y2) {
		dy = -1
	}

	// dy := iif(x1 < x2 && y1 < y2, 1, y2)

	fmt.Printf("DLINE xs=%d xm=%d dy=%d y=%d\n", xs, xm, dy, y)
	for x := xs; x <= xm; x++ {
		fmt.Printf("  setpoint %d,%d\n", x, y)
		t.setPoint(x, y, fmt.Sprintf("%d", lnNum))
		y = y + dy
		// if y > ym+1 {
		// 	logrus.Fatalf("y should be less than y2 %d < %d  xs=%d xm=%d ys=%d ym=%d", y, ym, xs, xm, ys, ym)
		// }
	}
}

func (t *TravelMap) addXLine(x, y1, y2 int, lnNum int) {
	for y := y1; y <= y2; y++ {
		t.setPoint(x, y, fmt.Sprintf("%d", lnNum))
	}
}

func (t *TravelMap) addYLine(y, x1, x2 int, lnNum int) {
	for x := x1; x <= x2; x++ {
		t.setPoint(x, y, fmt.Sprintf("%d", lnNum))
	}
}

func (t *TravelMap) setPoint(x, y int, marker string) {
	k := &Koord{x: x, y: y, overlaps: 0, line: marker}
	id := k.ID()
	c, ok := t.cells[id]
	if ok {
		c.overlaps = c.overlaps + 1
		return
	}
	t.cells[k.ID()] = k

	if k.x < t.minX {
		t.minX = k.x
	}
	if k.x > t.maxX {
		t.maxX = k.x
	}

	if k.y < t.minY {
		t.minY = k.y
	}
	if k.y > t.maxY {
		t.maxY = k.y
	}

}

func iif(expr bool, r1, r2 int) int {
	if expr {
		return r1
	}
	return r2
}

func atoi(data string) int {
	i, err := strconv.Atoi(strings.TrimSpace(data))
	if err != nil {
		logrus.Fatalf("Num=%s err=%s", data, err)
	}

	return i
}
