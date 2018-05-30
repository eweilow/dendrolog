package dendrolog

// TreeRenderer provides a type which print arbitrary trees to ASCII-like text.
type TreeRenderer struct {
	main stringBlock

	collectionResult *collected
}

type inputNode interface {
	Children() []inputNode
}

type collected struct {
	current  interface{}
	children []*collected
}

// CollectFromTree takes an arbitrary tree data type and runs `selector` with tree as its' node argument.
// The children of each node provided to `selector` be provided using the `child` function with non-pointer values.
// All nodes of the tree are visited and children are visited in the order given by the order of calls to `child`.
func (renderer *TreeRenderer) CollectFromTree(tree interface{}, selector func(node interface{}, child func(childPointer interface{}))) {
	root := renderer.collectFromTree(tree, selector)
	renderer.collectionResult = root
}

func (renderer *TreeRenderer) collectFromTree(tree interface{}, selector func(node interface{}, child func(childPointer interface{}))) *collected {
	var collectedChildren []*collected
	selector(tree, func(child interface{}) {
		var val *collected
		if child != nil {
			val = renderer.collectFromTree(child, selector)
		} else {
			val = &collected{
				current: nil,
			}
		}
		collectedChildren = append(collectedChildren, val)
	})
	return &collected{
		children: collectedChildren,
		current:  tree,
	}
}

// Render takes a result collected in `CollectFromTree` and renders a tree into text, running `stringifier` on each node of the tree in originally visited order.
// `stringifier` is called for every node of the tree provided to `CollectFromTree` and must return a string representation of each node.
// Simply returns the string "nil" if no nodes was collected, or `CollectFromTree` was never called.
func (renderer *TreeRenderer) Render(stringifier func(node interface{}) string) string {
	if renderer.collectionResult == nil {
		return "nil"
	}

	rendered := renderer.render(renderer.collectionResult, stringifier)
	return rendered.block.string() + "\n"
}

func (renderer *TreeRenderer) render(current *collected, stringifier func(node interface{}) string) renderedNode {
	strRenderer := stringRenderer{}

	val := stringifier(current.current)
	root := strRenderer.createBlockFromString(val)

	allNil := true
	for _, child := range current.children {
		if child.current != nil {
			allNil = false
			break
		}
	}
	if allNil {
		return renderedNode{block: root, start: 0, end: len(val) - 1}
	}

	renderedChildren := []renderedNode{}

	for _, child := range current.children {
		renderedChild := renderer.render(child, stringifier)

		renderedChildren = append(renderedChildren, renderedChild)
	}
	return renderLines(renderedChildren, root)

}
