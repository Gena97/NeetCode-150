func copyRandomList(head *Node) *Node {
	// Создаем карту (map) для отображения старых узлов на новые узлы.
	nodeMap := make(map[*Node]*Node)

	// Инициализируем указатель на текущий старый узел.
	oldNode := head

	// Инициализируем указатели на голову и текущий новый узел.
	var newHead *Node
	var newNodeTail *Node

	// Проходим по старому списку и создаем новый список, одновременно заполняя карту.
	for oldNode != nil {
		// Создаем новый узел и копируем значение из текущего старого узла.
		newNode := &Node{Val: oldNode.Val}

		// Связываем текущий старый узел с соответствующим новым узлом в карте.
		nodeMap[oldNode] = newNode

		// Обновляем указатель на следующий новый узел.
		if newHead == nil {
			newHead = newNode
			newNodeTail = newNode
		} else {
			newNodeTail.Next = newNode
			newNodeTail = newNodeTail.Next
		}

		// Переходим к следующему старому узлу.
		oldNode = oldNode.Next
	}

	// Второй проход по старому списку для установки произвольных указателей новым узлам.
	oldNode = head
	for oldNode != nil {
		// Устанавливаем произвольный указатель новому узлу, используя карту.
		newNode := nodeMap[oldNode]
		newNode.Random = nodeMap[oldNode.Random]

		// Переходим к следующему старому узлу.
		oldNode = oldNode.Next
	}

	// Возвращаем голову нового скопированного списка.
	return newHead
}