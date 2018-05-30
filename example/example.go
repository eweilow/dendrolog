package dendrolog

import (
	"fmt"
	"testing"

	"github.com/eweilow/dendrolog"

	"./treap"
)

func treapCollector(currentNode interface{}, child func(childPointer interface{})) {
	treapNode := (currentNode).(treap.Node)

	left := treapNode.GetLeft()
	if left != nil {
		child(*left)
	} else {
		child(nil)
	}

	right := treapNode.GetRight()
	if right != nil {
		child(*right)
	} else {
		child(nil)
	}
}

func treapRenderer(currentNode interface{}) string {
	if currentNode == nil {
		return "nil"
	}
	treapNode := (currentNode).(treap.Node)

	return fmt.Sprintf("%s_%d", treapNode.GetValue(), treapNode.GetPriority())
}

func stringify(treap *treap.Treap) string {
	root := treap.GetRoot()

	renderer := TreeRenderer{}
	renderer.CollectFromTree(*root, treapCollector)
	stringified := renderer.Render(treapRenderer)
	return stringified
}

func TestExample(t *testing.T) {
	fmt.Print("\n\nExample follows below:\n\n")
	tr := &treap.Treap{}

	tr.AddWithPriority("A", 2)
	fmt.Println(stringify(tr))

	tr.AddWithPriority("B", 3)
	fmt.Println(stringify(tr))

	tr.AddWithPriority("C", 1)
	fmt.Println(stringify(tr))

	tr.AddWithPriority("D", 4)
	fmt.Println(stringify(tr))

	tr.AddWithPriority("E", 6)
	fmt.Println(stringify(tr))

	tr.AddWithPriority("F", 0)
	fmt.Println(stringify(tr))

	tr.Delete("C")
	fmt.Println(stringify(tr))
}
