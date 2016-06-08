package treap

type Treap struct {
	root *node
}

type node struct {
	id       int
	priority int
	cnt      int
	left     *node
	right    *node
}

func NewTreap() *Treap {
	return &Treap{root: nil}
}

func (t *Treap) Add(itemId int, itemPriority int) *Treap {
	r := t.union(t.root, &node{id: itemId, priority: itemPriority, cnt:1})
	return &Treap{root: r}
}

func (t *Treap) Traverse(n *node, add int) {
	if n != nil {
		// curKey := add + t.cnt(n.left)
		t.Traverse(n.left, add)
		t.Traverse(n.right, add + t.cnt(n.left) + 1)
	}
}

func (t *Treap) union(this *node, that *node) *node {
	if this == nil {
		return that
	}
	if that == nil {
		return this
	}

	if this.priority > that.priority {
		this.right = t.union(this.right, that)
		t.updateCnt(this)
		return this;
	} else {
		that.left = t.union(this, that.left)
		t.updateCnt(that)
		return that;
	}

}

func (t *Treap) split(n *node, key int, add int) (*node, *node) {
	if n == nil {
		return nil, nil
	}

	curKey := add + t.cnt(n.left)

	if (key <= curKey) {
		left, right := t.split(n.left, key, add)
		n.left = right
		t.updateCnt(n);
		return left, n
	} else {
		left, right := t.split(n.right, key, add + 1 + t.cnt(n.left))
		n.right = left
		t.updateCnt(n);
		return n, right
	}
}

func (t *Treap) cnt(n *node) int {
	if n != nil {
		return n.cnt
	}
	return 0;
}

func (t *Treap) updateCnt(n *node) {
	if n != nil {
		n.cnt = 1 + t.cnt(n.left) + t.cnt(n.right)
	}
}
