func isValidSudoku(board [][]byte) bool {
	// Создаем три карты (хэш-таблицы) для отслеживания использованных чисел в каждой строке, колонке и блоке.
	var hashrow = make(map[int]*[9]bool) // Хранит информацию о числах в каждой строке.
	var hashcol = make(map[int]*[9]bool) // Хранит информацию о числах в каждой колонке.
	var hashbox = make(map[int]*[9]bool) // Хранит информацию о числах в каждом 3x3 блоке.

	// Проходим по всей доске 9x9.
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// Если ячейка не пустая (не '.'), начинаем проверку.
			if board[i][j] != '.' {
				// 1. Проверяем, использовалось ли это число в текущей строке.
				value := hashrow[i]
				if value == nil {
					// Если значение отсутствует, инициализируем массив для отслеживания использованных чисел.
					hashrow[i] = &[9]bool{}
					value = hashrow[i]
				}

				// Преобразуем символ в целое число.
				byteToInt, _ := strconv.Atoi(string(board[i][j]))

				// Если число уже использовалось в текущей строке, возвращаем false.
				if value[byteToInt-1] == true {
					return false
				} else {
					// Иначе отмечаем число как использованное.
					value[byteToInt-1] = true
				}

				// 2. Проверяем, использовалось ли это число в текущей колонке (аналогично строке).
				value = hashcol[j]
				if value == nil {
					hashcol[j] = &[9]bool{}
					value = hashcol[j]
				}

				if value[byteToInt-1] == true {
					return false
				} else {
					value[byteToInt-1] = true
				}

				// 3. Вычисляем индекс текущего 3x3 блока.
				boxIndex := (i/3)*3 + j/3

				// Проверяем, использовалось ли это число в текущем 3x3 блоке (аналогично строке и колонке).
				value = hashbox[boxIndex]
				if value == nil {
					hashbox[boxIndex] = &[9]bool{}
					value = hashbox[boxIndex]
				}

				if value[byteToInt-1] == true {
					return false
				} else {
					value[byteToInt-1] = true
				}
			}
		}
	}

	// Если все проверки пройдены успешно, возвращаем true.
	return true
}
