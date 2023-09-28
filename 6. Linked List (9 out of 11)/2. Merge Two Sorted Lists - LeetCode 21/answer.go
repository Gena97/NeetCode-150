func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	// Если значение первого узла первого списка меньше значения первого узла
	// второго списка, то объединим оставшиеся части первого списка и второго списка.
	if list1.Val < list2.Val {
		// Рекурсивно вызываем mergeTwoLists для следующего узла первого списка
		// и оставшегося второго списка.
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		// В противном случае объединяем оставшиеся части второго списка и первого списка.
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}