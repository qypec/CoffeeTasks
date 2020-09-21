package main

// description
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	listSize := 0
	nthList := &head

	for l := head; l != nil; l = l.Next {
		if listSize >= n {
			nthList = &((*nthList).Next)
		}
		listSize++
	}
	// remove
	*nthList = (*nthList).Next
	return head
}
