package graph

import (
	"errors"
	"johnson_utility"
)

// 定义三种颜色
const (
	CON_WHITE = 0
	CON_GRAY  = 1
	CON_BLACK = 2
)

// 定义图节点文件
type GraphNode struct {
	id    int   // 节点唯一标识
	edges []int // 标识邻接边的编号
	color int8  // 标识是否访问

	feature map[string]interface{} // 定义图节点本身的属性，以key-value的形式提供，用户负责转换，另一种方案是使用reflect
}

// 创建一个节点
func NewGraphNode(id int) *GraphNode {
	return &GraphNode{id, make([]int, 0), CON_WHITE, nil}
}

// 获取节点唯一标识符
func (node GraphNode) GetId() int {
	return node.id
}

// 加入邻居节点
func (pGNode *GraphNode) AddEdge(edgeIdx int) {
	if len(pGNode.edges) == 0 {
		pGNode.edges = append(pGNode.edges, edgeIdx)
		return
	}

	for _, v := range pGNode.edges {
		if v == edgeIdx {
			return
		}
	}

	pGNode.edges = append(pGNode.edges, edgeIdx)
}

// 获取所有的邻居节点
func (GNode GraphNode) GetEdges() []int {
	return GNode.edges
}

// 获取颜色 -- 代表是否访问
func (GNode GraphNode) GetColor() int8 {
	return GNode.color
}

// 设置颜色
func (pGNode *GraphNode) SetColor(iColor int8) {
	pGNode.color = iColor
}

// 增加属性, key必须唯一
func (pGNode *GraphNode) AddFeature(key string, value interface{}) error {
	if pGNode.feature == nil {
		pGNode.feature = make(map[string]interface{})
	}

	if _, ok := pGNode.feature[key]; !ok {
		pGNode.feature[key] = value
	} else {
		str := johnson_utility.ConcateString("属性", key, "已存在，没有必要添加第二次")
		return errors.New(str)
	}

	return nil
}

// 设置属性的值，要求属性必须存在
func (pGNode *GraphNode) SetFeature(key string, value interface{}) error {
	if pGNode.feature == nil {
		return errors.New("属性对值不存在")
	}

	if _, ok := pGNode.feature[key]; !ok {
		str := johnson_utility.ConcateString("属性", key, "不存在，请先添加属性，然后再进行设置")
		return errors.New(str)
	} else {
		pGNode.feature[key] = value
	}

	return nil
}

// 获取属性值
func (gnode GraphNode) GetFeature(key string) (interface{}, error) {
	if gnode.feature == nil {
		return nil, errors.New("属性对值不存在")
	}

	if _, ok := gnode.feature[key]; !ok {
		str := johnson_utility.ConcateString("属性", key, "不存在，请先添加属性，然后再进行设置")
		return nil, errors.New(str)
	}

	return gnode.feature[key], nil

}
