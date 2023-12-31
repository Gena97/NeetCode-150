// Функция minEatingSpeed вычисляет минимальную скорость поедания (k) бананов, которую Коко должна иметь,
// чтобы съесть все бананы в течение h часов.
func minEatingSpeed(piles []int, h int) int {
	// Найдем максимальное количество бананов в одной из куч.
	maxVal := 0
	for _, pile := range piles {
		maxVal = max(maxVal, pile)
	}

	left, right := 1, maxVal // Инициализация диапазона бинарного поиска.

	// Выполним бинарный поиск, чтобы найти минимальное значение k.
	for left < right {
		mid := left + (right-left)/2 // Середина текущего диапазона.
		hours := 0

		// Подсчитаем общее количество часов, которое потребуется Коко при текущей скорости поедания.
		// Например, если скорость - 3, а бананов - 4, то понадобится 2 часа. Для этого округляем в большую сторону.
		for _, pile := range piles {
			hours += int(math.Ceil(float64(pile) / float64(mid)))
		}

		// Если общее количество часов меньше или равно h, смещаем правую границу диапазона к текущей середине.
		if hours <= h {
			right = mid
		} else { // Иначе, смещаем левую границу диапазона к текущей середине + 1.
			left = mid + 1
		}
	}

	return left // Возвращаем минимальное значение k.
}

// Функция min возвращает минимум из двух чисел.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Функция max возвращает максимум из двух чисел.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
