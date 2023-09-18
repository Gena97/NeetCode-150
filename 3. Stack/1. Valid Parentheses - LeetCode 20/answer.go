func isValid(s string) bool {
	// Преобразуем строку в срез рун для удобства обработки
	sRune := []rune(s)

	// Вызываем функцию isBalanced для проверки сбалансированности скобок
	success := isBalanced(sRune)

	// Если скобки сбалансированы, возвращаем true, иначе false
	if success == true {
		return true
	} else {
		return false
	}
}

// Функция isOpen проверяет, является ли символ открывающей скобкой
func isOpen(s rune) bool {
	if s == '(' || s == '[' || s == '{' {
		return true
	} else {
		return false
	}
}

// Функция isPair проверяет, являются ли символы a и b парой скобок
func isPair(a rune, b rune) bool {
	if a == '(' && b == ')' || a == '[' && b == ']' || a == '{' && b == '}' {
		return true
	} else {
		return false
	}
}

// Функция isBalanced проверяет сбалансированность скобок в строке
func isBalanced(s []rune) bool {
	var stack []rune

	// Итерируемся по символам в строке
	for i := range s {
		if isOpen(s[i]) == true {
			// Если символ - открывающая скобка, добавляем ее в стек
			stack = append(stack, s[i])
		} else {
			// Если символ - закрывающая скобка
			if len(stack) > 0 && isPair(stack[len(stack)-1], s[i]) {
				// Если на вершине стека находится соответствующая открывающая скобка, удаляем ее из стека
				stack = stack[:len(stack)-1]
			} else {
				// Иначе, если не нашлось соответствующей открывающей скобки, возвращаем false
				return false
			}
		}
	}

	// Если стек пуст в конце итерации, значит, скобки сбалансированы
	if len(stack) == 0 {
		return true
	}

	return false
}