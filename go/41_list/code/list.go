// list.go.

// List Functions.

package main

import (
	"container/list"
	"fmt"
	"strings"
)

// Print Symbols.
const SymbolBorderLeft = "["
const SymbolBorderRight = "]"
const SymbolItemHead = SymbolBorderLeft + "HEAD" + SymbolBorderRight
const SymbolItemTail = SymbolBorderLeft + "TAIL" + SymbolBorderRight
const SymbolChainLink = "="

const FormatA = "%v"

// Represents the List as a text Snake.
func listToSnake(aList list.List) string {

	var item *list.Element
	var itemStr string
	var output strings.Builder
	var tmpStr string

	output.WriteString(SymbolItemHead)

	for item = aList.Front(); item != nil; item = item.Next() {

		itemStr = fmt.Sprintf(FormatA, item.Value)
		tmpStr = SymbolChainLink +
			SymbolBorderLeft +
			itemStr +
			SymbolBorderRight +
			SymbolChainLink
		output.WriteString(tmpStr)
	}

	output.WriteString(SymbolItemTail)

	return output.String()
}
