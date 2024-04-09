/*
 * @Descripttion:
 * @Author:
 * @Date: 2024-01-21 10:01:53
 * @LastEditTime: 2024-01-27 11:32:07
 */
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

// 将s放到队列中,  广度优先遍历 按层次来讲
func (this *Graph) BFS(s int, t int) {

	if s == t {
		return
	}
	// 前驱节点
	prev := make([]int, this.v)
	for index := range prev {
		prev[index] = -1
	}

	var queue []int
	visited := make([]bool, this.v)
	// 顶点
	queue = append(queue, s)
	visited[s] = true
	isFound := false

	for len(queue) > 0 && !isFound {
		top := queue[0]
		queue = queue[1:]
		linkedlist := this.adj[top]
		for e := linkedlist.Front(); e != nil; e = e.Next() {
			k := e.Value
			if !visited[k] {
				prev[k] = top
				if k == t {
					isFound = true
					break
				}
				queue = append(queue, k)
				visited[k] = true
			}
		}
	}
	if isFound {
		printPrev(prev, s, t)
	} else {
		fmt.Printf("no path found from %d to %d\n", s, t)

	}

}

func (this *Graph) DFS(s int, t int) {
	prev := make([]int, this.v)
	for index := range prev {
		prev[index] = -1
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
		k := e.Value
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
