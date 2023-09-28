func lengthOfLongestSubstring(s string) int {
	hashmap := make(map[rune]int)
	maxCounter := 0
	start := 0

	for i, v := range s {
		// Если символ v уже встречался ранее в текущей подстроке,
		// и его индекс находится после начала текущей подстроки, обновляем начало подстроки.
		if index, found := hashmap[v]; found && index >= start {
			start = index + 1
		}
		// Записываем индекс символа в хеш-карту.
		hashmap[v] = i
		// Вычисляем текущую длину подстроки без повторов и обновляем максимальную длину, если необходимо.
		if length := i - start + 1; length > maxCounter {
			maxCounter = length
		}
	}
	return maxCounter
}