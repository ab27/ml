package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)

var p = fmt.Println

type ft struct {
	f []int
	t int
}

type dist struct {
	target   int
	distance float64
}

type ByDistance []dist

func (this ByDistance) Len() int {
	return len(this)
}

func (this ByDistance) Less(i, j int) bool {
	return this[i].distance < this[j].distance
}

func (this ByDistance) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func main() {
	lines, err := readLines("digits.txt")
	check(err)

	dataSet := []ft{}

	for _, line := range lines {
		l := strings.Split(line, ",")
		target, _ := strconv.Atoi(strings.TrimSpace(l[1]))
		features := []int{}

		featuresString := strings.TrimSpace(l[0])
		featuresSlice := strings.Split(featuresString, " ")

		for _, v := range featuresSlice {
			f, _ := strconv.Atoi(v)
			features = append(features, f)
		}

		dataSet = append(dataSet, ft{features, target})

	}

	// randomize dataSet because it is initially ordered 0 to 9
	for i := range dataSet {
		j := rand.Intn(i + 1)
		dataSet[i], dataSet[j] = dataSet[j], dataSet[i]
	}

	correct := 0
	incorrect := 0

	for _, v := range dataSet[751:] {
		//p("target", dataSet[1400].t)
		n := findNeighbors(7, dataSet[:750], v.f)
		pred := predict(n)

		if pred == v.t {
			correct++
		} else {
			incorrect++
		}
	}

	p("correct: ", correct)
	p("incorrect: ", incorrect)
	p("Accuracy: ", (correct*100)/len(dataSet[751:]), "%")

}

func findNeighbors(k int, dataSet []ft, target []int) []dist {
	n := make([]dist, k)

	for i := 0; i < k; i++ {
		n[i].distance = 1000000.0
	}

	for _, v := range dataSet {
		d := euclideanDistance(target, v.f)

		if d < n[k-1].distance {
			n[k-1] = dist{v.t, d}
		}

		sort.Sort(ByDistance(n))

	}

	return n

}

func predict(n []dist) int {
	m := map[int]int{}
	mostFreq := -1

	for _, v := range n {
		m[v.target]++

		if m[v.target] > mostFreq {
			mostFreq = v.target
		}
	}

	return mostFreq

}

// mahnhatan distance

func euclideanDistance(x []int, y []int) float64 {
	sum := 0

	for k, v := range y {
		sum += (v - x[k]) * (v - x[k])
	}

	return math.Sqrt(float64(sum))
}
