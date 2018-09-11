package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	for i := 0; scanner.Scan(); i++ {
		t := scanner.Text()
		s := strings.Split(t, " ")
		solv(i, s)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func solv(i int, s []string) {
	fmt.Printf("#%d %v\n", i, s)
}
