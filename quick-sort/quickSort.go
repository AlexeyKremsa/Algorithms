package main

func quickSort(arr []int) {
	if len(arr) == 0 || len(arr) == 1 {
		return
	}

	lowestIndex := -1
	pivot := len(arr) - 1

	for i, val := range arr {
		if val < arr[pivot] {
			lowestIndex++
			swap(&arr[i], &arr[lowestIndex])
		}
	}

	lowestIndex++
	swap(&arr[pivot], &arr[lowestIndex])

	quickSort(arr[0:lowestIndex])
	quickSort(arr[lowestIndex+1 : len(arr)])
}

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
