func evalRPN(tokens []string) int {
	stack := []int{} // Стек для хранения чисел и промежуточных результатов
	operators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}

	// Итерация по элементам массива tokens
	for _, token := range tokens {
		if operators[token] {
			// Если текущий токен - оператор, выполняем операцию над двумя верхними числами в стеке и добавляем результат обратно в стек
			b := stack[len(stack)-1]       // Последний элемент в стеке
			a := stack[len(stack)-2]       // Предпоследний элемент в стеке
			stack = stack[:len(stack)-2]   // Удаляем использованные числа из стека
			result := operate(token, a, b) // Выполняем операцию
			stack = append(stack, result)  // Добавляем результат обратно в стек
		} else {
			// Если текущий токен - число, преобразуем его из строки в int и добавляем в стек
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	// Результат вычисления находится в вершине стека
	return stack[0]
}

// Функция operate выполняет математические операции (+, -, *, /) над двумя числами a и b в зависимости от оператора.
func operate(operator string, a, b int) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b // Поскольку задача требует, чтобы деление всегда округлялось к нулю
	default:
		return 0 // Если оператор неизвестен, возвращаем 0
	}
}