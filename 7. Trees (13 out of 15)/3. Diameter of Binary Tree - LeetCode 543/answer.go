/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// maxDiameter - глобальная переменная для хранения максимального диаметра дерева.
var maxDiameter int

// diameterOfBinaryTree - основная функция для вычисления диаметра бинарного дерева.
func diameterOfBinaryTree(root *TreeNode) int {
	// Инициализируем maxDiameter как 0 перед началом вычислений.
	maxDiameter = 0
	// Вызываем DFS для обхода дерева и вычисления диаметра.
	dfs(root, 0)
	// Возвращаем максимальный диаметр.
	return maxDiameter
}

// dfs - рекурсивная функция для обхода дерева и вычисления диаметра.
func dfs(root *TreeNode, depth int) int {
	// Если текущий узел nil, возвращаем 0.
	if root == nil {
		return 0
	}

	// Рекурсивно вызываем dfs для левого и правого поддеревьев.
	leftDepth := dfs(root.Left, depth+1)
	rightDepth := dfs(root.Right, depth+1)

	// Вычисляем диаметр через сумму глубин левого и правого поддеревьев.
	diameter := leftDepth + rightDepth
	// Если текущий диаметр больше максимального, обновляем максимальный диаметр.
	maxDiameter = max(diameter, maxDiameter)

	// Возвращаем максимальную глубину поддерева + 1 (текущий уровень глубины).
	return max(leftDepth, rightDepth) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
