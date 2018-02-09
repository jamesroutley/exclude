package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var usage string = `usage: exclude [FORBIDDEN file] [INPUT file]

Iterates though the lines in INPUT, and prints the lines that aren't also in
FORBIDDEN to stdout.`

func main() {
	flag.Parse()
	if len(flag.Args()) < 2 {
		exit(usage)
	}

	forbidden, err := getForbidden(flag.Arg(0))
	if err != nil {
		exit(err)
	}

	if err := exclude(forbidden, flag.Arg(1)); err != nil {
		exit(err)
	}
}

func exit(a interface{}) {
	fmt.Fprintln(os.Stderr, a)
	os.Exit(1)
}

func getForbidden(forbiddenFile string) (map[string]bool, error) {
	var forbidden = make(map[string]bool)
	file, err := os.Open(forbiddenFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		forbidden[scanner.Text()] = true
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return forbidden, err
}

func exclude(forbidden map[string]bool, inputFile string) error {
	file, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if !forbidden[s] {
			fmt.Println(s)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
