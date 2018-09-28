package stack

import "errors"

//使用链表实现栈
//以下是使用单向链表，使用头插法实现，也可以使用标准库的双向链表实现
//节点，需要动态申请内存空间
type Node struct {
	data int
	next *Node
}

//栈
type ListStack struct {
	top *Node
}

func NewListStack() *ListStack {
	return &ListStack{}
}

//是否为空
func (ls *ListStack) IsEmpty() bool {
	return ls.top == nil
}

//入栈,头插法
func (ls *ListStack) Push(i int) {
	//新建节点
	n := &Node{i, nil}
	if ls.top != nil {
		n.next = ls.top
	}
	ls.top = n
}

//出栈,从头弹出
func (ls *ListStack) Pop() (int, error) {
	if ls.top == nil {
		return 0, errors.New("栈空")
	}
	i := ls.top.data
	ls.top = ls.top.next
	return i, nil

}

//查看
func (ls *ListStack) Peek() (int, error) {
	if ls.top == nil {
		return 0, errors.New("栈空")
	}
	return ls.top.data, nil
}
