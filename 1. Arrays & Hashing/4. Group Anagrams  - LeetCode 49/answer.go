type Key [26]byte //Для каждого анаграма необходим ключ

func strKey(str string) (key Key) { //Функция перевода строки в ключ
	for i := range str {
		key[str[i]-'a']++ //'-а' нужно, чтобы получилось 0-25
	}
	return
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[Key][]string) //Мапа слайсов для анаграммов по ключам(Key)

	for _, v := range strs {
		key := strKey(v)                     //Получаем ключ анаграмма
		groups[key] = append(groups[key], v) //Добавляем в слайс по ключу
	}

	result := make([][]string, 0, len(groups)) //Для ответа инициализируем слайс слайсов строк с емкостью длины мапы

	for _, v := range groups {
		result = append(result, v)
	}

	return result
}