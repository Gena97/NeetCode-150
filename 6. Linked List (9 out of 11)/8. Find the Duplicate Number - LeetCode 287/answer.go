func findDuplicate(nums []int) int {
	// Создаем карту (map) для отслеживания уникальных чисел, которые мы уже встретили.
	m := make(map[int]bool)

	// Проходим по всем числам в срезе.
	for _, v := range nums {
		// Проверяем, есть ли текущее число v в карте m.
		if _, ok := m[v]; !ok {
			// Если число v не найдено в карте, добавляем его в карту как уже встреченное.
			m[v] = true
		} else {
			// Если число v уже найдено в карте, это означает, что оно повторяется, и мы его возвращаем.
			return v
		}
	}

	// Если в срезе нет повторяющихся чисел, возвращаем первое число в срезе (это предположительно допустимо в данной задаче).
	return nums[0]
}