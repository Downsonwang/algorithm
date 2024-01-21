package graph

import (
	"container/list"
	"fmt"
)

type Graph struct {
	adj []*list.List // 邻接表的数组
	v   int          //图中顶点数量
}

func newGraph(v int) *Graph {
	graph := &Graph{}
	graph.v = v
	graph.adj = make([]*list.List, v)
	for i := range graph.adj {
		graph.adj[i] = list.New()
	}
	return graph
}

// this.adj[s].PushBack(t) 是将顶点 t 添加到图中顶点 s 对应的邻接表中
// PushBack(t) 是链表的操作，将元素 t 添加到链表的末尾。
func (this *Graph) addEdge(s int, t int) {
	this.adj[s].PushBack(t)
	this.adj[t].PushBack(s)
}

// 将s放到队列中,
func (this *Graph) BFS(s int, t int) {
	// 广度节点
	if s == t {
		return
	}

	prev := make([]int, this.v)
	// 前驱索引未知
	for i := range prev {
		prev[i] = -1
	}

	var queue []int
	visited := make([]bool, this.v)
	// 加入源值
	queue = append(queue, s)
	visited[s] = true
	isFound := false

	for len(queue) > 0 && !isFound {
		// 出队
		// 取出队列的第一个元素
		top := queue[0]
		// 将已经处理的顶点出队
		queue = queue[1:]
		// 获取当前top对应的邻接表
		linkedlist := this.adj[top]
		// 遍历当前顶点的领结表 遍历与top相邻的顶点
		for e := linkedlist.Front(); e != nil; e = e.Next() {
			k := e.Value.(int)
			// 如果顶点 k 尚未被访问过。

			if !visited[k] {
				// 顶点k的前驱节点top
				prev[k] = top
				// 如果顶点 k 等于目标顶点 t，表示找到了从 s 到 t 的路径，将 isFound 设为 true 并退出循环。

				if k == t {
					isFound = true
					break
				}
				queue = append(queue, k)
				// 将顶点 k 标记为已访问，防止重复访问。

				visited[k] = true
			}
		}

	}
	if isFound {
		printPrev(prev, s, t)
	} else {
		fmt.Println("dont find any paths from s to t")
	}

}

func (this *Graph) DFS(s int, t int) {
	prev := make([]int, this.v)
	for i := range prev {
		prev[i] = -1
	}

	visited := make([]bool, this.v)
	visited[s] = true
	isFound := false
	this.recurse(prev, s, t, visited, isFound)
	printPrev(prev, s, t)
}

func (this *Graph) recurse(prev []int, s int, t int, visited []bool, isFound bool) {
	if s == t {
		return
	}
	visited[s] = true
	isFound = false
	// 获取当前顶点s对应的领接表
	linkedlist := this.adj[s]
	for e := linkedlist.Front(); e != nil; e = e.Next() {
		k := e.Value.(int)
		// 如果顶点 k 尚未被访问，则记录当前顶点 s 为顶点 k 的前驱顶点，并递归调用 recurse。

		if !visited[k] {
			prev[k] = s
			this.recurse(k, t, prev, visited, false)

		}
	}
}
func printPrev(prev []int, s int, t int) {
	if t == s || prev[t] == -1 {
		fmt.Println("%d", t)
	} else {
		printPrev(prev, s, t)
	}
}
