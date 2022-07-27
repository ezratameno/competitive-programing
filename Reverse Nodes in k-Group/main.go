package main

import (
	"fmt"
)

func main() {
	head := &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 5}}}}}
	fmt.Println(reverseKGroup(head, 2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var cur = head
	var values []int

	if head == nil {
		return nil
	}
	// get all the values to slice
	for cur != nil {
		values = append(values, cur.Val)
		cur = cur.Next
	}
	groups := cutSlice(values, k)
	var sortedValues []int
	// reverse each group
	for i := range groups {
		// reverse only groups that are in size k
		if len(groups[i]) == k {
			sortedValues = append(sortedValues, Reverse(groups[i])...)
		} else {
			sortedValues = append(sortedValues, groups[i]...)
		}
	}
	// create the new list
	newHead := &ListNode{Val: sortedValues[0]}
	cur = newHead
	for _, val := range sortedValues[1:] {
		cur.Next = &ListNode{Val: val}
		cur = cur.Next

	}

	return newHead
}

// taken from https://golangprojectstructure.com/reversing-go-slice-array/#:~:text=%20Reversing%20a%20Slice%20%201%20Naive%20Approach.,millions...%204%20Nice%20and%20Concise.%20%20More%20
func Reverse(input []int) []int {
	var output []int
	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}
	return output
}

//cutSlice cut slice into groups of size k.
func cutSlice(slice []int, k int) [][]int {
	var res [][]int
	var start int
	for start+k < len(slice) {
		res = append(res, slice[start:start+k])
		start += k
	}
	// add all the left overs that don't divide to k
	if start < len(slice) {
		res = append(res, slice[start:])
	}
	return res
}
