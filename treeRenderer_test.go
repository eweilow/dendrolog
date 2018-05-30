package dendrolog

import (
	"testing"
)

func TestTreeRenderer(t *testing.T) {
	t.Run("Should work for compact trees", func(t *testing.T) {
		tree := trieTree{
			value: "a",
		}
		left, middle, right := tree.setChildren("a", "b", "c")
		left.setChildren("e1", "f1", "g1")
		middle.setChildren("e2", "f2", "g2")
		right.setChildren("e3", "f3", "g3")

		renderer := TreeRenderer{}
		renderer.CollectFromTree(tree, trieCollector)
		stringified := renderer.Render(trieRenderer)

		testMatch(t, `
			&        ______________a_______________        ;
			&       /              |               \       ;
			&   ___a___         ___b___          ___c___   ;
			&  /  |    \       /  |    \        /  |    \  ;
			&e1   f1    g1   e2   f2    g2    e3   f3    g3;
		`, stringified)
	})

	t.Run("Should work for non-compact trees", func(t *testing.T) {
		tree := trieTree{
			value: "a",
		}
		left, middle, right := tree.setChildren("averylongnode", "bverylongnode", "cverylongnode")
		left.setChildren("e1", "f1", "g1verylongnode")
		middle.setChildren("e1verylongnode", "f2", "g2")
		right.setChildren("e3", "f3", "g3verylongnode")

		renderer := TreeRenderer{}
		renderer.CollectFromTree(tree, trieCollector)
		stringified := renderer.Render(trieRenderer)

		testMatch(t, `
&                 __________________________a___________________________                             ;
&                /                                      |               \                            ;
&   averylongnode                                 bverylongnode          cverylongnode               ;
&  /     |       \                               /     |       \        /     |       \              ;
&e1      f1       g1verylongnode   e1verylongnode      f2       g2    e3      f3       g3verylongnode;
		`, stringified)
	})
	t.Run("Should work for trees where nodes have 1 or 2 children", func(t *testing.T) {
		tree := trieTree{
			value: "a",
		}
		left, middle, _ := tree.setChildren("a", "b", "")
		left.setChildren("e1", "", "g1")
		middle.setChildren("", "f3", "g3")

		renderer := TreeRenderer{}
		renderer.CollectFromTree(tree, trieCollectorNoNil)
		stringified := renderer.Render(trieRenderer)

		testMatch(t, `
			&     ___a___     ;
			&    /       \    ;
			&   a         b   ;
			&  / \       / \  ;
			&e1   g1   f3   g3;
		`, stringified)
	})

	t.Run("Should work for trees where nodes have 1 or 2 children and items are wide", func(t *testing.T) {
		tree := trieTree{
			value: "11231231312",
		}
		left, middle, _ := tree.setChildren("cccc", "b", "")
		left.setChildren("", "222", "")
		middle.setChildren("", "", "adasdasd")

		renderer := TreeRenderer{}
		renderer.CollectFromTree(tree, trieCollectorNoNil)
		stringified := renderer.Render(trieRenderer)

		testMatch(t, `
			&     __11231231312_      ;
			&    /              \     ;
			&cccc                b    ;
			& |                  |    ;
			&222              adasdasd;
		`, stringified)
	})

	t.Run("Should work for trees without children", func(t *testing.T) {
		tree := trieTree{
			value: "tree",
		}

		renderer := TreeRenderer{}
		renderer.CollectFromTree(tree, trieCollectorNoNil)
		stringified := renderer.Render(trieRenderer)

		testMatch(t, `
			&tree;
		`, stringified)
	})

	t.Run("Should work for trees that have not been collected", func(t *testing.T) {
		renderer := TreeRenderer{}
		stringified := renderer.Render(trieRenderer)

		testMatch(t, `
			&nil;
		`, stringified)
	})
}
