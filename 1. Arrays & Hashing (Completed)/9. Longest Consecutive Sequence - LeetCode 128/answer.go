func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Создаем множество для хранения уникальных чисел из входного массива
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxLength := 0

	// Проходимся по числам и находим самую длинную последовательность
	for num := range numSet {
		// Проверяем, является ли текущее число началом последовательности
		if !numSet[num-1] {
			currNum := num
			currLength := 1

			// Находим длину последовательности
			for numSet[currNum+1] {
				currNum++
				currLength++
			}

			// Обновляем максимальную длину, если необходимо
			if currLength > maxLength {
				maxLength = currLength
			}
		}
	}

	return maxLength
}