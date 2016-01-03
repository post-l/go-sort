package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	// ArrSize define the size of the set to be sort
	ArrSize = 10000
	// NumberRange define the range of the random numbers in the array
	NumberRange = 1000
)

// QuickSortInt sorts an array of int in an average of O(n log n) with Quicksort Algorithm.
func QuickSortInt(arr []int) {
	if len(arr) <= 5 {
		InsertionSortInt(arr)
	} else {
		p := partition(arr)
		QuickSortInt(arr[:p])
		QuickSortInt(arr[p+1:])
	}
}

func partition(arr []int) int {
	p := choosePivot(arr)
	end := len(arr) - 1
	arr[p], arr[end] = arr[end], arr[p]
	k := 0
	for i := 0; i < end; i++ {
		if arr[i] < arr[end] {
			arr[i], arr[k] = arr[k], arr[i]
			k++
		}
	}
	arr[k], arr[end] = arr[end], arr[k]
	return k
}

func choosePivot(arr []int) int {
	n := len(arr)
	end := n - 1
	mid := n / 2
	if arr[0] < arr[end] {
		if arr[0] >= arr[mid] {
			return 0
		} else if arr[mid] < arr[end] {
			return mid
		}
		return end
	}
	if arr[end] >= arr[mid] {
		return end
	} else if arr[mid] < arr[0] {
		return mid
	}
	return 0
}

// HeapSortInt sorts an array of int in an average of O(n log n) with HeapSort Algorithm
func HeapSortInt(arr []int) {
	heapify(arr)
	for end := len(arr) - 1; end > 0; {
		arr[end], arr[0] = arr[0], arr[end]
		end--
		siftDown(arr, 0, end)
	}
}

func heapify(arr []int) {
	for start, end := (len(arr)-2)/2, len(arr)-1; start >= 0; start-- {
		siftDown(arr, start, end)
	}
}

func siftDown(arr []int, start, end int) {
	for root := start; root*2+1 <= end; {
		child := root*2 + 1
		swap := root
		if arr[swap] < arr[child] {
			swap = child
		}
		if child+1 <= end && arr[swap] < arr[child+1] {
			swap = child + 1
		}
		if swap == root {
			return
		}
		arr[root], arr[swap] = arr[swap], arr[root]
		root = swap
	}
}

// MergeSortInt sorts an array of int in an average of O(n log n) with Merge Sort Algorithm.
// return a new sorted int slice. Not in-place, space complexity O(n).
func MergeSortInt(arr []int) []int {
	n := len(arr)
	if n > 1 {
		mid := n / 2
		a1 := MergeSortInt(arr[:mid])
		a2 := MergeSortInt(arr[mid:])
		res := make([]int, n)
		for i, j, k := 0, 0, 0; k < len(res); k++ {
			if i < len(a1) && (j == len(a2) || a1[i] <= a2[j]) {
				res[k] = a1[i]
				i++
			} else {
				res[k] = a2[j]
				j++
			}
		}
		return res
	}
	return arr
}

// InsertionSortInt sorts an array of int in an average of O(n^2) with Insertion Sort Algorithm.
func InsertionSortInt(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}

// BubbleSortInt sorts an array of int in an average of O(n^2) with an optimize Bubble Sort Algorithm.
func BubbleSortInt(arr []int) {
	for n := len(arr); n != 0; {
		newn := 0
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				newn = i
			}
		}
		n = newn
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	arr := make([]int, ArrSize)
	for i := range arr {
		arr[i] = rand.Intn(NumberRange)
	}
	fmt.Println("The array has", ArrSize, "elements")
	arr2 := duplicateIntSlice(arr)
	t1 := time.Now()
	QuickSortInt(arr2)
	fmt.Println("QuickSort:\t", time.Since(t1))
	arr2 = duplicateIntSlice(arr)
	t1 = time.Now()
	HeapSortInt(arr2)
	fmt.Println("HeapSort:\t", time.Since(t1))
	t1 = time.Now()
	MergeSortInt(arr) // Not in-place so no need to duplicate the array
	fmt.Println("MergeSort:\t", time.Since(t1))
	arr2 = duplicateIntSlice(arr)
	t1 = time.Now()
	InsertionSortInt(arr2)
	fmt.Println("InsertionSort:\t", time.Since(t1))
	arr2 = duplicateIntSlice(arr)
	t1 = time.Now()
	BubbleSortInt(arr2)
	fmt.Println("BubbleSort:\t", time.Since(t1))
}

func duplicateIntSlice(arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	return newArr
}
