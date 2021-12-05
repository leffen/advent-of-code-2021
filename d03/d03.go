package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	lines, err := importFile("data.txt")
	if err != nil {
		logrus.Fatal(err)
	}

	gamma := calcGamma(lines)
	epsi := calcEpsilon(lines)
	fmt.Printf("Power %d\n", gamma*epsi)

	oxygen := calcOxygen(lines)
	scrubber := calcScrubber(lines)
	fmt.Printf("Rating %d\n", oxygen*scrubber)
}

func importFile(fileName string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	return strings.Split(string(fileBytes), "\n"), nil
}

func calcGamma(lines []string) int64 {
	maxWidth := len(lines[0])
	nums := ""
	for i := 0; i < maxWidth; i++ {
		bits := []int{0, 0}
		for _, l := range lines {
			bits[AtoI(l[i])]++
		}
		if bits[0] > bits[1] {
			nums = nums + "0"
		} else {
			nums = nums + "1"
		}
	}

	fmt.Printf("gamma = %v\n", nums)
	return BinaryToInt(nums)
}

func calcEpsilon(lines []string) int64 {
	maxWidth := len(lines[0])
	nums := ""
	for i := 0; i < maxWidth; i++ {
		bits := []int{0, 0}
		for _, l := range lines {
			bits[AtoI(l[i])]++
		}
		if bits[0] < bits[1] {
			nums = nums + "0"
		} else {
			nums = nums + "1"
		}
	}

	fmt.Printf("epsilon = %v\n", nums)
	return BinaryToInt(nums)
}

func calcOxygen(lines []string) int64 {
	maxWidth := len(lines[0])
	i := 0
	for {
		bits := []int{0, 0}
		for _, l := range lines {
			bits[AtoI(l[i])]++
		}
		if bits[0] > bits[1] {
			lines = filter(i, '0', lines)
		} else {
			lines = filter(i, '1', lines)
		}
		if len(lines) == 1 {
			rc := BinaryToInt(lines[0])

			fmt.Printf("oxygen = %v l=%v\n", rc, lines[0])
			return rc
		}

		i++
		if i >= maxWidth {
			logrus.Fatalf("Failed to located row. i=%d rc len=%d", i, len(lines))
		}
	}
}

func calcScrubber(lines []string) int64 {
	maxWidth := len(lines[0])

	i := 0
	for {
		bits := []int{0, 0}
		for _, l := range lines {
			bits[AtoI(l[i])]++
		}
		if bits[0] <= bits[1] {
			lines = filter(i, '0', lines)
		} else {
			lines = filter(i, '1', lines)
		}
		if len(lines) == 1 {
			rc := BinaryToInt(lines[0])
			fmt.Printf("scrubber = %v l=%v\n", rc, lines[0])
			return rc
		}

		i++
		if i >= maxWidth {
			logrus.Fatalf("Failed to located row. i=%d rc len=%d", i, len(lines))
		}
	}
}

func filter(pos int, mask byte, lines []string) []string {
	rc := []string{}
	for _, l := range lines {
		if l[pos] == mask {
			rc = append(rc, l)
		}
	}

	return rc
}

func AtoI(data byte) int {
	i, err := strconv.Atoi(string(data))
	if err != nil {
		logrus.Fatal(err)
	}

	return i
}

func BinaryToInt(item string) int64 {
	i, err := strconv.ParseInt(item, 2, 64)
	if err != nil {
		logrus.Fatal(err)
	}
	return i
}
