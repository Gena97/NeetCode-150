func Encode(stringsToEncode []string) string {
	encodedStrings := make([]string, len(stringsToEncode))

	for i, str := range stringsToEncode {
		// Вычислить длину текущей строки.
		strLen := len(str)

		// Добавить длину, затем '#' и саму строку.
		encodedStrings[i] = strconv.Itoa(strLen) + "#" + str
	}

	// Объединить все закодированные строки с разделителем.
	return strings.Join(encodedStrings, "")
}

func Decode(encodedString string) []string {
	decodedStrings := []string{}

	for len(encodedString) > 0 {
		// Найти позицию символа '#'.
		hashPos := strings.Index(encodedString, "#")

		if hashPos == -1 {
			// Если '#' не найден, кодировка неверна.
			break
		}

		// Распарсить длину строки.
		strLen, err := strconv.Atoi(encodedString[:hashPos])

		if err != nil {
			// Ошибка при парсинге длины.
			break
		}

		// Извлечь закодированную строку на основе ее длины.
		if hashPos+strLen+1 <= len(encodedString) {
			decodedStrings = append(decodedStrings, encodedString[hashPos+1:hashPos+1+strLen])
			encodedString = encodedString[hashPos+1+strLen:]
		} else {
			// Закодированная строка неполная.
			break
		}
	}

	return decodedStrings
}