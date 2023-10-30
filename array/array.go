/*
 * @Descripttion:
 * @Author:
 * @Date: 2023-10-12 15:23:19
 * @LastEditTime: 2023-10-12 17:09:49
 */
package array

import "fmt"

/**
  数组的插入 删除
*/

type Array struct {
	data []int
	len  int
}

// 初始化内存
func NewArray(cap int) *Array {
	if cap == 0 {
		return nil
	}
	return &Array{
		make([]int, cap),
		0,
	}
}

// 获取长度
func (this *Array) Len() int {
	return this.len
}

// 判断是否越界
func (this *Array) IsIndexCrosstheBound(index int) bool {
	if index >= cap(this.data) {
		return true //越界返回0
	} else {
		return false //没越界返回1
	}
}

// 通过索引查找数组中的值
func (this *Array) FindDataByIndex(index int) int {
	if this.IsIndexCrosstheBound(index) {
		return -1
	}
	a := this.data[index]
	println(a)
	return a
}

// 插入数据到索引index上
func (this *Array) InsertNumIntoDataInIndex(num int, index int) (res []int, err string) {
	if this.IsIndexCrosstheBound(index) {
		return nil, "index exception"
	}
	//先移动后面的值
	for i := this.len; i > index; i-- {
		this.data[i] = this.data[i+1]
	}

	this.data[index] = num
	this.len++
	return this.data, ""
}

// 插入到末尾
func (this *Array) InsertToTail() {
	_, _ = this.InsertNumIntoDataInIndex(1, this.len)
}

//删除在某个索引上的值
func (this *Array) DeleteNumAtIndexInData(index int) (res []int, err string) {
	if this.IsIndexCrosstheBound(index) {
		return nil, "index exception"
	}

	for i := index; i < this.len-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.len--
	return this.data, ""
}

// 打印数列
func (this *Array) Print() {
	var format string
	for i := 0; i < this.len; i++ {
		format += fmt.Sprintf("|%+v", this.data[i])
	}
	fmt.Println(format)
}
