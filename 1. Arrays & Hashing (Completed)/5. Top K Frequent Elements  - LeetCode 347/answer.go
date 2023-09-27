func topKFrequent(nums []int, k int) []int {

	// Создаем словарь для подсчета частоты встречаемости чисел.
	dic := map[int]int{}
	for _, v := range nums {
		dic[v]++
	}

	// Создаем обратный словарь, где ключ - это частота, а значение - срез чисел с этой частотой.
	revdic := map[int][]int{}
	for v, count := range dic {
		revdic[count] = append(revdic[count], v)
	}

	// Инициализируем результат как пустой срез.
	result := []int{}

	// Начинаем с самой большой частоты(кол-во элементов в мапе) и добавляем числа в результат, пока не достигнем k элементов.
	for i := len(nums); len(result) != k; i-- {
		for _, n := range revdic[i] {
			if len(result) != k {
				result = append(result, n)
			}
		}
	}

	return result
}