func searchMatrix(matrix [][]int, target int) bool {
	// Итерируемся по строкам матрицы.
	for _, v := range matrix {
		left := 0
		right := len(v) - 1
		var mid int
		// Запускаем бинарный поиск в текущей строке.
		for left <= right {
			mid = left + (right-left)/2
			switch {
			// Если средний элемент больше целевого значения, сужаем диапазон справа.
			case v[mid] > target:
				right = mid - 1
			// Если средний элемент меньше целевого значения, сужаем диапазон слева.
			case v[mid] < target:
				left = mid + 1
			// Если средний элемент равен целевому значению, возвращаем true, так как нашли цель.
			default:
				return true
			}
		}
	}
	// Если дошли до этой точки, значит, элемент не найден в матрице, возвращаем false.
	return false
}