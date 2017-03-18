package queue

import (
	"errors"
	"log"
)

// 队列管理结构
// 用slice实现了一个环形队列，当head > tail时，队列元素空间为[tail, head]，其余空间空闲；当 head < tail, 队列元素空间为[0, head] + [tail, size - 1],其余空间空闲；初识时head=tail=0且isEmpty=true; 队列满有两种情形：1）当head=size-1且tail=0时，队列满；2）当head=tail-1时，队列满。 队列空的判断：head=tail=0，且isEmpty = true
type Queue struct {
	nodes   []*QueueNode
	head    int  // 指示队列头部元素
	tail    int  // 指示队列尾部元素
	isEmpty bool // 指示是否时空闲

	lockChan chan bool // 控制访问
}

func (q Queue) PrintQueue() {
	log.Println("<=====start print queue=====>")

	if q.isEmpty {
		log.Println("队列为空")
	} else {
		if q.tail <= q.head {
			for i := q.head; i >= q.tail; i-- {
				log.Println(*(q.nodes[i]))
			}
		} else {
			for i := q.head; i >= 0; i-- {
				log.Println(*(q.nodes[i]))
			}
			for i := len(q.nodes) - 1; i >= q.tail; i-- {
				log.Println(*(q.nodes[i]))
			}
		}
	}

	log.Println("<=====end print queue=====>")
}

// 创建队列
func NewQueue(iSize int) *Queue {
	return &Queue{make([]*QueueNode, iSize), 0, 0, true, make(chan bool, 1)}
}

// 加锁及解锁操作
func (pQ *Queue) lock() {
	pQ.lockChan <- true
}

func (pQ *Queue) unlock() {
	<-pQ.lockChan
}

// 向队列中添加对象
func (pQueue *Queue) Push(pNode *QueueNode) error {
	go pQueue.lock()
	defer pQueue.unlock()

	if pNode == nil {
		return errors.New("添加的元素不能为空")
	}

	if pQueue.head == pQueue.tail-1 || (pQueue.head == len(pQueue.nodes)-1 && pQueue.tail == 0) {
		return errors.New("队列已满！！！")
	}

	if pQueue.isEmpty {
		pQueue.nodes[pQueue.head] = pNode
	} else if pQueue.head < len(pQueue.nodes)-1 {
		pQueue.head = pQueue.head + 1
		pQueue.nodes[pQueue.head] = pNode
	} else {
		pQueue.head = 0
		pQueue.nodes[pQueue.head] = pNode
	}
	pQueue.isEmpty = false

	return nil
}

// 队列中元素个数
func (q Queue) Len() int {
	go q.lock()
	defer q.unlock()

	if q.isEmpty {
		return 0
	}

	if q.head >= q.tail {
		return q.head - q.tail + 1
	} else {
		return len(q.nodes) - (q.tail - q.head - 1)
	}
}

// 队列的容量
func (q Queue) Capacity() int {
	return len(q.nodes)
}

// 出队列
func (pQueue *Queue) Pop() (*QueueNode, error) {
	go pQueue.lock()
	defer pQueue.unlock()

	if pQueue.Len() == 0 {
		return nil, errors.New("队列中无元素")
	}

	ret := pQueue.nodes[pQueue.tail]

	// 如果出队列时，tail与head相等，说明已是队列的最后一个元素
	if pQueue.tail == pQueue.head {
		pQueue.tail, pQueue.head = 0, 0
		pQueue.isEmpty = true
	} else {
		pQueue.tail++
	}

	return ret, nil
}
