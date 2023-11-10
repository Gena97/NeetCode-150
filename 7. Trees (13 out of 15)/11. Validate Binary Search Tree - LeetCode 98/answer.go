/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, nil, nil)
}

// Вспомогательная функция для проверки допустимости BST.
// node - текущий узел, min - минимальное значение, которое этот узел должен превысить, max - максимальное значение, которое этот узел должен быть меньше.
func isValidBSTHelper(node *TreeNode, min *TreeNode, max *TreeNode) bool {
	// Если узел является пустым (nil), то он не нарушает свойство BST, и мы возвращаем true.
	if node == nil {
		return true
	}

	// Проверяем, находится ли значение текущего узла в допустимом диапазоне, заданном min и max.
	// Если значение не находится в этом диапазоне, возвращаем false, так как дерево не является BST.
	if (min != nil && node.Val <= min.Val) || (max != nil && node.Val >= max.Val) {
		return false
	}

	// Рекурсивно вызываем isValidBSTHelper для левого и правого поддеревьев с обновленными min и max.
	// Для левого поддерева устанавливаем новое значение max как текущий узел, а для правого - новое значение min.
	// Это гарантирует, что все узлы в левом поддереве будут меньше текущего узла, а в правом - больше.
	return isValidBSTHelper(node.Left, min, node) && isValidBSTHelper(node.Right, node, max)
}
