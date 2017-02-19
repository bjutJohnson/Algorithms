package tree

// defines methods for tree structure

type Search interface {
	Search(key int) *TreeNode
}

// defines extreme values
type Extreme interface {
	Max() *TreeNode
	Min() *TreeNode
}

// insert and delete operation
type Ops interface {
	Insert(insertKey int)
	Delete(deleteKey int) *TreeNode
}

// traverse
type Traverse interface {
	Preorder() []*TreeNode
	Inorder() []*TreeNode
	Postorder() []*TreeNode
}

type LegalCheck interface{
	IsLegal() bool
}
