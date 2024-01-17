package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AlihanE/challenge-sort/sorts"
	"github.com/pkg/profile"
)

var isUniq bool

func main() {
	defer profile.Start(profile.MemProfileAllocs, profile.ProfilePath(".")).Stop()

	if len(os.Args) < 2 {
		panic("invalid params")
	}

	filename := ""
	var sorterFunc func([]string) []string
	for i, v := range os.Args {
		switch v {
		case "-u":
			isUniq = true
		case "-count":
			sorterFunc = sorts.CountingSort
		default:
			if sorterFunc == nil {
				sorterFunc = sorts.DefaultSort
			}
		}

		if i == len(os.Args)-1 {
			filename = v
		}
	}

	f := read(filename)

	words := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		words = append(words, s.Text())
	}

	uniq := make(map[string]struct{})
	for _, v := range sorterFunc(words) {
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
