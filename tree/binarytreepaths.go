/*
 * @Descripttion: 二叉树路径
 * @Author: DW
 * @Date: 2023-12-26 23:09:16
 * @LastEditTime: 2023-12-26 23:19:39
 */
package tree

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	// 定义一个递归函数 travel，用于深度优先搜索并记录路径
	var traverse func(node *TreeNode, s string)
	var str []string
	traverse = func(node *TreeNode, s string) {
		// 如果当前节点是叶子节点，将路径添加到字符串数组中
		if node.Left == nil && node.Right == nil {
			v := s + strconv.Itoa(node.Val)
			str = append(str, v)
			return
		}
		// 将当前节点的值添加到路径中，并继续递归遍历左右子树

		s = s + strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			traverse(node.Left, s)
		}
		if node.Right != nil {
			traverse(node.Right, s)
		}
	}
	// 调用递归函数，初始路径为空字符串

	traverse(root, "")
	return str

}
