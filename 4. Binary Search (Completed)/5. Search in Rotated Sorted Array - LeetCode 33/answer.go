func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {

		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// Проверяем, отсортирована ли левая половина массива
		if nums[left] <= nums[mid] { // Левая половина отсортирована
			// Проверяем, находится ли целевой элемент в отсортированной левой половине
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1 // Сужаем поиск к левой половине
			} else {
				left = mid + 1 // Сужаем поиск к правой половине
			}
		} else { // Правая половина отсортирована
			// Проверяем, находится ли целевой элемент в отсортированной правой половине
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1 // Сужаем поиск к правой половине
			} else {
				right = mid - 1 // Сужаем поиск к левой половине
			}
		}
	}

	return -1
}