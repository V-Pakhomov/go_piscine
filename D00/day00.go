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

	arr := make(map[int]int)
	slice := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
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
	fmt.Println("mean =", mean)
	fmt.Println(slice)
}