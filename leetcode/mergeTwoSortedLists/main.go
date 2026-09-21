package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}
	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}
	return dummy.Next
}

func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	var out []int
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	tests := []struct {
		name     string
		list1    []int
		list2    []int
		expected []int
	}{
		{"both empty", nil, nil, nil},
		{"first empty", nil, []int{1, 2, 3}, []int{1, 2, 3}},
		{"second empty", []int{1, 2, 3}, nil, []int{1, 2, 3}},
		{"equal length", []int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"first shorter", []int{1}, []int{2, 3, 4}, []int{1, 2, 3, 4}},
		{"second shorter", []int{1, 2, 3}, []int{4}, []int{1, 2, 3, 4}},
		{"with duplicates", []int{1, 1, 2}, []int{1, 3, 3}, []int{1, 1, 1, 2, 3, 3}},
		{"negative values", []int{-5, -1, 3}, []int{-2, 0, 4}, []int{-5, -2, -1, 0, 3, 4}},
		{"all from first", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"all from second", []int{4, 5, 6}, []int{1, 2, 3}, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		got := listToSlice(mergeTwoLists(buildList(tt.list1), buildList(tt.list2)))
		if !equal(got, tt.expected) {
			fmt.Printf("FAIL [%s]: got %v, want %v\n", tt.name, got, tt.expected)
		} else {
			fmt.Printf("OK   [%s]: %v\n", tt.name, got)
		}
	}
}
