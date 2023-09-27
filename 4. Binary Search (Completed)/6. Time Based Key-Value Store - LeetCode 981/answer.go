type TimeMap struct {
	m map[string][]Value // Хранилище для ключей и значений
}

type Value struct {
	value     string // Значение
	timestamp int    // Отметка времени, когда это значение было установлено
}

// Конструктор для создания нового объекта TimeMap
func Constructor() TimeMap {
	return TimeMap{make(map[string][]Value)}
}

// Метод Set позволяет установить значение для ключа с указанной временной отметкой
func (this *TimeMap) Set(key string, value string, timestamp int) {
	newValue := Value{
		value:     value,
		timestamp: timestamp,
	}
	// Добавляем новое значение в список значений для данного ключа
	this.m[key] = append(this.m[key], newValue)
}

// Метод Get позволяет получить значение ключа для указанной временной отметки
func (this *TimeMap) Get(key string, timestamp int) string {
	var list []Value
	var ok bool
	if list, ok = this.m[key]; !ok {
		return "" // Если ключ не существует, возвращаем пустую строку
	}

	l := 0
	r := len(list)

	// Бинарный поиск в списке значений по временным отметкам
	for l < r {
		m := (l + r) / 2
		if list[m].timestamp > timestamp {
			r = m // Сужаем диапазон поиска к левой половине
		} else {
			l = m + 1 // Сужаем диапазон поиска к правой половине
		}
	}

	if l == 0 {
		return "" // Если не найдено значение для указанной временной отметки, возвращаем пустую строку
	}
	return list[l-1].value // Возвращаем значение, ближайшее к указанной временной отметке
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */