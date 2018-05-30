package dendrolog

import (
	"bytes"
	"strings"
)

// stringBlock defines a rectangular block of width by height chars
type stringBlock struct {
	backingData []string
	width       int
	height      int
}

// stringRenderer defines a type containing a stringBlock, which can be used for advanced rendering
type stringRenderer struct {
	main stringBlock
}

// CreateBlock creates a rectangular block containing chars
func (renderer *stringRenderer) createBlockFromString(str string) stringBlock {
	backingData := strings.Split(str, "\n")
	width := maxStringLength(backingData)

	return stringBlock{
		backingData: backingData,
		width:       width,
		height:      len(backingData),
	}
}

// String returns a stringified (with newlines) version of this
func (block *stringBlock) string() string {
	var data bytes.Buffer
	for cy := 0; cy < block.height; cy++ {
		if cy < block.height-1 {
			data.WriteString(block.backingData[cy] + "\n")
		} else {
			data.WriteString(block.backingData[cy])
		}
	}
	return data.String()
}

// Render a stringBlock into another stringBlock pointer
func (block *stringBlock) renderInto(x int, y int, otherBlock *stringBlock) {
	for i := 0; i < block.height; i++ {
		*otherBlock = otherBlock.appendString(x, y+i, block.backingData[i])
	}
}

// Append a string into this stringBlock
func (block *stringBlock) appendString(x int, y int, str string) stringBlock {
	backingData := block.backingData
	//extend existing array
	for cy := block.height; cy <= y; cy++ {
		backingData = append(backingData, "")
	}

	newWidth := intMax(block.width, x+len(str))
	newHeight := intMax(block.height, y+1)

	var originalRowLength int
	if y < block.height {
		originalRowLength = len(backingData[y])
	}

	var row bytes.Buffer
	for cx := 0; cx < newWidth; cx++ {
		if cx >= x && cx < x+len(str) {
			row.WriteByte(str[cx-x])
		} else if cx < originalRowLength {
			row.WriteByte(backingData[y][cx])
		} else {
			row.WriteString(" ")
		}
	}
	backingData[y] = row.String()

	return stringBlock{
		backingData: backingData,
		width:       newWidth,
		height:      newHeight,
	}
}
