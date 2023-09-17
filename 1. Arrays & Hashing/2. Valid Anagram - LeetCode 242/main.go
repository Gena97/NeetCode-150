func isAnagram(s string, t string) bool {
	sRune := []rune(s) //конвертируем в руны
	tRune := []rune(t)
	runeMap := make(map[rune]int) //создаем мапу для отслеживания количества встречающихся символов

	for _, v := range sRune { //добавляем по 1 за каждую руну первого
		runeMap[v] += 1
	}

	for _, v := range tRune { //вычитаем по 1 за каждую руну второго слова
		runeMap[v] -= 1
	}

	for _, v := range runeMap { //если встречается число в мапе отличное от 0, то возвращаем false
		if v != 0 {
			return false
		}
	}

	return true
}