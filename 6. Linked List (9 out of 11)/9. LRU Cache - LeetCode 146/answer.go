// LRUCache представляет кэш с политикой "Наименее недавно используемые".
type LRUCache struct {
	cache    map[int]*list.Element // Карта для хранения ключей и элементов связанного списка
	linklist *list.List            // Связанный список для управления порядком элементов
	capacity int                   // Емкость кэша
}

// Constructor создает и инициализирует новый объект LRUCache с заданной емкостью.
func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    make(map[int]*list.Element, capacity), // Инициализация карты с заданной емкостью
		linklist: list.New(),                            // Инициализация связанного списка
		capacity: capacity,                              // Установка емкости кэша
	}
}

// Get возвращает значение, связанное с заданным ключом, или -1, если ключ не найден.
func (this *LRUCache) Get(key int) int {
	// Проверяем, есть ли ключ в кэше
	if _, ok := this.cache[key]; !ok {
		return -1
	}

	// Если ключ найден, перемещаем элемент в начало связанного списка (обновление LRU)
	elem := this.cache[key]
	this.linklist.MoveToFront(elem)

	// Возвращаем значение из элемента
	return elem.Value.([]int)[1]
}

// Put добавляет значение в кэш с заданным ключом.
func (this *LRUCache) Put(key int, value int) {
	// Проверяем, существует ли ключ в кэше
	if elem, ok := this.cache[key]; ok {
		// Если ключ существует, обновляем его значение и перемещаем элемент в начало списка
		this.linklist.Remove(elem)
		newelem := this.linklist.PushFront([]int{key, value})
		this.cache[key] = newelem
		return
	}

	// Если емкость кэша достигнута, удаляем самый старый элемент из связанного списка и из карты
	if len(this.cache) == this.capacity {
		elem := this.linklist.Back()
		v := this.linklist.Remove(elem)
		delete(this.cache, v.([]int)[0])
	}

	// Создаем новый элемент и добавляем его в начало связанного списка и в карту
	newelem := this.linklist.PushFront([]int{key, value})
	this.cache[key] = newelem
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */