func trap(height []int) int {
	if len(height) < 3 {
		return 0 // Если в "ландшафте" менее трех элементов, вода не может быть заперта, возвращаем 0.
	}

	res := 0

	left, right := 1, len(height)-2 // Указатели для движения слева и справа от ландшафта.
	maxL := height[0]               // Максимальная высота слева, изначально равна первому элементу.
	maxR := height[len(height)-1]   // Максимальная высота справа, изначально равна последнему элементу.

	for left <= right {
		if maxL < maxR {
			// Если текущая высота слева меньше максимальной, то вода может быть заперта.
			if maxL-height[left] > 0 {
				res += maxL - height[left] // Добавляем объем воды, который может быть заперт.
			}
			maxL = max(maxL, height[left]) // Обновляем максимальную высоту слева.
			left++                         // Перемещаемся вправо.
		} else {
			// Аналогично для случая, когда текущая высота справа меньше максимальной.
			if maxR-height[right] > 0 {
				res += maxR - height[right]
			}
			maxR = max(maxR, height[right])
			right-- // Перемещаемся влево.
		}
	}

	return res // Возвращаем общий объем воды, который может быть заперт.
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}