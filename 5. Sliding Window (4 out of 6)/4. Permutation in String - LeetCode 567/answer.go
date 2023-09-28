func checkInclusion(s1 string, s2 string) bool {
	// Если длина s1 больше длины s2, то s1 не может быть включена в s2.
	if len(s1) > len(s2) {
		return false
	}

	// Создаем массивы для подсчета частоты символов в s1 и первых len(s1) символах s2.
	arrs1 := [26]int{} // 26 - количество маленьких латинских букв 'a'-'z'.
	for i := range s1 {
		arrs1[s1[i]-'a']++
	}

	arrs2 := [26]int{}
	for i := 0; i < len(s1); i++ {
		arrs2[s2[i]-'a']++
	}

	// Если частоты символов в s1 и первых len(s1) символах s2 совпадают, то s1 может быть включена в s2.
	if arrs1 == arrs2 {
		return true
	}

	// Проходим по остальной части s2 и обновляем частоту символов в окне размером len(s1).
	for i := len(s1); i < len(s2); i++ {
		arrs2[s2[i]-'a']++         // Увеличиваем частоту символа, добавленного в окно.
		arrs2[s2[i-len(s1)]-'a']-- // Уменьшаем частоту символа, ушедшего из окна.

		// Если частоты символов в текущем окне совпадают с частотами символов в s1, то s1 может быть включена в s2.
		if arrs1 == arrs2 {
			return true
		}
	}

	return false
}
