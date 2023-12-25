package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//runWithTimestamp(D12, 12)
	//runWithTimestamp(D13, 13)
	//runWithTimestamp(D14, 14)
	//runWithTimestamp(D15, 15)
	//runWithTimestamp(D16, 16)
	//runWithTimestamp(D17, 17)
	//runWithTimestamp(D18, 18)
	//runWithTimestamp(D19, 19)
	runWithTimestamp(D20, 20)
	//runWithTimestamp(D21, 21)
	//runWithTimestamp(D22, 22)
	//runWithTimestamp(D23, 23)
	//runWithTimestamp(D24, 24)
	//runWithTimestamp(D25, 25)
}

func removeElement[Item any](arr []Item, i int) []Item {
	arr[0], arr[i] = arr[i], arr[0]
	return arr[1:]
}

func removeElementByValue[Item comparable](arr []Item, value Item) []Item {
	for i, el := range arr {
		if el == value {
			arr = removeElement[Item](arr, i)
			break
		}
	}
	return arr
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sqr(a int) int {
	return a * a
}

func runWithTimestamp[Item string | int](callable func() []Item, day int) {
	from := time.Now()
	res := callable()
	to := time.Now()

	output := make([]string, 0, len(res))
	for _, item := range res {
		output = append(output, fmt.Sprint(item))
	}

	fmt.Printf("Result %d: %s\nTime: %dms\n", day, output, to.UnixMilli()-from.UnixMilli())
}

func parseIntArray(input string, sep string) []int {
	raw := strings.Split(input, sep)

	res := make([]int, 0)
	for _, s := range raw {
		if len(s) == 0 {
			continue
		}
		i, e := strconv.Atoi(s)
		if e != nil {
			fmt.Println(raw, s, e)
			os.Exit(1)
		}
		res = append(res, i)
	}

	return res
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	if a == b {
		return a
	}
	if a < b {
		a, b = b, a
	}
	return gcd(a-b, b)
}
