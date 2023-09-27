func containsDuplicate(nums []int) bool {

	m := make(map[int]int)

	for _, v := range nums { //Проходим по всему массиву чисел
		m[v] = m[v] + 1 //Добавляем каждый раз 1, когда число встречается
		if m[v] == 2 {  //Если встретилось 2 раза, возвращаем true
			return true
		}
	}
	return false
}