func maxProfit(prices []int) int {
	res := 0
	left, right := 0, 0

	for right < len(prices) {
		// Если текущая цена больше, находящейся слева,
		// то вычисляем прибыль, которую можно получить, продав акции справа.
		if prices[left] < prices[right] {
			res = max(res, prices[right]-prices[left])
		} else {
			// Если текущая цена меньше или равна цене слева,
			// обновляем left, чтобы сдвинуть интервал покупки.
			left = right
		}
		// Увеличиваем right для перехода к следующей цене.
		right += 1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
