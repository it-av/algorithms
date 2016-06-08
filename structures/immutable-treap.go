package structures


type ImmutableTreap struct {
	root *node
	callback func(key, value int)
}


type node struct {
	id       int
	priority int
	cnt      int
	left     *node
	right    *node
}

func NewImmutableTreap() *ImmutableTreap {
	return &ImmutableTreap{root: nil}
}

func (t *ImmutableTreap) Upsert(itemId int, itemPriority int) *ImmutableTreap {
	r := t.union(t.root, &node{id: itemId, priority: itemPriority, cnt: 1})
	return &ImmutableTreap{root: r}
}

func (t *ImmutableTreap) Traverse(fn func(id, value int)) {
	t.callback = fn
	t.traverse(t.root, 0)

}

func (t *ImmutableTreap) union(this *node, that *node) *node {
	if this == nil {
		return that
	}
	if that == nil {
		return this
	}

	var res *node

	if this.priority > that.priority {
		res = &node{
			id:       this.id,
			priority: this.priority,
			left:     this.left,
			right:    t.union(this.right, that),
		}
	} else {
		res = &node{
			id:       that.id,
			priority: that.priority,
			left:     t.union(this, that.left),
			right:    that.right,
		}
	}

	t.updateCnt(res)
	return res

}
func (t *ImmutableTreap) split(n *node, key int, add int) (*node, *node) {
	if n == nil {
		return nil, nil
	}
	curKey := add + t.cnt(n.left)
	var res *node

	if key <= curKey {
		left, right := t.split(n.left, key, add)
		res = &node{
			id:       n.id,
			priority: n.priority,
			left:     right,
			right:    n.right,
		}
		t.updateCnt(res)
		return left, res
	} else {
		left, right := t.split(n.right, key, add+1+t.cnt(n.left))
		res = &node{
			id:       n.id,
			priority: n.priority,
			left:     n.left,
			right:    left,
		}
		t.updateCnt(res)
		return res, right
	}
}

func (t *ImmutableTreap) cnt(n *node) int {
	if n != nil {
		return n.cnt
	}
	return 0
}

func (t *ImmutableTreap) updateCnt(n *node) {
	if n != nil {
		n.cnt = 1 + t.cnt(n.left) + t.cnt(n.right)
	}
}

func (t *ImmutableTreap) traverse(n *node, add int) {
	if n != nil {
		t.callback(add+t.cnt(n.left), n.id)
		t.traverse(n.left, add)
		t.traverse(n.right, add+t.cnt(n.left)+1)

	}
}


