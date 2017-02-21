package tree

import (
	"fmt"
)

// define tree node element
type TreeNode struct {
	parent *TreeNode
	left   *TreeNode
	right  *TreeNode

	key int // key element, according tree by this element
}

// constructor
func NewTreeNode(key int, parent, left, right *TreeNode) *TreeNode {
	ret := new(TreeNode)

	ret.left = left
	ret.right = right
	ret.parent = parent
	ret.key = key

	return ret
}

// get and set tree node
func (pNode *TreeNode) SetLeftNode(left *TreeNode) {
	pNode.left = left

	if left != nil{
		left.parent = pNode
	}
}

func (node TreeNode) GetLeftNode() *TreeNode{
	return node.left 
}

func (pNode *TreeNode) SetRightNode(right *TreeNode) {
	pNode.right = right

	if right != nil{
		right.parent = pNode
	}
}

func (node TreeNode) GetRightNode() *TreeNode{
	return node.right
}

func (pNode *TreeNode) SetParentNode(parent *TreeNode) {
	pNode.parent = parent
}

func (node TreeNode) GetParentNode() *TreeNode{
	return node.parent
}

// get key
func (node TreeNode) GetKey() int {
	return node.key
}

// check features which must be satisfied by bst
func (node TreeNode) IsBST() bool {
	parent := node.parent
	if parent != nil {
		if parent.left == &node && node.key > parent.key {
			return false
		}

		if parent.right == &node && node.key < parent.key {
			return false
		}
	}

	left := node.left
	if left != nil && left.key > node.key {
		return false
	}

	right := node.right
	if right != nil && right.key < node.key {
		return false
	}

	return true
}

// print tree node one by one
func PrintNodes(pNodes []*TreeNode){
	if len(pNodes) != 0{
		for _, v := range(pNodes){
			fmt.Println(v.GetKey())
		}
	}
}