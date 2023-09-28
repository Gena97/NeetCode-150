func characterReplacement(s string, k int) int {
	charCount := make(map[byte]int)
	maxLen := 0
	maxRepeatCount := 0
	start := 0

	// Проходимся по строке s, используя два указателя (start и end) для создания окна.
	for end := 0; end < len(s); end++ {
		charCount[s[end]]++                                     // Увеличиваем счетчик текущего символа.
		maxRepeatCount = max(maxRepeatCount, charCount[s[end]]) // Обновляем максимальное повторение символа.

		// Если текущая длина окна минус максимальное повторение символа больше k,
		// смещаем начало окна вправо, уменьшая счетчик символа в начале окна.
		if end-start+1-maxRepeatCount > k {
			charCount[s[start]]-- // Уменьшаем счетчик символа, который ушел из окна.
			start++               // Сдвигаем начало окна вправо.
		}

		maxLen = max(maxLen, end-start+1)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
