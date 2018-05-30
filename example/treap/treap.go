package treap

import (
	"math/rand"
)

// Treap implements a randomized binary tree
type Treap struct {
	root      *Node
	nodeCount int
}

// Size returns the size of this Treap
func (treap *Treap) Size() int {
	return treap.nodeCount
}

// GetRoot returns the root of this Treap
func (treap *Treap) GetRoot() *Node {
	return treap.root
}

// Add adds an element to this Treap. Returns true if an element was added and false if an element was already in the list (and thus not added)
func (treap *Treap) Add(data dataType) bool {
	const priorityMax int32 = 10000

	return treap.AddWithPriority(data, int(rand.Int31n(priorityMax)))
}

// AddWithPriority adds an element to this Treap with fixed priority. Returns true if an element was added and false if an element was already in the list (and thus not added)
func (treap *Treap) AddWithPriority(data dataType, priority int) bool {
	const priorityMax int32 = 10000

	if treap.root == nil {
		treap.root = &Node{
			priority: priority,
			value:    data,
		}
		treap.nodeCount = 1
		return true
	}
	// if this treap already contains this data, do nothing
	if treap.Find(data) {
		return false
	}

	treap.root.add(treap, priority, data)
	treap.nodeCount++
	return true
}

// Delete deletes an element in this Treap and returns true if it was deleted
func (treap *Treap) Delete(data dataType) bool {
	if !treap.Find(data) {
		return false
	}

	treap.nodeCount--
	deleteRoot, replaceRoot := treap.root.delete(treap, data)
	if deleteRoot {
		treap.root = nil
	}
	if replaceRoot != nil {
		treap.root = replaceRoot
	}

	return true
}

// Find returns true if an element exist in this Treap, otherwise false
func (treap *Treap) Find(data dataType) bool {
	if treap.root == nil {
		return false
	}

	current := treap.root
	for current != nil {
		if current.value == data {
			return true
		}
		if current.value < data {
			current = current.right
		} else {
			current = current.left
		}
	}
	return false
}

// Stringify returns a string of all the elements of this Treap in order
func (treap *Treap) Stringify() string {
	if treap.root == nil {
		return ""
	}
	return treap.root.stringify()
}

func (treap *Treap) healthy() bool {
	if treap.root == nil {
		return true
	}
	return treap.root.healthy()
}
