func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	var mid int

	for left <= right {
		// Вычисляем средний индекс.
		mid = left + (right-left)/2

		switch {
		// Если средний элемент больше целевого значения, сужаем диапазон поиска справа.
		case nums[mid] > target:
			right = mid - 1
		// Если средний элемент меньше целевого значения, сужаем диапазон поиска слева.
		case nums[mid] < target:
			left = mid + 1
		// Если средний элемент равен целевому значению, возвращаем его индекс.
		default:
			return mid
		}
	}

	return -1
}