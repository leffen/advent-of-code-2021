package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	nums := importFileAsIntA("data.txt")
	cnt := count(nums)
	fmt.Printf("Count: %d\n", cnt)

	cnt2 := countTris(nums)
	fmt.Printf("Count tries: %d\n", cnt2)
}

func count(nums []int) int64 {
	rc := int64(0)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			rc++
		}
	}
	return rc
}

func countTris(nums []int) int64 {
	rc := int64(0)
	for i := 0; i < len(nums)-3; i++ {
		s1 := nums[i] + nums[i+1] + nums[i+2]
		s2 := nums[i+1] + nums[i+2] + nums[i+3]
		if s2 > s1 {
			rc++
		}
	}
	return rc
}

func importFileAsIntA(fileName string) []int {
	lines, err := importFile(fileName)
	if err != nil {
		logrus.Fatal(err)
	}
	return linesToIntA(lines)
}

func importFile(fileName string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	return strings.Split(string(fileBytes), "\n"), nil
}

func linesToIntA(lines []string) []int {
	rc := []int{}
	for _, l := range lines {
		if len(strings.TrimSpace(l)) == 0 {
			continue
		}
		i, err := strconv.Atoi(l)
		if err != nil {
			logrus.Fatal(err)
		}
		rc = append(rc, i)
	}
	return rc
}
