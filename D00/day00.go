package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"sort"
	"math"
)

func arrToMap(arr []int) map[int]int {
	dict := make(map[int]int)
	for _, v := range arr {
		dict[v]++
	}
	return dict
}

func arrSum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func readStdin() []int {
	arr := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			if len(scanner.Text()) != 0 {
				fmt.Println("It is not an integer")
				os.Exit(2)
			}
		} else {
			arr = append(arr, num)
		}
	}
	sort.Ints(arr)
	return arr
}

func calcMean(arr []int) float32 {
	return float32(arrSum(arr)) / float32(len(arr))
}

func calcMedian(arr []int) float32 {
	if len(arr) % 2 == 0 {
		return float32(arr[len(arr)/2] + arr[len(arr)/2 - 1]) / 2
	} else {
		return float32(arr[len(arr)/2])
	}
}

func calcMode(arr []int) int {
	mode := 0
	max_amount := 0
	dict := arrToMap(arr)
	for _, v := range arr {
		if dict[v] > max_amount {
			max_amount = dict[v]
			mode = v
		}
	}
	return mode
}

func calcSD(arr []int) float64 {
	sd := 0.0
	mean := calcMean(arr)
	for _, v := range arr {
		sd += math.Pow(float64(float32(v) - mean), 2)
	}
	sd = math.Sqrt(sd)
	return sd
}

func main() {
	arr := readStdin()
	mean := calcMean(arr)
	median := calcMedian(arr)
	mode := calcMode(arr)
	sd := calcSD(arr)
	fmt.Printf("mean = %.2f\nmedian = %.2f\nmode = %d\nSD = %.2f\n", mean, median, mode, sd)
}