func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int          // Переменная для переноса.
	var sum int            // Переменная для суммы текущих узлов.
	var result *ListNode   // Голова результирующего связанного списка.
	var currNode *ListNode // Текущий узел результирующего связанного списка.

	// Проходим по обоим связанным спискам до тех пор, пока есть узлы для сложения.
	for l1 != nil || l2 != nil {
		sum = carry // Начинаем сумму с текущего значения переноса.

		// Если есть узел в l1, добавляем его значение к сумме и переходим к следующему узлу.
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		// Если есть узел в l2, добавляем его значение к сумме и переходим к следующему узлу.
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10 // Вычисляем перенос для следующей итерации.
		sum %= 10        // Оставляем только последнюю цифру суммы.

		// Создаем новый узел с текущим значением суммы и добавляем его к результирующему списку.
		newNode := &ListNode{
			Val:  sum,
			Next: nil,
		}

		if result == nil {
			result = newNode
			currNode = result
		} else {
			currNode.Next = newNode
			currNode = currNode.Next
		}
	}

	// Если после всех итераций остался перенос, создаем новый узел для него.
	if carry > 0 {
		newNode := &ListNode{
			Val:  carry,
			Next: nil,
		}
		currNode.Next = newNode
	}

	return result // Возвращаем голову результирующего связанного списка.
}
