package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
	"io"
)

func checkPaths(filename1, filename2, command string) {
	paths := make(map[string]int)
	file, err := os.Open(filename1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		paths[string(line)]++
	}
	file, err = os.Open(filename2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	reader = bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if paths[string(line)] == 0 {
			fmt.Println(command, string(line))
		}
	}
}

func main() {
	newFlag := flag.String("new", "", "Filename with new snapshot")
	oldFlag := flag.String("old", "", "Filename with original snapshot")
	flag.Parse()
	if *newFlag == "" || *oldFlag == "" {
		os.Exit(1)
	}
	newFilename := *newFlag
	oldFilename := *oldFlag
	checkPaths(newFilename, oldFilename, "REMOVED")
	checkPaths(oldFilename, newFilename, "ADDED")
}
