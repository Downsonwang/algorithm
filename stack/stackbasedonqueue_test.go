/*
 * @Descripttion:测试
 * @Author:DW
 * @Date: 2023-11-12 10:38:09
 * @LastEditTime: 2023-11-12 10:44:11
 */
package stack

import (
	"fmt"
	"testing"
)

func TestMyStack(t *testing.T) {
	stack := Constructor()
	if !stack.Empty() {
		t.Error("Expected stack to be empty")
	}

	stack.Push(1)
	if top := stack.Top(); top != 1 {
		//t.Error("Expected top element to be 1, got %v", top)
	}
	stack.Push(2)
	if popped := stack.Pop(); popped != 2 {
		t.Errorf("expected %v", popped)
	}

	if !stack.Empty() {
		t.Error("Expected stack to be empty after pop")
	}

	popped := stack.Pop()
	fmt.Println(popped)

	top := stack.Top()
	if top != -1 {
		fmt.Println(top)
	}
}
