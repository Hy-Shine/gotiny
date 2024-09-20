package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	list := []int{1, 6, 3, 0, 45, -1, 89, 23}
	BubbleSort(list)
	fmt.Println(list)
}
