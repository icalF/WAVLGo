package wavlgo

import "fmt"

type keytype interface {
	LessThan(interface{}) bool
}

type valuetype interface{}

type node struct {
	left, right, parent *node
	Key                 keytype
	Value               valuetype
}

// Tree is a struct of red-black tree
type Tree struct {
	root *node
	size int
}

// NewTree returns a new rbtree
func NewTree() *Tree {
	return &Tree{}
}

// Find finds the node and return its value
func (t *Tree) Find(key keytype) interface{} {
	n := t.findnode(key)
	if n != nil {
		return n.Value
	}
	return nil
}

// FindIt finds the node and return it as a iterator
func (t *Tree) FindIt(key keytype) *node {
	return t.findnode(key)
}

// Empty checks whether the rbtree is empty
func (t *Tree) Empty() bool {
	if t.root == nil {
		return true
	}
	return false
}

// Iterator creates the rbtree's iterator that points to the minmum node
func (t *Tree) Iterator() *node {
	return minimum(t.root)
}

// Size returns the size of the rbtree
func (t *Tree) Size() int {
	return t.size
}

// Clear destroys the rbtree
func (t *Tree) Clear() {
	t.root = nil
	t.size = 0
}

// Insert inserts the key-value pair into the rbtree
func (t *Tree) Insert(key keytype, value valuetype) {}

// Delete deletes the node by key
func (t *Tree) Delete(key keytype) {}

// Preorder prints the tree in pre order
func (t *Tree) Preorder() {
	fmt.Println("preorder begin!")
	if t.root != nil {
		t.root.preorder()
	}
	fmt.Println("preorder end!")
}

// findnode finds the node by key and return it,if not exists return nil
func (t *Tree) findnode(key keytype) *node {
	return nil
}

// transplant transplants the subtree u and v
func (t *Tree) transplant(u, v *node) {}

// Next returns the node's successor as an iterator
func (n *node) Next() *node {
	return successor(n)
}

func (n *node) preorder() {
	fmt.Printf("(%v %v)", n.Key, n.Value)
	if n.parent == nil {
		fmt.Printf("nil")
	} else {
		fmt.Printf("whose parent is %v", n.parent.Key)
	}
	if n.left != nil {
		fmt.Printf("%v's left child is ", n.Key)
		n.left.preorder()
	}
	if n.right != nil {
		fmt.Printf("%v's right child is ", n.Key)
		n.right.preorder()
	}
}

// successor returns the successor of the node
func successor(x *node) *node {
	if x.right != nil {
		return minimum(x.right)
	}
	y := x.parent
	for y != nil && x == y.right {
		x = y
		y = x.parent
	}
	return y
}

// minimum finds the minimum node of subtree n.
func minimum(n *node) *node {
	for n.left != nil {
		n = n.left
	}
	return n
}

// maximum finds the maximum node of subtree n.
func maximum(n *node) *node {
	for n.right != nil {
		n = n.right
	}
	return n
}
