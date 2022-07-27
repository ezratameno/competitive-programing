package main

import "fmt"

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	list3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 6,
		},
	}
	list := &[]*ListNode{list1, list2, list3}
	fmt.Println(mergeKLists(*list))
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	res := &ListNode{}
	var helper []*ListNode

	// will run on the res to add values
	runner := res

	// initilize pointers of each list to the helper slice
	helper = append(helper, lists...)

	for {
		min, index := findMin(helper)
		// if -1 then we have only nil pointers in the slice so we stop
		if index == -1 {
			break
		}
		// add the min to the result
		runner.Next = &ListNode{Val: min}
		runner = runner.Next
		// update the pointer of the list that have the min to the next value
		helper[index] = helper[index].Next

	}

	// ignore the first value (the default 0 that we have)
	res = res.Next
	return res

}

// findMin returns the min and index
// if we return -1 as the index then we have only nil pointers in the slice
func findMin(helper []*ListNode) (int, int) {
	var seenNonNull bool
	var min int
	index := -1
	for i, val := range helper {
		// if its the first time that we see a value that is not nil
		if !seenNonNull && val != nil {
			min = val.Val
			index = i
			seenNonNull = true
		}
		// if we have seen a non nil value and is less than the min -> update min and index
		if seenNonNull && val != nil && min > val.Val {
			min = val.Val
			index = i
		}

	}
	return min, index

}
