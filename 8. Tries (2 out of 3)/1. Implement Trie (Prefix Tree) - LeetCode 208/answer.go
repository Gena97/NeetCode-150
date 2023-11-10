/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

// Определение структуры узла бора (Trie).
type Node struct {
	Children [26]*Node // Массив дочерних узлов для каждой буквы алфавита (a-z).
	isEnd    bool      // Флаг, указывающий, является ли текущий узел концом слова.
}

// Определение структуры бора (Trie).
type Trie struct {
	root *Node // Корневой узел бора.
}

// Конструктор для инициализации бора (Trie).
func Constructor() Trie {
	result := Trie{root: &Node{}} // Создание нового бора с корневым узлом.
	return result
}

// Вставка слова в бор.
func (this *Trie) Insert(word string) {
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

// Поиск слова в боре.
func (this *Trie) Search(word string) bool {
	wordLength := len(word)
	currentNode := this.root // Начинаем с корневого узла бора.

	// Перебираем буквы в слове и проверяем их наличие в боре.
	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'
		if currentNode.Children[charIndex] == nil {
			return false // Если буква отсутствует, слово не найдено.
		}
		currentNode = currentNode.Children[charIndex] // Переходим к следующему узлу.
	}

	return currentNode.isEnd // Возвращаем true, если последний узел помечен как конец слова.
}

// Поиск префикса (начала слова) в боре.
func (this *Trie) StartsWith(prefix string) bool {
	wordLength := len(prefix)
	currentNode := this.root // Начинаем с корневого узла бора.

	// Перебираем буквы в префиксе и проверяем их наличие в боре.
	for i := 0; i < wordLength; i++ {
		charIndex := prefix[i] - 'a'
		if currentNode.Children[charIndex] == nil {
			return false // Если буква отсутствует, префикс не найден.
		}
		currentNode = currentNode.Children[charIndex] // Переходим к следующему узлу.
	}

	return true // Если все буквы префикса найдены, возвращаем true.
}
