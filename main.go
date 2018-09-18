package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := os.Args[1]
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		log.Fatal("Can not scan input file")
	}
	t := scanner.Text()
	t = strings.TrimRight(t, "\n")
	expectCaseNum, _ := strconv.Atoi(t)

	actualCaseNum := 0
	for i := 1; scanner.Scan(); i++ {
		t := scanner.Text()
		t = strings.TrimRight(t, "\n")
		a := strings.Split(t, " ")
		s := solv(i, a)
		fmt.Printf("Case #%d: %v\n", i, s)
		actualCaseNum = i
	}
	if expectCaseNum != actualCaseNum {
		log.Fatal("Total case number does not match actual case number.")
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func solv(i int, a []string) []string {
	return a
}
