package sort

func BubbleSort(list []int) {
	if len(list) <= 1 {
		return
	}
	for i := 0; i < len(list); i++ {
		for j := 1; j < len(list)-i; j++ {
			if list[j-1] > list[j] {
				list[j-1], list[j] = list[j], list[j-1]
			}
		}
	}
}
