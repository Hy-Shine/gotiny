package linkedlist

import (
	"fmt"
	"testing"
)

type node struct {
	data int
	next *node
}

type MyHashSet struct {
	bucket [1000]*node
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func hashFunc(key int) int {
	return key % 1000
}

func (this *MyHashSet) Add(key int) {
	h := hashFunc(key)
	head := this.bucket[h]
	if head == nil {
		this.bucket[h] = &node{data: key}
	} else {
		temp := head
		for temp != nil {
			if temp.data == key {
				return
			}
			if temp.next == nil {
				temp.next = &node{data: key}
				break
			}
			temp = temp.next
		}
	}
}

func (this *MyHashSet) Remove(key int) {
	h := hashFunc(key)
	head := this.bucket[h]
	if head != nil {
		x := &node{next: head}
		temp := x
		for temp != nil {
			if temp.next != nil {
				if temp.next.data == key {
					temp.next = temp.next.next
					break
				}
			}
			temp = temp.next
		}
		this.bucket[h] = x.next
	}
}

func (this *MyHashSet) Contains(key int) bool {
	h := hashFunc(key)
	head := this.bucket[h]
	if head != nil {
		temp := head
		for temp != nil {
			if temp.data == key {
				return true
			}
			temp = temp.next
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

func TestMyHash(t *testing.T) {
	obj := Constructor()
	obj.Add(1)
	obj.Add(1001)
	obj.Add(2)
	obj.Add(3)
	fmt.Println(obj.Contains(1))
	obj.Add(1)
	obj.Remove(1)
	fmt.Println(obj.Contains(1))
}

func TestMoveZero(t *testing.T) {
	f := func(nums []int) {
		slow, fast := 0, 0
		for fast < len(nums) {
			if nums[fast] != 0 {
				nums[slow] = nums[fast]
				slow++
			}
			fast++
		}
	}

	x := []int{1, 2, 6, 5, 0, 2, 3, 0, 89, 6, 0, 56}
	f(x)
	fmt.Println(x)
}
