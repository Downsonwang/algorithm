/*
 * @Descripttion: Prims 在加权连通图中构建一个最小生成树
 * @Author:DW
 * @Date: 2024-04-08 13:49:50
 * @LastEditTime: 2024-04-08 15:23:49
 */
package main

import (
	"fmt"
	"math"
)

type Graph struct {
	Verxs  int
	Data   []rune
	Weight [][]int
}

func main() {

	inf := math.MaxInt32

	// 构建一张图
	vertices := []rune{'A', 'B', 'C', 'D', 'E', 'F'}

	v_len := len(vertices)

	weight := [][]int{
		{inf, 5, 6, 4, inf, inf},   //A
		{5, inf, 1, 2, inf, inf},   //B
		{6, 1, inf, 2, 5, 3},       //C
		{4, 2, 2, inf, inf, 4},     //D
		{inf, inf, 5, inf, inf, 4}, //E
		{inf, inf, 3, 4, 4, inf},   // F
	}

	graph := NewGraph(v_len)
	// 创建图
	CreateGraph(graph, v_len, vertices, weight)
	fmt.Println("图的邻接矩阵:")
	ShowGraph(graph)

	fmt.Println("最小生成树的顶点以及权值:")
	Prim(graph, 0)
}

func NewGraph(v_len int) *Graph {
	return &Graph{
		Verxs:  v_len,
		Data:   make([]rune, v_len),
		Weight: make([][]int, v_len),
	}
}

func CreateGraph(graph *Graph, v_len int, vetices []rune, weight [][]int) {
	for i := 0; i < v_len; i++ {
		graph.Data[i] = vetices[i]
		graph.Weight[i] = make([]int, v_len)
		copy(graph.Weight[i], weight[i])
	}
	fmt.Println(graph)
}

func ShowGraph(graph *Graph) {
	for _, link := range graph.Weight {
		fmt.Println(link)
	}
}

func Prim(graph *Graph, v int) {
	// 用于标记顶点是否被访问
	visited := make([]int, graph.Verxs)
	visited[v] = 1

	//
	var h1, h2 int

	minWeight := math.MaxInt32
	// 控制边数 n个顶点 n-1条边
	for k := 0; k < graph.Verxs-1; k++ {
		// 开始迭代 访问过的顶点
		for i := 0; i < graph.Verxs; i++ {
			// 访问未访问过的顶点
			for j := 0; j < graph.Verxs; j++ {
				if visited[i] == 1 && visited[j] == 0 && graph.Weight[i][j] < minWeight {
					minWeight = graph.Weight[i][j]
					// 记录目前权值最小的两个顶点
					h1 = i
					h2 = j
				}
			}
		}
		fmt.Printf("<%c, %c> 权值：%d\n", graph.Data[h1], graph.Data[h2], minWeight)

		visited[h2] = 1
		minWeight = math.MaxInt32

	}
}
