/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	return isSameTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Если одно из деревьев nil, а другое - нет, то они не идентичны.
	if p == nil && q != nil || q == nil && p != nil {
		return false
	}

	// Если оба дерева не nil, то сравниваем значения и их поддеревья рекурсивно.
	if p != nil {
		// Если значения узлов не совпадают, деревья не идентичны.
		if p.Val != q.Val {
			return false
		}

		// Рекурсивно вызываем isSameTree для левых и правых поддеревьев.
		if isSameTree(p.Left, q.Left) == false || isSameTree(p.Right, q.Right) == false {
			return false
		}
	}

	// Если ни одно из условий выше не выполнилось, деревья идентичны.
	return true
}