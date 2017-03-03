package tree

//import "fmt"

// this struct used to manage all tree nodes
type BSTree struct {
	nodes []*TreeNode
	root 	*TreeNode
}

// constructor
func NewBSTree(pNodes []*TreeNode, pRoot *TreeNode) *BSTree{
	ret := new(BSTree)
	ret.nodes = make([]*TreeNode, 0)
	ret.root = pRoot

	ret.nodes = append(ret.nodes, pNodes...)

	return ret
}

// get and set
func (bst BSTree) GetNodes() []*TreeNode{
	return bst.nodes 
}

// Is it a legal binary search tree??
// it requires: satisfy the features of binary search tree
func (bst BSTree) IsLegal() bool {
	if len(bst.nodes) == 0 {
		return true
	}

	for _, v := range bst.nodes {
		if !v.IsBST() {
			return false
		}
	}

	return true
}

// traverse
func (bst BSTree) Preorder() []*TreeNode{
	if bst.root == nil{
		return nil
	}

	ret := make([]*TreeNode, 0)
	ret = append(ret, bst.root)

	savedRoot := bst.root

	left := bst.root.left
	
	if left != nil{
		bst.root = left
		leftNodes := bst.Preorder()

		ret = append(ret, leftNodes...)
	}

	bst.root = savedRoot

	right := bst.root.right

	if right != nil{
		bst.root = right
		rightNodes := bst.Preorder()

		ret = append(ret, rightNodes...)
	}

	bst.root = savedRoot

	return ret
}

func (bst BSTree) Inorder() []*TreeNode{
	if bst.root == nil{
		return nil
	}

	ret := make([]*TreeNode, 0)
	savedRoot := bst.root

	left := bst.root.GetLeftNode()
	if left != nil{
		bst.root = left
		leftNodes := bst.Inorder()

		ret = append(ret, leftNodes...)
	}

	bst.root = savedRoot
	ret = append(ret, bst.root)

	right := savedRoot.GetRightNode()
	if right != nil{
		bst.root = right
		rightNodes := bst.Inorder()

		ret = append(ret, rightNodes...)
	}
	bst.root = savedRoot

	return ret
}

func (bst BSTree) Postorder() []*TreeNode{
	if bst.root == nil{
		return nil
	}

	ret := make([]*TreeNode, 0)
	savedRoot := bst.root

	left := bst.root.GetLeftNode()
	right := bst.root.GetRightNode()

	if left == nil && right == nil{
		ret = append(ret, bst.root)
	}else{
		if left != nil{
			bst.root = left
			leftNodes := bst.Postorder()
			ret = append(ret, leftNodes...)
		}
	
		bst.root = savedRoot
	
		if right != nil{
			bst.root = right
			rightNodes := bst.Postorder()
			ret = append(ret, rightNodes...)
		}

		bst.root = savedRoot

		ret = append(ret, bst.root)
	}

	return ret
}

// search 
func (bst BSTree) Search(key int) *TreeNode{
	if key == bst.root.GetKey(){
		return bst.root
	}else if key < bst.root.GetKey(){
		bst.root = bst.root.GetLeftNode()
		if bst.root == nil{
			return nil
		}else{
			return bst.Search(key)
		}		
	}else{
		bst.root = bst.root.GetRightNode()
		if bst.root == nil{
			return nil
		}else{
			return bst.Search(key)
		}		
	}

	return nil
}

// max
func (bst BSTree) Max() *TreeNode{
	if bst.root == nil{
		return nil
	}

	right := bst.root.GetRightNode()
	if right == nil{
		return bst.root
	}else{
		bst.root = right
		return bst.Max()
	}
}

// min
func (bst BSTree) Min() *TreeNode{
	if bst.root == nil{
		return nil
	}

	left := bst.root.GetLeftNode()
	if left == nil{
		return bst.root
	}else{
		bst.root = left
		return bst.Min()
	}
}

// insert
func (pBst *BSTree) Insert(insertKey int){
	insertNode := NewTreeNode(insertKey, nil, nil, nil)

	if pBst.root == nil{
		pBst.root = insertNode
	}else{
		cursor := pBst.root
		prevCursor := pBst.root
	
		for cursor != nil{
			prevCursor = cursor
			if insertKey < cursor.GetKey() {
				cursor = cursor.GetLeftNode()
			}else{
				cursor = cursor.GetRightNode()
			}
		}
	
		if insertKey < prevCursor.GetKey(){
			prevCursor.SetLeftNode(insertNode)
		}else{
			prevCursor.SetRightNode(insertNode)
		}
	}
	pBst.nodes = append(pBst.nodes, insertNode)
}

// delete
func (pBst *BSTree) Delete(deleteKey int) *TreeNode{
	pDelNode := pBst.Search(deleteKey)
	if pDelNode == nil{
		return nil
	}

	left := pDelNode.GetLeftNode()
	right := pDelNode.GetRightNode()
	parent := pDelNode.GetParentNode()

	var replace *TreeNode

	savedRoot := pBst.root

	if left == nil && right == nil{
		replace = nil
	}else if left == nil {
		replace = right
	}else if right == nil{
		replace = left
	}else{
		pBst.root = left
		leftMax := pBst.Max()

		if leftMax == left{		
			replace = left
		}else{
			leftMaxLeft := leftMax.GetLeftNode()
			if leftMaxLeft != nil{
				leftMaxParent := leftMax.GetParentNode()
				leftMaxParent.SetRightNode(leftMaxLeft)
			}
			replace = leftMax
		}
	}

	// not delete root
	if parent != nil{
		if parent.GetLeftNode() == pDelNode{
			parent.SetLeftNode(replace)
		}else{
			parent.SetRightNode(replace)
		}
		pBst.root = savedRoot
	}else{ // delete root
		pBst.root = replace
	}

	if replace == left{
		replace.SetRightNode(right)
	}else{
		replace.SetLeftNode(left)
	}

	return replace
}

// print bstree perfectly
func (bst BSTree) PrintBstree(){
	// compute hierachy of the tree
	
}