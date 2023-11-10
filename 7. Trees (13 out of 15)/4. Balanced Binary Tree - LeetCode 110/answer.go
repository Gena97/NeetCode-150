/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	// Если корень дерева nil, то считаем его сбалансированным.
	if root == nil {
		return true
	}

	// Вызываем функцию maxDepth для вычисления глубины левого и правого поддеревьев.
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	// Проверяем, соблюдается ли условие сбалансированности для текущего узла.
	if abcCheck(leftDepth, rightDepth) {
		// Если условие соблюдается, рекурсивно проверяем сбалансированность левого и правого поддеревьев.
		return isBalanced(root.Left) && isBalanced(root.Right)
	}

	// Если условие не соблюдается, возвращаем false, так как дерево не сбалансировано.
	return false
}

// maxDepth - функция для вычисления максимальной глубины бинарного дерева.
func maxDepth(root *TreeNode) int {
	// Если корень дерева nil, глубина равна 0.
	if root == nil {
		return 0
	}

	// Рекурсивно вычисляем максимальную глубину, учитывая левое и правое поддеревья.
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// abcCheck - функция для проверки соблюдения условия сбалансированности (разница в глубинах <= 1).
func abcCheck(a, b int) bool {
	c := a - b
	if c >= -1 && c <= 1 {
		return true
	}
	return false
}