package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := os.Args[1]
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		solv(t)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func solv(f string) {
	fmt.Printf("%v\n", f)
}
