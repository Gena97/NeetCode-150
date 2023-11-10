/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Функция goodNodes вызывает функцию abChecker для поиска хороших узлов в дереве,
// начиная с корневого узла.
func goodNodes(root *TreeNode) int {
	// Если дерево пустое, возвращаем 0, так как нет хороших узлов.
	if root == nil {
		return 0
	}

	// Начинаем с корневого узла и считаем его как хороший узел (1).
	// Затем вызываем abChecker для левого и правого поддеревьев и суммируем результаты.
	return 1 + abChecker(root.Left, root.Val) + abChecker(root.Right, root.Val)
}

// Функция abChecker выполняет поиск хороших узлов в поддереве,
// принимая на вход текущий узел (root) и значение, с которым будем сравнивать (v).
func abChecker(root *TreeNode, v int) int {
	// Если узел пустой, возвращаем 0.
	if root == nil {
		return 0
	}

	// Если значение текущего узла больше или равно v, считаем его как хороший узел (1).
	// Затем рекурсивно вызываем abChecker для левого и правого поддеревьев и суммируем результаты.
	if root.Val >= v {
		return 1 + abChecker(root.Left, root.Val) + abChecker(root.Right, root.Val)
	}

	// Если значение текущего узла меньше v, то не считаем его хорошим узлом,
	// и продолжаем поиск хороших узлов в левом и правом поддеревьях.
	return abChecker(root.Left, v) + abChecker(root.Right, v)
}
