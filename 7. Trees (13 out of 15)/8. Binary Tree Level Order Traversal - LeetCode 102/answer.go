/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	// Если дерево пустое, возвращаем пустой двумерный срез
	if root == nil {
		return [][]int{}
	}

	// Создаем двумерный срез, в котором будем хранить результат
	var out [][]int

	// Создаем срез для текущего уровня узлов и срез для следующего уровня
	var level []*TreeNode
	var nextLevel = []*TreeNode{root}

	// Пока есть узлы для обхода на следующем уровне
	for len(nextLevel) > 0 {
		// Текущий уровень становится следующим уровнем
		level = nextLevel
		// Сбрасываем срез для следующего уровня
		nextLevel = []*TreeNode{}
		// Добавляем новый пустой срез в результат для текущего уровня
		out = append(out, []int{})

		// Обходим узлы на текущем уровне
		for _, node := range level {
			// Добавляем значение текущего узла в текущий уровень результата
			out[len(out)-1] = append(out[len(out)-1], node.Val)

			// Если есть левый потомок, добавляем его в следующий уровень
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			// Если есть правый потомок, добавляем его в следующий уровень
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
	}

	// Возвращаем результат - двумерный срез с уровнями значений узлов
	return out
}
