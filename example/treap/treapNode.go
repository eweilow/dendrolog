package treap

type dataType = string

var dataDefault dataType

type Node struct {
	priority int
	value    dataType
	left     *Node
	right    *Node
}

// GetLeft returns the left child of this Node
func (currentNode *Node) GetLeft() *Node {
	return currentNode.left
}

// GetRight returns the right child of this Node
func (currentNode *Node) GetRight() *Node {
	return currentNode.right
}

// GetValue returns the value of this Node
func (currentNode *Node) GetValue() dataType {
	return currentNode.value
}

// GetPriority returns the priority of this Node
func (currentNode *Node) GetPriority() int {
	return currentNode.priority
}

func (currentNode *Node) healthy() bool {
	if currentNode.right != nil {
		if !currentNode.right.healthy() {
			return false
		}
		if currentNode.right.value < currentNode.value {
			return false
		}
	}
	if currentNode.left != nil {
		if !currentNode.left.healthy() {
			return false
		}
		if currentNode.left.value > currentNode.value {
			return false
		}
	}
	return true
}

func (currentNode *Node) rotateRight(root *Treap) {
	current := *currentNode
	copy := currentNode.left

	if currentNode == root.root {
		root.root = copy
	}

	current.left = copy.right
	copy.right = &current
}

func (currentNode *Node) rotateLeft(root *Treap) {
	current := *currentNode
	copy := currentNode.right

	if currentNode == root.root {
		root.root = copy
	}

	current.right = copy.left
	copy.left = &current
}

func (currentNode *Node) add(treap *Treap, priority int, data dataType) *Node {
	if data > currentNode.value {
		if currentNode.right != nil {
			newNode := currentNode.right.add(treap, priority, data)
			if newNode != nil {
				currentNode.right = newNode
			}
		} else {
			currentNode.right = &Node{
				priority: priority,
				value:    data,
			}
		}
		if currentNode.right.priority < currentNode.priority {
			currentNode.rotateLeft(treap)
			return currentNode.right
		}
	} else {
		if currentNode.left != nil {
			newNode := currentNode.left.add(treap, priority, data)
			if newNode != nil {
				currentNode.left = newNode
			}
		} else {
			currentNode.left = &Node{
				priority: priority,
				value:    data,
			}
		}
		if currentNode.left.priority < currentNode.priority {
			currentNode.rotateRight(treap)
			return currentNode.left
		}
	}
	return nil
}

func (currentNode *Node) stringify() string {
	output := ""
	if currentNode.left != nil {
		output += currentNode.left.stringify() + ", "
	}
	output += currentNode.value
	if currentNode.right != nil {
		output += ", " + currentNode.right.stringify()
	}

	return output
}

// returns true if this node is to be deleted,
// and an optional node that indicates if we are to replace anything
func (currentNode *Node) delete(treap *Treap, data dataType) (bool, *Node) {
	if currentNode.left == nil && currentNode.right == nil {
		if currentNode.value != data {
			// We ended up in a leaf that is not the one we wanted to delete
			panic("Something has gone wrong - we should never be here")
		}
	}

	if data == currentNode.value {
		// delete currentNode in some way

		if currentNode.right != nil && currentNode.left == nil {
			// the easy conditions when we replace the current node with a subtree
			return false, currentNode.right
		} else if currentNode.right == nil && currentNode.left != nil {
			// the easy conditions when we replace the current node with a subtree
			return false, currentNode.left
		} else if currentNode.right != nil && currentNode.left != nil {
			// the trickier cases when we have to rotate the tree in order to delete properly
			if currentNode.left.priority > currentNode.right.priority {
				currentNode.rotateLeft(treap)
				return currentNode.right.delete(treap, data)
			}
			currentNode.rotateRight(treap)
			return currentNode.left.delete(treap, data)
		}
		return currentNode.value == data, nil
	}

	if data > currentNode.value {
		// attempt to delete on right subtree

		if currentNode.right == nil {
			// We would have ended up in a right subtree that is nil
			panic("Something has gone wrong - we should never be here")
		}

		deleteRight, replaceRight := currentNode.right.delete(treap, data)
		if deleteRight {
			currentNode.right = nil
		}
		if replaceRight != nil {
			currentNode.right = replaceRight
		}
	} else {
		// attempt to delete on left subtree

		if currentNode.left == nil {
			// We would have ended up in a left subtree that is nil
			panic("Something has gone wrong - we should never be here")
		}

		deleteLeft, replaceLeft := currentNode.left.delete(treap, data)
		if deleteLeft {
			currentNode.left = nil
		}
		if replaceLeft != nil {
			currentNode.left = replaceLeft
		}
	}
	return false, nil
}
