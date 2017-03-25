package graph

import (
	"errors"
	"fmt"
	"johnson_utility"
	"log"
	"queue"
	"sort"
	"strconv"
	"sync"
)

const (
	QUEUE_SIZE = 100
)

// 管理图的所有节点
type GraphManager struct {
	id2Node map[int]*GraphNode // 以map的形式记录所有的节点
	id2Edge map[int]*GraphEdge // 以map的形式记录所有的边

	nodeCounter int // 计数器，申请的图节点数量
	edgeCounter int // 边计数器

	isDirection bool // 有向图还是无向图，必须初始时给定

	isCirculation bool // 是否存在循环，初始时为false

	mutex *sync.Mutex // 控制对管理对象的操作，一个时刻只能有一个goroutine对其进行修改

	logicTime int // 用于控制深度遍历的时刻的逻辑时钟，只表示先后关系，如t0=0, t1=1表示t1是紧邻t0的下一个时刻
}

// 创建图的管理者
func NewGraphManager(direction bool) *GraphManager {
	return &GraphManager{make(map[int]*GraphNode, 0), make(map[int]*GraphEdge), 0, 0, direction, false, new(sync.Mutex), 0}
}

// 打印所有的边
func (gm GraphManager) PrintEdges() {
	gm.applyChan()
	defer gm.releaseChan()

	log.Println("<=====start : id2Edge=====>")
	for _, v := range gm.id2Edge {
		log.Println(*v)
	}
	log.Println("<=====end :id2Edge=====>")
}

// 打印图节点
func (gm GraphManager) PrintNode() {
	gm.applyChan()
	defer gm.releaseChan()

	fmt.Println("<====starting to print graph node======>")

	for _, v := range gm.id2Node {
		discovery, _ := v.GetFeature("discovery")
		finish, _ := v.GetFeature("finish")

		fmt.Print("id: ")
		fmt.Print(v.GetId())
		fmt.Print(", discovery time: ", discovery)
		fmt.Println(", finishing time: ", finish)
	}
	fmt.Println("<====end of printing graph node =======>")
}

// 增加一个图节点, 返回标识唯一节点的编号
func (pGManager *GraphManager) AddNode() int {
	pGManager.applyChan()
	defer pGManager.releaseChan()

	newNode := NewGraphNode(pGManager.nodeCounter)
	pGManager.id2Node[pGManager.nodeCounter] = newNode

	pGManager.nodeCounter++

	return newNode.GetId()
}

// 加锁
func (pGManager *GraphManager) applyChan() {
	pGManager.mutex.Lock()
}

// 释放锁
func (pGManager *GraphManager) releaseChan() {
	pGManager.mutex.Unlock()
}

// 增加一条边， 当节点不存在时，返回错误
func (pGManager *GraphManager) AddEdge(from, to int) error {
	pGManager.applyChan()
	defer pGManager.releaseChan()

	if _, ok := pGManager.id2Node[from]; !ok {
		errStr := johnson_utility.ConcateString("不存在的from节点号:", strconv.FormatInt(int64(from), 10))
		return errors.New(errStr)
	}

	if _, ok := pGManager.id2Node[to]; !ok {
		errStr := johnson_utility.ConcateString("不存在的to节点号:", strconv.FormatInt(int64(to), 10))
		return errors.New(errStr)
	}

	newEdge := NewGraphEdge(pGManager.edgeCounter, from, to)
	pGManager.id2Edge[pGManager.edgeCounter] = newEdge
	pGManager.id2Node[from].AddEdge(pGManager.edgeCounter)

	pGManager.edgeCounter++

	if pGManager.isDirection == false {
		newEdge2 := NewGraphEdge(pGManager.edgeCounter, to, from)
		pGManager.id2Edge[pGManager.edgeCounter] = newEdge2
		pGManager.id2Node[to].AddEdge(pGManager.edgeCounter)

		pGManager.edgeCounter++
	}

	return nil
}

// 返回图是否是有向图
func (gm GraphManager) IsDirection() bool {
	return gm.isDirection
}

// 获取标识符为idx的节点
func (gm GraphManager) getNodeById(idx int) *GraphNode {
	if v, ok := gm.id2Node[idx]; ok {
		return v
	} else {
		return nil
	}
}

// 获取标识为idx的边
func (gm GraphManager) getEdgeById(idx int) *GraphEdge {
	if v, ok := gm.id2Edge[idx]; ok {
		return v
	} else {
		return nil
	}
}

// 获取一个节点的所有邻接节点的编号
func (gm GraphManager) getAllAdjacentNodes(sourceIdx int) ([]int, error) {
	v := gm.getNodeById(sourceIdx)
	if v == nil {
		return nil, errors.New("调用getAllAdjacentNodes时，指定的源节点编号不存在！")
	}

	// 获取所有领结边的编号
	vEdges := v.GetEdges()
	if len(vEdges) == 0 {
		return nil, nil
	}

	ret := make([]int, 0)
	for _, edgeIdx := range vEdges {
		edgeEntity := gm.getEdgeById(edgeIdx)
		if edgeEntity == nil {
			errStr := johnson_utility.ConcateString("调用getAllAdjacentNodes时，指定的源节点的", strconv.FormatInt(int64(sourceIdx), 10), "不存在边：", strconv.FormatInt(int64(edgeIdx), 10))
			return nil, errors.New(errStr)
		} else {
			ret = append(ret, edgeEntity.GetDestination())
		}
	}

	return ret, nil
}

// 广度遍历
func (gm *GraphManager) BFS(sourceIdx int) ([]*GraphNode, error) {
	gm.applyChan()
	defer gm.releaseChan()

	if _, ok := gm.id2Node[sourceIdx]; !ok {
		return nil, errors.New("广度遍历的源节点不存在")
	}

	ret := make([]*GraphNode, 0)
	q := queue.NewQueue(QUEUE_SIZE)

	currentVertex := gm.id2Node[sourceIdx]
	currentVertex.SetColor(CON_GRAY)
	q.Push(queue.NewQueueNode(currentVertex.GetId()))

	//log.Println("q: ", q)

	for q.Len() != 0 {
		qnode, err := q.Pop()
		if err != nil {
			return nil, err
		}

		// only for debug
		//log.Println("qnode", qnode)

		idx := qnode.Value
		allAdjacent, err1 := gm.getAllAdjacentNodes(idx)

		//log.Println("all adjacent: ", allAdjacent)
		if err1 != nil {
			return nil, err1
		}

		for _, vIdx := range allAdjacent {
			vertex := gm.getNodeById(vIdx)
			if vertex == nil {
				errStr := johnson_utility.ConcateString("编号为", strconv.FormatInt(int64(vIdx), 10), "的节点不存在")
				return nil, errors.New(errStr)
			}
			if vertex.GetColor() == CON_WHITE {
				vertex.SetColor(CON_GRAY)
				q.Push(queue.NewQueueNode(vIdx))
			} else {
				gm.isCirculation = true
			}
		}

		gm.id2Node[idx].SetColor(CON_BLACK)
		ret = append(ret, gm.id2Node[idx])

		//q.PrintQueue()
	}
	return ret, nil
}

// 深度遍历
func (gm *GraphManager) DFS() {
	gm.applyChan()
	defer gm.releaseChan()

	// 给所有的节点添加属性
	for _, v := range gm.id2Node {
		v.AddFeature("parent", nil)
		v.AddFeature("discovery", -1)
		v.AddFeature("finish", -1)
	}

	for _, v := range gm.id2Node {
		if v.GetColor() == CON_WHITE {
			gm.dfs_Visit(v)
		}
	}
}

// 深度遍历具体细节
func (gm *GraphManager) dfs_Visit(pgNode *GraphNode) error {
	gm.logicTime = gm.logicTime + 1
	pgNode.SetFeature("discovery", gm.logicTime)
	pgNode.SetColor(CON_GRAY)

	// 获取所有的邻接节点
	allAdjacent, err := gm.getAllAdjacentNodes(pgNode.GetId())
	if err != nil {
		str := johnson_utility.ConcateString("获取编号为", strconv.FormatInt(int64(pgNode.GetId()), 10), "的邻接节点时出错")
		return errors.New(str)
	}

	for _, v := range allAdjacent {
		node := gm.getNodeById(v)
		if node != nil {
			if node.GetColor() == CON_WHITE {
				node.SetFeature("parent", pgNode)
				gm.dfs_Visit(node)
			} else {
				// 当父节点存在还未完全访问结束时，如果通过直接点又访问到了父节点，则说明存在环
				if node.GetColor() == CON_GRAY {
					gm.isCirculation = true
				}
			}
		}
	}

	pgNode.SetColor(CON_BLACK)

	gm.logicTime = gm.logicTime + 1
	pgNode.SetFeature("finish", gm.logicTime)

	return nil
}

// 判断当前图是否存在循环，循环设置在深度遍历或广度遍历时设置
func (gm GraphManager) ExistCircle() bool {
	return gm.isCirculation
}

// 拓扑排序：根据深度遍历获取，返回节点
// 算法：当未深度遍历图时，深度遍历此图，然后对每个节点的访问结束时间到排序即可
func (gm *GraphManager) TopologySort() ([]*GraphNode, error) {
	if gm == nil || len(gm.id2Node) == 0 {
		return nil, errors.New("图中没有节点！")
	}

	if gm.isDirection == false {
		return nil, errors.New("无向图不可以进行拓扑排序")
	}

	pNode := gm.getNodeById(0)

	// 如果没有访问结束时间标识，表明未尽兴过深度遍历，对图进行深度遍历
	if _, err := pNode.GetFeature("finish"); err != nil {
		gm.DFS()
	}

	if gm.isCirculation == true {
		return nil, errors.New("图中存在环，不可能进行拓扑排序")
	}

	// 根据节点访问结束时间进行排序
	//retNodes := new(GraphNodes4SortByFinish)
	retNodes := make([]*GraphNode, 0)

	for _, value := range gm.id2Node {
		retNodes = append(retNodes, value)
	}

	sort.Sort(GraphNodes4SortByFinish(retNodes))

	return retNodes, nil
}
