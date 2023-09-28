func reverseList(head *ListNode) *ListNode {
	// Инициализируем revHead, чтобы он указывал на начальный узел списка.
	revHead := head

	// Если список пуст или состоит из одного элемента, возвращаем его.
	if head == nil || head.Next == nil {
		return head
	}

	// Рекурсивно вызываем функцию для оставшейся части списка.
	revHead = reverseList(head.Next)

	// Устанавливаем связь между текущим узлом и следующим узлом в обратном порядке.
	head.Next.Next = head

	// Обнуляем указатель Next текущего узла, чтобы избежать цикла.
	head.Next = nil

	// Возвращаем revHead, который теперь указывает на новый начальный узел.
	return revHead
}