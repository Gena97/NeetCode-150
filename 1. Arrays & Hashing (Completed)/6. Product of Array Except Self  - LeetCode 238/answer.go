func productExceptSelf(nums []int) []int {
	// Создаем результирующий массив с тем же размером, что и входной массив nums.
	res := make([]int, len(nums), len(nums))

	// Инициализируем первый элемент результата как 1, так как умножение на 1 не изменяет значение.
	res[0] = 1

	// Вычисляем произведение всех элементов слева от текущего элемента и сохраняем в res.
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	// Инициализируем переменную postfix как 1 для хранения произведения элементов справа от текущего элемента.
	postfix := 1

	// Запускаем обратный цикл для вычисления произведения элементов справа и умножения его на res[i-1].
	for i := len(nums) - 1; i > 0; i-- {
		postfix *= nums[i]
		res[i-1] *= postfix
	}

	// Возвращаем результирующий массив, который содержит ответ.
	// Т.к. мы проделывали операции в результурующем массиве, дополнительная память нам не понадобилась.
	return res
}