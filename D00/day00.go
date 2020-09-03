package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"sort"
)

func main() {
	var sum int = 0
	var mean float32
	var mode int

	arr := make(map[int]int)
	slice := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			if len(scanner.Text()) != 0 {
				fmt.Println("ERROR")
				os.Exit(2)
			}
		} else {
			slice = append(slice, num)
			arr[num]++
			sum += num
		}
	}
	fmt.Println("sum =", sum)
	fmt.Println(arr)
	sort.Ints(slice)
	if len(slice) % 2 == 1 {
		mean = float32(slice[len(slice)/2])
	} else {
		mean = float32(slice[len(slice)/2 - 1] + slice[len(slice)/2]) / 2
	}
	max_count := 0
	for _, n := range slice {
		if arr[n] > max_count {
			mode = n
			max_count = arr[n]
		}
	}
	fmt.Println("mean =", mean, "mode =", mode)
	fmt.Println(slice)
}