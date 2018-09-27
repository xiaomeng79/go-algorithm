package binaryTree

//二叉树
//二叉树是每个节点最多有两个子树的树结构
type ElementType int//节点数据
//结点
type Node struct {
	Value ElementType
	Parent *Node
	Left *Node
	Right *Node
}

//创建节点
func NewNode(v ElementType) *Node {
	return &Node{Value:v}
}

/*
*节点比较
* n>m:1 n<m:-1 n=m:0
 */
func (n *Node) Compare(m *Node) int {
	if n.Value < m.Value {
		return -1
	} else if n.Value > m.Value {
		return 1
	} else {
		return 0
	}
}

//树
type Tree struct {
	Head *Node
	Size int
}

//创建树
func NewTree(n *Node) *Tree {
	if n == nil {
		return &Tree{}
	}
	return &Tree{Head:n,Size:1}
}

//插入,相同的节点值，放到右子树
func (t *Tree) Insert(i ElementType) {
	n := NewNode(i)//创建节点
	if t.Head == nil {//判断树的根节点
		t.Head = n
		t.Size++
		return
	}

	h := t.Head

	for {
		if n.Compare(h) == -1 {//小于parent，到左子节点
			if h.Left == nil {//无左子节点
				h.Left = n
				n.Parent = h
				break
			} else {
				h = h.Left
			}
		} else {//大于parent
			if h.Right == nil {
				h.Right = n
				n.Parent = h
				break
			} else {
				h = h.Right
			}
		}
	}
	t.Size++
}

//搜索
func (t *Tree) Search(i ElementType) *Node {
	h := t.Head
	n := NewNode(i)
	for h != nil {
		switch h.Compare(n) {
		case -1:
			h = h.Right
		case 1:
			h = h.Left
		case 0:
			return h
		default:
			panic("Node not found")
		}
	}
	panic("Node not found")
}

//删除
func (t *Tree) Delete(i ElementType) bool {
	var parent *Node

	h := t.Head
	n := NewNode(i)

	for h != nil {
		switch n.Compare(h) {
		case -1:
			parent = h
			h = h.Left
		case 1:
			parent = h
			h = h.Right
		case 0:
			if h.Left != nil {
				right := h.Right
				h.Value = h.Left.Value
				h.Left = h.Left.Left
				h.Right = h.Left.Right

				if right != nil {
					subTree := &Tree{Head: h}
					IterOnTree(right, func(n *Node) {
						subTree.Insert(n.Value)
					})
				}
				t.Size--
				return true
			}

			if h.Right != nil {
				h.Value = h.Right.Value
				h.Left = h.Right.Left
				h.Right = h.Right.Right

				t.Size--
				return true
			}

			if parent == nil {
				t.Head = nil
				t.Size--
				return true
			}

			if parent.Left == n {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
			t.Size--
			return true
		}
	}
	return false
}

func IterOnTree(n *Node, f func(*Node)) bool {
	if n == nil {
		return true
	}
	if !IterOnTree(n.Left, f) {
		return false
	}

	f(n)

	return IterOnTree(n.Right, f)
}