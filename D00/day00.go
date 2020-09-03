package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"sort"
	"math"
)

func main() {
	var sum int = 0
	var median float32
	var mean float32
	var mode int
	var sd float64 = 0

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
	sort.Ints(slice)
	if len(slice) % 2 == 1 {
		median = float32(slice[len(slice)/2])
	} else {
		median = float32(slice[len(slice)/2 - 1] + slice[len(slice)/2]) / 2
	}
	mean = float32(sum) / float32(len(arr))
	max_count := 0
	for _, n := range slice {
		sd += math.Pow(float64(float32(n) - mean), 2)
		if arr[n] > max_count {
			mode = n
			max_count = arr[n]
		}
	}
	sd = math.Sqrt(sd)
	fmt.Printf("mean = %.2f\nmedian = %.2f\nmode = %d\nSD = %.2f\n", mean, median, mode, sd)
}