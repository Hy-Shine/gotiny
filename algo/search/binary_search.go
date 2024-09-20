package search

import "golang.org/x/exp/constraints"

func BinarySearch[K constraints.Ordered](arr []K, target K) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)>>1
		if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func BinarySearchContains[K constraints.Ordered](arr []K, target K) bool {
	return BinarySearch(arr, target) != -1
}
