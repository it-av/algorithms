package structures

type Treap struct {
	root *treapNode
}

type treapNode struct {
	id       int
	priority int
	cnt      int
	left     *treapNode
	right    *treapNode
}

func NewTreap() *Treap {
	return &Treap{root: nil}
}

func (t *Treap) Add(itemId int, itemPriority int) {
	t.root = t.union(t.root, &treapNode{id: itemId, priority: itemPriority, cnt: 1})
}

func (t *Treap) Traverse(n *treapNode, add int) {
	if n != nil {
		// curKey := add + t.cnt(n.left)
		t.Traverse(n.left, add)
		t.Traverse(n.right, add+t.cnt(n.left)+1)
	}
}

func (t *Treap) union(this *treapNode, that *treapNode) *treapNode {
	if this == nil {
		return that
	}
	if that == nil {
		return this
	}

	if this.priority > that.priority {
		this.right = t.union(this.right, that)
		t.updateCnt(this)
		return this
	} else {
		that.left = t.union(this, that.left)
		t.updateCnt(that)
		return that
	}

}

func (t *Treap) split(n *treapNode, key int, add int) (*treapNode, *treapNode) {
	if n == nil {
		return nil, nil
	}

	curKey := add + t.cnt(n.left)

	if key <= curKey {
		left, right := t.split(n.left, key, add)
		n.left = right
		t.updateCnt(n)
		return left, n
	} else {
		left, right := t.split(n.right, key, add+1+t.cnt(n.left))
		n.right = left
		t.updateCnt(n)
		return n, right
	}
}

func (t *Treap) cnt(n *treapNode) int {
	if n != nil {
		return n.cnt
	}
	return 0
}

func (t *Treap) updateCnt(n *treapNode) {
	if n != nil {
		n.cnt = 1 + t.cnt(n.left) + t.cnt(n.right)
	}
}
