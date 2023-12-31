func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	// Инициализируем переменную prev (предыдущий узел) как nil,
	// так как начнем с конца списка.
	var prev *ListNode

	// Инициализируем curr (текущий узел) как голову списка.
	curr := head

	// Проходим по списку.
	for curr != nil {
		// Сохраняем ссылку на следующий узел, чтобы не потерять его.
		nxt := curr.Next

		// Разворачиваем указатель текущего узла, чтобы он указывал на предыдущий узел.
		curr.Next = prev

		// Перемещаем указатели prev и curr на следующий шаг.
		prev = curr
		curr = nxt
	}

	// Возвращаем новую голову списка, которая теперь является последним элементом
	// исходного списка (после его переворота).
	return prev
}