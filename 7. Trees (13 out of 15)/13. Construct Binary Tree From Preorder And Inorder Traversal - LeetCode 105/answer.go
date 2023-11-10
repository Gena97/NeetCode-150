/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Функция buildTree принимает два среза - preorder и inorder,
// и восстанавливает бинарное дерево на основе данных срезов.
func buildTree(preorder []int, inorder []int) *TreeNode {
	// Если срезы пусты, возвращаем nil, так как нет узлов для построения дерева.
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// Создаем корень дерева, устанавливая его значение из первого элемента preorder.
	root := &TreeNode{
		Val: preorder[0],
	}

	// Находим индекс корневого элемента в inorder.
	var mid int
	for i, v := range inorder {
		if v == root.Val {
			mid = i
			break
		}
	}

	// Разделяем inorder на левую и правую части от корневого элемента.
	inLeft := inorder[:mid]
	inRight := inorder[mid+1:]

	// Разделяем preorder на левую и правую части на основе размера левой части inorder.
	preLeft := preorder[1 : len(inLeft)+1]
	preRight := preorder[len(inLeft)+1:]

	// Рекурсивно строим левое и правое поддеревья и присваиваем их корню.
	root.Left = buildTree(preLeft, inLeft)
	root.Right = buildTree(preRight, inRight)

	// Возвращаем корень, который представляет собой корневой узел восстановленного дерева.
	return root
}
