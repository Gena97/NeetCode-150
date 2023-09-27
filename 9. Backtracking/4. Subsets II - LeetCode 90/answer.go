func subsets(nums []int) [][]int {
	res := [][]int{}
	dif(nums, []int{}, 0, &res)
	return res
}

func dif(nums, tempAr []int, curN int, res *[][]int) {
	if curN == len(nums) {
		// Создаем копию tempAr и добавляем ее в res, чтобы избежать изменения уже добавленных подмножеств
		temp := make([]int, len(tempAr))
		copy(temp, tempAr)
		*res = append(*res, temp)
		return
	}

	// Пропускаем текущий элемент nums[curN]
	dif(nums, tempAr, curN+1, res)

	// Включаем текущий элемент nums[curN]
	tempAr = append(tempAr, nums[curN])
	dif(nums, tempAr, curN+1, res)
}
