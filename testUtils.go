package dendrolog

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

type trieTree struct {
	value  string
	left   *trieTree
	middle *trieTree
	right  *trieTree
}

func (trie *trieTree) setChildren(left string, middle string, right string) (leftTree *trieTree, middleTree *trieTree, rightTree *trieTree) {
	if left != "" {
		leftTree = &trieTree{
			value: left,
		}
		trie.left = leftTree
	}
	if middle != "" {
		middleTree = &trieTree{
			value: middle,
		}
		trie.middle = middleTree
	}
	if right != "" {
		rightTree = &trieTree{
			value: right,
		}
		trie.right = rightTree
	}
	return
}

func trieCollector(node interface{}, child func(childPointer interface{})) {
	trie := (node).(trieTree)
	if trie.left != nil {
		child(*trie.left)
	} else {
		child(nil)
	}
	if trie.middle != nil {
		child(*trie.middle)
	} else {
		child(nil)
	}
	if trie.right != nil {
		child(*trie.right)
	} else {
		child(nil)
	}
}
func trieCollectorNoNil(node interface{}, child func(childPointer interface{})) {
	trie := (node).(trieTree)
	if trie.left != nil {
		child(*trie.left)
	}
	if trie.middle != nil {
		child(*trie.middle)
	}
	if trie.right != nil {
		child(*trie.right)
	}
}

func trieRenderer(node interface{}) string {
	if node == nil {
		return "nil"
	}
	trie := (node).(trieTree)

	return trie.value
}

func testMatch(t *testing.T, expected string, stringified string) {
	re := regexp.MustCompile(`&([^;&]*);`)

	matches := re.FindAllStringSubmatch(expected, -1)

	expectedRows := make([]string, len(matches))
	for i, row := range matches {
		expectedRows[i] = strings.TrimRight(row[1], " ")
		//fmt.Printf("'%s'\n", expectedRows[i])
	}

	for i, row := range strings.Split(strings.TrimRight(stringified, "\n "), "\n") {
		trimmedRow := strings.TrimRight(row, " ")

		if trimmedRow != expectedRows[i] {
			fmt.Print(stringified)
			t.Fatalf("Expected row %d to be equal.\n Expected row: '%s'\n      Got row: '%s'\n Expected Tree:\n%s\n      Got Tree:\n%s", i, expectedRows[i], trimmedRow, strings.Join(expectedRows, "\n"), stringified)
		}
	}
}
