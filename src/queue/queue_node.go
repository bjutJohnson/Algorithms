package queue

// 队列中的节点
// 建议使用时对待用节点用整型的id唯一标识
type QueueNode struct {
	Value int
}

// 创建队列节点
func NewQueueNode(value int) *QueueNode {
	return &QueueNode{value}
}
