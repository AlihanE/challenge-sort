package sorts

func QuickSort(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	}

	quickSort(arr, 0, len(arr)-1)

	return arr
}

func partition(arr []string, low, high int) int {
	last := high

	buf := low - 1
	for i := low; i <= high-1; i++ {
		if arr[last] > arr[i] {
			buf++
			arr[buf], arr[i] = arr[i], arr[buf]
		}
	}

	buf++
	arr[last], arr[buf] = arr[buf], arr[last]

	return buf
}

func quickSort(arr []string, low, high int) {
	if low < high {
		p := partition(arr, low, high)

		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}
