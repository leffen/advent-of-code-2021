package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	cmds, err := loadDirectionsFile("data.txt")
	if err != nil {
		logrus.Fatal(err)
	}
	num := walk(cmds)
	fmt.Printf("Num:%d\n", num)

	num = walk2(cmds)
	fmt.Printf("Num 2:%d\n", num)
}

type Cmd struct {
	direction string
	distance  int64
}

func loadDirectionsFile(fileName string) ([]*Cmd, error) {
	rc := []*Cmd{}

	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	for _, ln := range strings.Split(string(fileBytes), "\n") {
		fields := strings.Split(ln, " ")
		i, err := strconv.Atoi(fields[1])
		if err != nil {
			logrus.Fatal(err)
		}

		cmd := &Cmd{direction: fields[0], distance: int64(i)}
		rc = append(rc, cmd)
	}

	return rc, nil
}

func walk(cmds []*Cmd) int64 {

	horiz := int64(0)
	depth := int64(0)
	for _, c := range cmds {
		switch {
		case c.direction == "forward":
			horiz += c.distance
		case c.direction == "down":
			depth += c.distance
		case c.direction == "up":
			depth -= c.distance
		}
	}
	return horiz * depth
}

func walk2(cmds []*Cmd) int64 {

	aim := int64(0)
	depth := int64(0)
	horiz := int64(0)
	for _, c := range cmds {
		switch {
		case c.direction == "forward":
			horiz += c.distance
			depth += (aim * c.distance)
		case c.direction == "down":
			aim += c.distance
		case c.direction == "up":
			aim -= c.distance
		}
	}
	//logrus.Infof("Horiz:%d depth:%d\n", horiz, depth)
	return horiz * depth
}
