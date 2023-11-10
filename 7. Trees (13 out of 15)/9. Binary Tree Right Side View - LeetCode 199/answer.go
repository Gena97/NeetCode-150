/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	// Если дерево пустое, возвращаем пустой срез
	if root == nil {
		return []int{}
	}

	// Создаем срез для хранения результатов
	var res []int

	// Создаем срез для текущего уровня узлов и срез для следующего уровня
	var level []*TreeNode
	var nextLevel = []*TreeNode{root}

	// Пока есть узлы на следующем уровне
	for len(nextLevel) > 0 {
		// Текущий уровень становится следующим уровнем
		level = nextLevel
		// Сбрасываем срез для следующего уровня
		nextLevel = []*TreeNode{}

		// Обходим узлы на текущем уровне
		for _, node := range level {
			// Если есть левый потомок, добавляем его в следующий уровень
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			// Если есть правый потомок, добавляем его в следующий уровень
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		// Добавляем значение последнего узла на текущем уровне (с правой стороны) в результат
		res = append(res, level[len(level)-1].Val)
	}

	// Возвращаем результат - срез значений узлов, видимых с правой стороны
	return res
}