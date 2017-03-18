package graph

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

	feature map[string]interface{} // 定义图节点本身的属性，以key-value的形势提供
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

//
