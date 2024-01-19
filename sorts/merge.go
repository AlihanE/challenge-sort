package sorts

func MergeSort(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2
	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])
	res := merge(left, right)
	return res
}

func merge(left, right []string) []string {
	res := make([]string, len(left)+len(right))
	i, j := 0, 0
	for k := 0; k < len(res); k++ {
		if i >= len(left) {
			res[k] = right[j]
			j++
		} else if j >= len(right) {
			res[k] = left[i]
			i++
		} else if left[i] < right[j] {
			res[k] = left[i]
			i++
		} else {
			res[k] = right[j]
			j++
		}
	}

	return res
}
