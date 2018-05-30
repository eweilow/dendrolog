![logo](./logo.png)

# Dendrolog

Dendrolog is a Go package for debugging trees, by printing them into an ASCII-art-like representation.

Implemented as final project of the course [DD1327 - Fundamentals of computer science](https://www.kth.se/social/course/DD1327/) at KTH Royal Institute of Technology.

## Examples

Take a treap as an example:

```go
type Treap struct {
  left *Treap
  right *Treap
  value string
  priority int
}

func (treap *Treap) AddWithPriority(value string, priority int) {
  // Your basic Treap implementation
}

treap := Treap{}
treap.AddWithPriority("Z", 5)
treap.AddWithPriority("2", 25)
treap.AddWithPriority("B", 150)
treap.AddWithPriority("A", 2)
treap.AddWithPriority("C", 13)
treap.AddWithPriority("F", 14)
```

An example of the above treap, printed with Dendrolog, looks like this:

```
  ______A_2
 /         \
2_25      Z_5__
         /     \
       _C_13    nil
      /     \
     B_150 F_14
```

The code needed to print the above is:

```go
// treap is initialized as above

renderer := TreeRenderer{}
renderer.CollectFromTree(treap, func(node interface{}, child func(childPointer interface{})) {
  treapNode := node.(Treap) // we are free to cast the provided node to whatever we want
  if treapNode.left != nil {
    child(*treapNode.left) // child must receive non-pointer values
  } else {
    child(nil);
  }
  if treapNode.right != nil {
    child(*treapNode.right) // child must receive non-pointer values
  } else {
    child(nil);
  }
})

stringified := renderer.Render(func(node interface{}) string {
  treapNode := node.(Treap) // we are free to cast the provided node to whatever we want

  return fmt.Sprintf("%s_%d", treapNode.value, treapNode.priority)
})

fmt.Print(stringified)
```

## API

### TreeRenderer

TreeRenderer provides a type which print arbitrary trees to ASCII-like text.

```go
type TreeRenderer struct {}
```

#### func CollectFromTree

```go
func (renderer *TreeRenderer) CollectFromTree(
  tree interface{},
  selector func(
    node interface{},
    child func(childPointer interface{})
  ),
)
```

CollectFromTree walks an arbitrary tree and expects the selector callback to tell what the children of each walked node are.

For every invocation, `selector` must call `child` with each child of the `node` argument. `child` can be provided `nil`, to make certain trees more clear.

#### func Render

```go
func (renderer *TreeRenderer) Render(stringifier func(node interface{}) string) string
```

Render takes a result collected in `CollectFromTree` and renders a tree into text, running `stringifier` on each node of the tree.

`stringifier` is called for every node of the tree provided to `CollectFromTree` and must return a string representation of each node.
Simply returns the string "nil" if no nodes was collected, or `CollectFromTree` was never called.

## Future

* The API of this library is not frozen, but adhere to semantic versioning
* Only major version bumps will break existing functionality
