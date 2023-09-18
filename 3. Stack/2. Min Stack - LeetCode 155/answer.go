type MinStack struct {
	stacks []stack
}

type Element struct {
	value int
	min   int
}

// Конструктор для создания объекта MinStack
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	minValue := 0
	if len(this.stacks) == 0 {
		// Если стек пуст, то минимальным значением будет сам элемент
		minValue = val
	} else {
		// В противном случае, минимальное значение - минимум из текущего минимума и нового значения
		minValue = min(val, this.GetMin())
	}

	// Создаем новый Element и добавляем его в стек
	newStack := Element{
		value: val,
		min:   minValue,
	}
	this.stacks = append(this.stacks, newStack)
}

func (this *MinStack) Pop() {
	if len(this.stacks) > 0 {
		// Удаляем верхний элемент стека
		this.stacks = this.stacks[:len(this.stacks)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stacks) == 0 {
		return 0 // Возвращаем 0, если стек пуст
	}
	return this.stacks[len(this.stacks)-1].value
}

// Метод GetMin возвращает минимальное значение в стеке
func (this *MinStack) GetMin() int {
	if len(this.stacks) == 0 {
		return 0 // Возвращаем 0, если стек пуст
	}
	return this.stacks[len(this.stacks)-1].min
}

// Вспомогательная функция для нахождения минимума из двух чисел
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}