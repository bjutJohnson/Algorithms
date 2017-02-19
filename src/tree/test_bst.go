package tree

import "fmt"

func Test1(){
	node7 := NewTreeNode(7, nil, nil, nil)
	node3 := NewTreeNode(3, nil, nil, nil)
	node1 := NewTreeNode(1, nil, nil, nil)
	node5 := NewTreeNode(5, nil, nil, nil)
	node12 := NewTreeNode(12, nil, nil, nil)
	node9 := NewTreeNode(9, nil, nil, nil)
	node16 := NewTreeNode(16, nil, nil, nil)

	node7.SetLeftNode(node3)
//	node3.SetParentNode(node7)

	node7.SetRightNode(node12)
//	node12.SetParentNode(node7)

	node3.SetLeftNode(node1)
//	node1.SetParentNode(node3)

	node3.SetRightNode(node5)
//	node5.SetParentNode(node3)
	
	node12.SetLeftNode(node9)
//	node9.SetParentNode(node12)

	node12.SetRightNode(node16)
//	node16.SetParentNode(node12)

	bst := NewBSTree([]*TreeNode{node7, node3, node1, node5, node12, node9, node16}, node7)

	fmt.Println("search 0")
	search0 := bst.Search(0)
	if search0 != nil{
		fmt.Println(*search0)
	}else{
		fmt.Println("0 not exist")
	}

	fmt.Println("search 5")
	search5 := bst.Search(5)
	if search5 != nil{
		fmt.Println(*search5)
	}else{
		fmt.Println("5 not exist")
	}

	fmt.Println("Max:")
	max := bst.Max()
	if max != nil{
		fmt.Println(*max)
	}else{
		fmt.Println("no element in the tree")
	}

	fmt.Println("Min:")
	min := bst.Min()
	if min != nil{
		fmt.Println(*min)
	}else{
		fmt.Println("no element in the tree")
	}

	fmt.Println("insert 8")
	bst.Insert(8)

	fmt.Println("Delete 12")
	replace := bst.Delete(12)
	fmt.Println(*replace)

	retPre := bst.Preorder()
	retIn := bst.Inorder()
	retPost := bst.Postorder()

	fmt.Println("preorder:")
	PrintNodes(retPre)
	fmt.Println("inorder:")
	PrintNodes(retIn)
	fmt.Println("postorder:")
	PrintNodes(retPost)	
}
