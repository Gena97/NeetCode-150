func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	dif(candidates, []int{}, 0, target, &res)
	return res
}

func dif(candidates, cur []int, curN, target int, res *[][]int) {
	if target == 0 {
		temp := make([]int, len(cur))
		copy(temp, cur)
		*res = append(*res, temp)
		return
	}

	if target < 0 || curN == len(candidates) {
		return
	}

	// Вызываем dif с текущим элементом
	dif(candidates, append(cur, candidates[curN]), curN, target-candidates[curN], res)

	// Вызываем dif без текущего элемента, двигаясь к следующему
	dif(candidates, cur, curN+1, target, res)
}