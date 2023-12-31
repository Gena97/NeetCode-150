func twoSum(numbers []int, target int) []int {
	// Используем два указателя: один начинается с начала массива, другой с конца.
	i := 0
	j := len(numbers) - 1

	// Пока левый указатель меньше правого, продолжаем поиск.
	for i < j {
		left := numbers[i]
		right := numbers[j]

		// Если сумма чисел на левом и правом указателях больше цели, уменьшаем правый указатель.
		if left+right > target {
			j--
		} else if left+right < target {
			// Если сумма меньше цели, увеличиваем левый указатель.
			i++
		} else {
			// Если сумма равна цели, возвращаем индексы (плюс один, так как массив 1-индексированный).
			return []int{i + 1, j + 1}
		}
	}

	// Если не найдено подходящих чисел, возвращаем массив с нулями.
	return []int{0}
}