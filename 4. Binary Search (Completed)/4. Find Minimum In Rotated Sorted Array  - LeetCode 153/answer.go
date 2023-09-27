func findMin(nums []int) int {
	len := len(nums)

	if len == 1 {
		return nums[0]
	}

	left, right := 0, len-1

	// Выполняем бинарный поиск до тех пор, пока правая граница меньше левой.
	for nums[right] < nums[left] {

		mid := left + (right-left)/2 // Вычисляем средний индекс.

		// Если элемент в середине больше или равен элементу в начале массива,
		// это означает, что минимум находится справа от mid.
		if nums[mid] >= nums[left] {
			left = mid + 1 // Смещаем левую границу вправо.
		} else {
			right = mid // Иначе, минимум находится слева от mid, включая mid.
		}
	}

	return nums[left] // Возвращаем минимальный элемент.
}
