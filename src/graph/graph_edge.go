package graph

// 定义图上的边
type GraphEdge struct {
	id       int                    // 唯一标识符
	from     int                    // 出节点
	to       int                    // 入节点
	features map[string]interface{} // 边上的各种属性，暂时可能用不着
}

// 创建一条边
func NewGraphEdge(idx, fromId, toId int) *GraphEdge {
	return &GraphEdge{idx, fromId, toId, nil}
}

func (ge GraphEdge) GetDestination() int {
	return ge.to
}
