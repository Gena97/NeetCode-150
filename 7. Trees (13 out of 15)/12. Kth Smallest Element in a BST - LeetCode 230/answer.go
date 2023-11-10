/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Функция kthSmallest находит k-й наименьший элемент в бинарном поисковом дереве.
func kthSmallest(root *TreeNode, k int) int {
	// Создаем срез (массив) arr для хранения узлов дерева в порядке обхода.
	var arr []*TreeNode

	// Начинаем бесконечный цикл.
	for {
		// Обходим все узлы в левом поддереве, добавляя их в срез arr.
		for root != nil {
			arr = append(arr, root)
			root = root.Left
		}

		// Восстанавливаем значение root из последнего узла в срезе arr.
		root = arr[len(arr)-1]
		arr = arr[:len(arr)-1]

		// Уменьшаем k, так как нашли один из наименьших элементов.
		k--

		// Если k достигло нуля, то мы нашли k-й наименьший элемент и возвращаем его значение.
		if k == 0 {
			return root.Val
		}

		// Переходим к правому поддереву для продолжения поиска.
		root = root.Right
	}
}
