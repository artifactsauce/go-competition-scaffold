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
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		log.Fatal("Can not read stdin")
	}
	t := scanner.Text()
	t = strings.TrimRight(t, "\n")
	expectCaseNum, _ := strconv.Atoi(t)

	actualCaseNum := 0
	for i := 1; scanner.Scan(); i++ {
		t := scanner.Text()
		t = strings.TrimRight(t, "\n")
		s := strings.Split(t, " ")
		r := solv(i, s)
		fmt.Printf("Case #%d: %v\n", i, r)
		actualCaseNum = i
	}
	if expectCaseNum != actualCaseNum {
		log.Fatal("Total case number does not match actual case number.")
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func solv(i int, s []string) []string {
	return s
}
