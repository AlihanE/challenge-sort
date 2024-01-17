package sorts

import (
	"sort"
)

func DefaultSort(words []string) []string {
	sort.Strings(words)

	return words
}
