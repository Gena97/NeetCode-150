/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

// Определение структуры для хранения словаря слов.
type WordDictionary struct {
	root *Node
}

// Определение структуры узла бора (Trie).
type Node struct {
	Children [26]*Node // Массив дочерних узлов для каждой буквы алфавита (a-z).
	isEnd    bool      // Флаг, указывающий, является ли текущий узел концом слова.
}

// Конструктор для инициализации словаря.
func Constructor() WordDictionary {
	result := WordDictionary{root: &Node{}} // Создание нового словаря с корневым узлом бора.
	return result
}

// Добавление слова в словарь.
func (this *WordDictionary) AddWord(word string) {
	wordLength := len(word)
	currentNode := this.root // Начинаем с корневого узла бора.

	// Перебираем буквы в слове и добавляем их в бор.
	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a' // Преобразуем букву в индекс (0-25).
		if currentNode.Children[charIndex] == nil {
			currentNode.Children[charIndex] = &Node{} // Создаем новый узел, если не существует.
		}
		currentNode = currentNode.Children[charIndex] // Переходим к следующему узлу.
	}
	currentNode.isEnd = true // Помечаем последний узел как конец слова.
}

// Поиск слова в словаре.
func (this *WordDictionary) Search(word string) bool {
	return this.searchRecursive(word, 0, this.root)
}

// Рекурсивная функция для поиска слова в словаре.
func (this *WordDictionary) searchRecursive(word string, index int, node *Node) bool {
	if index == len(word) {
		return node.isEnd // Если достигнут конец слова, возвращаем значение флага isEnd узла.
	}

	char := word[index]

	if char != '.' { // Если текущий символ не равен '.', это обычная буква.
		charIndex := char - 'a'
		if node.Children[charIndex] == nil {
			return false // Если буква отсутствует, слово не найдено.
		}
		return this.searchRecursive(word, index+1, node.Children[charIndex]) // Переходим к следующему узлу.
	}

	// Если текущий символ - '.', это символ-джокер, который соответствует любой букве.
	// Перебираем всех дочерних узлов для текущего символа и рекурсивно ищем слово в них.
	for _, child := range node.Children {
		if child != nil && this.searchRecursive(word, index+1, child) {
			return true // Если слово найдено, возвращаем true.
		}
	}

	return false // Если не удалось найти слово, возвращаем false.
}
