// ...

// unsafe.go.

package main

import (
	"fmt"
	"github.com/legacy-vault/library/go/compact_double_link_list"
)

func exampleUnsafe() {

	var i uint64
	var itemX *cdllist.ListItem
	var itemY *cdllist.ListItem
	var itemZ *cdllist.ListItem
	var list *cdllist.List
	var ok bool
	var stressTestCounts uint64

	list = cdllist.CreateNewList()
	fmt.Println(list.EnlistAllItemValues())

	list.InsertHead(&cdllist.ListItem{Data: "B"})
	list.InsertHead(&cdllist.ListItem{Data: "C"})
	list.InsertHead(&cdllist.ListItem{Data: "D"})
	ok = list.IsIntegral()
	fmt.Println(list.EnlistAllItemValues(), ok)

	list.InsertTail(&cdllist.ListItem{Data: "A"})
	fmt.Println(list.EnlistAllItemValues())

	itemX = &cdllist.ListItem{Data: "X"}
	list.InsertHead(itemX)
	list.InsertHead(&cdllist.ListItem{Data: "E"})
	list.InsertHead(&cdllist.ListItem{Data: "F"})
	list.InsertHead(&cdllist.ListItem{Data: "G"})
	fmt.Println(list.EnlistAllItemValues())

	itemY = &cdllist.ListItem{Data: "Y"}
	list.InsertNextItemUnsafe(itemX, itemY)
	list.InsertPreviousItemUnsafe(&cdllist.ListItem{Data: "W"}, itemX)
	fmt.Println(list.EnlistAllItemValues())

	list.MoveItemToHeadPositionUnsafe(itemX)
	list.MoveItemToTailPositionUnsafe(itemX.GetNextItem().GetNextItem())
	fmt.Println(list.EnlistAllItemValues())

	list.MoveItemToBeforeReferenceUnsafe(itemY, itemX)
	// 'K' does not exist, so this Move does not happen.
	//list.MoveItemToAfterReferenceUnsafe(itemX, &cdllist.ListItem{Data: "K"})
	fmt.Println(list.EnlistAllItemValues())

	list.MoveItemToAfterReferenceUnsafe(itemX, list.GetTail().GetPreviousItem())
	fmt.Println(list.EnlistAllItemValues())

	list.RemoveTail()
	list.RemoveHead()
	fmt.Println(list.EnlistAllItemValues())

	itemZ = &cdllist.ListItem{Data: "Z"}
	list.InsertNextItemUnsafe(itemX, itemZ)
	fmt.Println(list.EnlistAllItemValues())

	list.RemoveItemUnsafe(itemZ)
	fmt.Println(list.EnlistAllItemValues())

	fmt.Println(
		list.IsIntegral(),
		list.GetHead().Data,
		list.GetTail().Data,
		list.GetSize(),
		list.IsEmpty(),
		list.HasItems(),
		list.HasAnItem(itemX),
		list.HasAnItem(itemZ),
	)

	stressTestCounts = 1000
	for i = 1; i <= stressTestCounts; i++ {
		list.RemoveHead()
		list.InsertTail(&cdllist.ListItem{Data: i})
	}
	for i = 1; i <= stressTestCounts; i++ {
		list.RemoveTail()
		list.InsertHead(&cdllist.ListItem{Data: i})
	}
	fmt.Println(list.EnlistAllItemValues())

	list.RemoveTail()
	list.RemoveHead()
	fmt.Println(list.EnlistAllItemValues())

	ok = cdllist.ClearList(list)
	fmt.Println(
		ok,
		list.GetHead(),
		list.GetTail(),
		list.GetSize(),
		list.IsEmpty(),
		list.HasItems(),
		list.HasAnItem(itemX),
		list.HasAnItem(itemZ),
	)
	fmt.Println(list.EnlistAllItems())
	fmt.Println(list.EnlistAllItemValues())
}
