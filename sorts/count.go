package sorts

import "fmt"

const length = 255

func CountingSort(words []string) {
	max := getMax(words)
	var res []string
	for i := 0; i < max; i++ {
		res = countSort(words, i)
	}

	fmt.Println("res", res)
}

func getMax(words []string) int {
	max := 0
	for _, v := range words {
		if len(v) > max {
			max = len(v)
		}
	}

	return max
}

func countSort(words []string, index int) []string {
	counterArr := make([]int, length+1)
	resArr := make([]string, len(words))

	for _, word := range words {
		char := '*'
		if index < len(word) {
			char = rune(word[index])
		}
		counterArr[char] += 1
	}

	for i := 1; i < len(counterArr); i++ {
		counterArr[i] += counterArr[i-1]
	}

	for i := len(words) - 1; i >= 0; i-- {
		char := '*'
		if index < len(words[i]) {
			char = rune(words[i][index])
		}
		id := counterArr[char] - 1
		counterArr[char] = id
		resArr[id] = words[i]
	}

	return resArr
}
