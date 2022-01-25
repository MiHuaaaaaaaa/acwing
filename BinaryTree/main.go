package main

import "fmt"

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func (t *TreeNode) print() {
	fmt.Println(t.val)
}

// 前序遍历
// 1.访问根节点
// 2.前序遍历左子树
// 3.前序遍历右子树
func (t *TreeNode) PreTraverse() {
	if t == nil {
		return
	}

	t.print()
	t.left.PreTraverse()
	t.right.PreTraverse()

}

// 后序遍历

// 1.后序遍历左子树
// 2.后序遍历右子树
// 3.访问根节点
func (t *TreeNode) PostTraverse() {
	if t == nil {
		return
	}

	t.left.PostTraverse()
	t.right.PostTraverse()
	t.print()

}

// 中序遍历

// 1.中序遍历左子树
// 2.访问根节点
// 3.中序遍历右子树
func (t *TreeNode) MidTraverse() {
	if t == nil {
		return
	}

	t.left.MidTraverse()
	t.print()
	t.right.MidTraverse()

}

func main() {
	t := TreeNode{
		val: 666,
		left: &TreeNode{
			val:   1,
			left:  &TreeNode{val: 0},
			right: &TreeNode{val: 888},
		},
		right: &TreeNode{val: 5},
	}
	t.MidTraverse()
}
