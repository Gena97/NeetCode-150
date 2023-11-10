// Функция нахождения меньшего общего предка
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    // Начинаем с корневого узла
    cur := root 

    // Итеративно проходим по дереву
    for cur != nil {
        // Если оба узла p и q находятся в правом поддереве,
        // двигаемся вправо
        if p.Val > cur.Val && q.Val > cur.Val {
            cur = cur.Right
        } 
        // Если оба узла p и q находятся в левом поддереве,
        // двигаемся влево
        else if p.Val < cur.Val && q.Val < cur.Val { 
            cur = cur.Left
        } 
        // Если текущий узел cur находится между узлами p и q
        // или равен одному из них, это означает, что текущий узел - LCA
        else {
            return cur
        }
    }

    // Если не найдено общего предка, возвращаем nil
    return nil
}
