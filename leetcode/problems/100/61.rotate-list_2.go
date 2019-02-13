package problems

//todo Memory Limit Exceeded
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if k <= 0 {
		return head
	}

	end := head

	sum := 1
	for end.Next != nil {
		end = end.Next
		sum++
	}

	//set as ring
	end.Next = head
	k = k % sum
	if k == 0 {
		return head
	}

	for sum-k > 1 {
		head = head.Next
		k++
	}
	r := head.Next
	head.Next = nil

	return r
}
