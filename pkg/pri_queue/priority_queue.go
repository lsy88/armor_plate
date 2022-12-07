package priqueue

import "container/heap"

//Queue 是我们在优先队列中管理的东西
type Queue struct {
	DepotID  uint64  //仓库id
	Distance float64 //距离
	Priority int     //仓库优先级
	Amount   int     //仓库持有货物量
	Index    int     //在堆中的索引
}

//优先级队列
type PriorityQueue []*Queue

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// 我们希望 Pop 给我们最高而不是最低的优先级，所以我们使用比这里更大的优先级。
	//赋予距离80%权重，仓库优先级15%权重，仓库持有量5%权重设置优先级队列的比较器
	return (pq[i].Distance*80 + float64(pq[i].Priority*15+int(pq[i].Amount)*5)) >
		(pq[j].Distance*80 + float64(pq[j].Priority*15+int(pq[j].Amount)*5))
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
	return
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Queue)
	item.Index = n
	*pq = append(*pq, item)
	return
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // 避免内存泄漏
	item.Index = -1 // 为了安全
	*pq = old[0 : n-1]
	return item
}

// GenQueue 生成优先级队列
func GenQueue(pg *PriorityQueue) {
	heap.Init(pg)
}
