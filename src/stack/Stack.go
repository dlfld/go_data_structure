package stack

import (
	"sync"
)

type Stack struct {
	size  int
	array []interface{}
	sync.RWMutex
}

// NewStack /**  获取一个栈实例
func NewStack() *Stack {
	stack := &Stack{size: -1}
	return stack
}

// Top /** 返回栈顶元素
func (stack *Stack) Top() interface{} {
	return stack.array[stack.size]
}

// IsEmpty /** 判断栈内元素是否为空
func (stack *Stack) IsEmpty() bool {
	if stack == nil || stack.size == -1 {
		return true
	}
	return false
}

// IsNotEmpty /** 判断栈内元素是否为空
func (stack *Stack) IsNotEmpty() bool {
	return stack.IsEmpty()
}

// Size /** 返回栈中元素个数，size从0开始计数，因此需要+1
func (stack *Stack) Size() int {
	return stack.size + 1
}

// Push /** 向栈内添加一个元素
func (stack *Stack) Push(value interface{}) bool {
	stack.Lock()
	stack.array = append(stack.array, value)
	stack.size += 1
	defer stack.Unlock()
	return true
}

// Pop /** 返回栈顶元素
func (stack *Stack) Pop() interface{} {
	//如果栈为空则返回空
	if stack.IsEmpty() {
		return nil
	}
	stack.Lock()
	res := stack.array[stack.size]
	stack.size -= 1
	defer stack.Unlock()
	return res
}
