package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/pkg/profile"
)

var isUniq bool

type SortedWords struct {
	arr []string
}

func (sw *SortedWords) Less(i, j int) bool {
	return sw.arr[i] < sw.arr[j]
}

func (sw *SortedWords) Len() int {
	return len(sw.arr)
}

func (sw *SortedWords) Swap(i, j int) {
	sw.arr[i], sw.arr[j] = sw.arr[j], sw.arr[i]
}

func (sw *SortedWords) Add(word string) {
	sw.arr = append(sw.arr, word)
}

func main() {
	defer profile.Start(profile.MemProfileAllocs, profile.ProfilePath(".")).Stop()

	if len(os.Args) < 2 {
		panic("invalid params")
	}

	filename := ""
	for i, v := range os.Args {
		switch v {
		case "-u":
			isUniq = true
		}

		if i == len(os.Args)-1 {
			filename = v
		}
	}

	f := read(filename)

	words := &SortedWords{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		words.Add(s.Text())
	}

	sort.Sort(words)

	uniq := make(map[string]struct{})
	for _, v := range words.arr {
		if isUniq {
			if _, ok := uniq[v]; !ok {
				uniq[v] = struct{}{}
				fmt.Println(v)
			}
		} else {
			fmt.Println(v)
		}
	}
}

func read(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return f
}
