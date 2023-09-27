func permute(nums []int) [][]int {
	res := [][]int{}
	permutation(nums, 0, &res)
	return res
}

func permutation(nums []int, index int, result *[][]int) {
	if index == len(nums) {
		resultCopy := make([]int, len(nums))
		copy(resultCopy, nums)
		*result = append(*result, resultCopy)
		return
	}

	for i := index; i < len(nums); i++ {
		// Меняем местами элементы для создания новой перестановки
		nums[index], nums[i] = nums[i], nums[index]
		permutation(nums, index+1, result)
		// Восстанавливаем оригинальный порядок элементов
		nums[index], nums[i] = nums[i], nums[index]
	}
}