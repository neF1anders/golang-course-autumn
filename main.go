package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func twoSum(nums []int, target int) []int {
	hm := make(map[int]int)
	for i, j := range nums {
		if val, ok := hm[target-j]; ok {
			return []int{val, i}
		}
		hm[j] = i
	}
	return []int{}
}

func main() {
	if inFile, err := os.Open("input.txt"); err == nil {
		os.Stdin = inFile
		defer inFile.Close()
	}
	if outFile, err := os.Create("output.txt"); err == nil {
		os.Stdout = outFile
		defer outFile.Close()
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	readInt := func() int {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		return val
	}
	t := readInt()

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for tc := 0; tc < t; tc++ {
		n := readInt()
		target := readInt()

		nums := make([]int, n)
		for i := 0; i < n; i++ {
			nums[i] = readInt()
		}
		ans := twoSum(nums, target)
		fmt.Fprintf(writer, "%v %v\n", ans[0], ans[1])
	}
}
