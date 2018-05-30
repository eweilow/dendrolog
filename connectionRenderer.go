package dendrolog

type renderedNode struct {
	block stringBlock
	// Defines at which x coord in this block on which the zone we can connect lines to starts
	start int
	// Defines at which x coord in this block on which the zone we can connect lines to end
	end int
}

const baseSpacing = 3

//const debugConnections = false
const lineRow = 0

//const debugRow = 0
const connectRow = 1

func renderSingleChild(
	child renderedNode,
	self stringBlock,
) renderedNode {
	renderer := stringRenderer{}
	into := renderer.createBlockFromString("")

	start := 0
	end := 0

	x := (self.width - child.block.width) / 2
	if self.width > child.block.width {
		self.renderInto(0, 0, &into)
		child.block.renderInto(x, self.height+connectRow, &into)
		into = into.appendString(x+(child.end-child.start)/2, self.height+lineRow, "|")

		start = 0
		end = self.width

		/*if debugConnections {
			into = into.appendString(x, self.height+debugRow, repeat(child.end-child.start+1, debugChar))
		}*/
	} else {
		self.renderInto(-x, 0, &into)
		child.block.renderInto(0, self.height+connectRow, &into)
		into = into.appendString(-x, self.height+lineRow, "|")

		start = -x
		end = -x + self.width

		/*if debugConnections {
			into = into.appendString(0, self.height+debugRow, repeat(child.end-child.start+1, debugChar))
		}*/
	}

	return renderedNode{
		block: into,
		start: start,
		end:   end - 1,
	}
}

func renderBinaryChildren(
	left renderedNode,
	right renderedNode,
	self stringBlock,
) renderedNode {
	renderer := stringRenderer{}
	into := renderer.createBlockFromString("")

	spacing := intMax(intMax(baseSpacing, self.width-(left.block.width-left.end-1+right.start-1)), self.width+2)

	left.block.renderInto(0, self.height+connectRow, &into)
	right.block.renderInto(left.block.width+spacing, self.height+connectRow, &into)

	/*if debugConnections {
		into = into.appendString(left.start, self.height+debugRow, repeat(left.end-left.start+1, debugChar))
		into = into.appendString(left.block.width+spacing+right.start, self.height+debugRow, repeat(right.end-right.start+1, debugChar))
	}*/

	into = into.appendString(left.end+1, self.height+lineRow, "/")
	into = into.appendString(left.block.width+spacing+right.start-1, self.height+lineRow, "\\")

	leftMost := left.end + 1
	rightMost := left.block.width + spacing + right.start - 1

	selfStart := leftMost + intMax(0, (rightMost-leftMost-self.width)/2) + 1
	into = into.appendString(leftMost+1, 0, repeat(rightMost-leftMost-1, "_"))
	self.renderInto(selfStart, 0, &into)

	return renderedNode{
		block: into,
		start: selfStart,
		end:   selfStart + self.width - 1,
	}
}

func renderLines(
	children []renderedNode,
	self stringBlock,
) renderedNode {
	childCount := len(children)
	switch childCount {
	case 1:
		return renderSingleChild(children[0], self)
	case 2:
		return renderBinaryChildren(children[0], children[1], self)
	}

	renderer := stringRenderer{}
	into := renderer.createBlockFromString("")

	innerWidth := 0
	for _, child := range children[1 : len(children)-1] {
		innerWidth = innerWidth + child.block.width
	}

	freeSpace := intMax(0, self.width+2-innerWidth)
	childSpacing := intMax(baseSpacing, freeSpace/(len(children)-1))

	leftAngle := 0
	rightAngle := into.width

	offset := 0
	for i, child := range children {

		if i == 0 {
			leftAngle = offset + child.end + 2
			into = into.appendString(offset+child.end+1, self.height+lineRow, "/")
			child.block.renderInto(0, self.height+connectRow, &into)
			/*if debugConnections {
				into = into.appendString(child.start, self.height+debugRow, repeat(child.end-child.start+1, debugChar))
			}*/
		} else if i == len(children)-1 {
			rightAngle = offset + child.start
			into = into.appendString(rightAngle, self.height+lineRow, "\\")
			child.block.renderInto(rightAngle-child.start+1, self.height+connectRow, &into)
			/*if debugConnections {
				into = into.appendString(rightAngle+1, self.height+debugRow, repeat(child.end-child.start+1, debugChar))
			}*/
		} else {
			into = into.appendString(offset+child.start+(child.end-child.start)/2, self.height+lineRow, "|")
			child.block.renderInto(offset, self.height+connectRow, &into)
			/*if debugConnections {
				into = into.appendString(offset+child.start, self.height+debugRow, repeat(child.end-child.start+1, debugChar))
			}*/
		}

		offset = offset + childSpacing + child.block.width
	}

	x := (leftAngle + rightAngle - self.width) / 2

	self.renderInto(x, 0, &into)

	into = into.appendString(leftAngle, 0, repeat(x-leftAngle, "_"))
	into = into.appendString(x+self.width, 0, repeat(rightAngle-(x+self.width), "_"))

	return renderedNode{
		block: into,
		start: x,
		end:   x + self.width - 1,
	}
}
