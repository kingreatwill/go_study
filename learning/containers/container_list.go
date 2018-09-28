// GO语言中的双向链表(container/list包);
package containers

import "fmt"
import "container/list"
import "github.com/kingreatwill/go_study/learning/common"

// list99;
func list99() {
	common.Println("container_list.go")

	// 创建一个链表;
	// func New() *List;
	l := list.New()
	// 清空链表;
	// func (l *List) Init() *List;
	l.Init()
	// 链表中元素的个数;
	// func (l *List) Len() int;
	fmt.Println("Len:", l.Len())

	// 返回链表第一个元素或nil;
	// func (l *List) Front() *Element;
	fmt.Println("Front:", l.Front())
	// 返回链表最后一个元素或nil;
	// func (l *List) Back() *Element;
	fmt.Println("Back:", l.Back())

	// 将一个值为v的新元素插入链表的第一个位置, 返回生成的新元素;
	// func (l *List) PushFront(v interface{}) *Element;
	fmt.Println("PushFront:", l.PushFront(1).Value)
	printList(l)
	// 创建链表other的拷贝, 并将拷贝的最后一个位置连接到链表l的第一个位置;
	// func (l *List) PushFrontList(other *List);
	// 将一个值为v的新元素插入链表的最后一个位置, 返回生成的新元素;
	// func (l *List) PushBack(v interface{}) *Element;
	fmt.Println("PushBack:", l.PushBack(99).Value)
	printList(l)
	// 创建链表other的拷贝, 并将链表l的最后一个位置连接到拷贝的第一个位置;
	// func (l *List) PushBackList(other *List);

	// 将一个值为v的新元素插入到mark前面, 并返回生成的新元素;
	// func (l *List) InsertBefore(v interface{}, mark *Element) *Element;
	fmt.Println("InsertBefore:", l.InsertBefore(98, l.Back()).Value)
	// 将一个值为v的新元素插入到mark后面, 并返回新生成的元素;
	// func (l *List) InsertAfter(v interface{}, mark *Element) *Element;
	fmt.Println("InsertAfter:", l.InsertAfter(2, l.Front()).Value)

	// 将元素e移动到链表的第一个位置;
	// func (l *List) MoveToFront(e *Element);
	l.MoveToFront(l.Back())
	printList(l)
	// 将元素e移动到链表的最后一个位置;
	// func (l *List) MoveToBack(e *Element);
	l.MoveToBack(l.Front())
	printList(l)
	// 将元素e移动到mark的前面;
	// func (l *List) MoveBefore(e, mark *Element);
	// 将元素e移动到mark的后面;
	// func (l *List) MoveAfter(e, mark *Element);

	// 删除链表中的元素e, 并返回e.Value;
	// func (l *List) Remove(e *Element) interface{};
	l.Remove(l.Back())
	printList(l)

	// 编译输出:
	// Len: 0
	// Front: <nil>
	// Back: <nil>
	// PushFront: 1
	// 1
	// PushBack: 99
	// 1 99
	// InsertBefore: 98
	// InsertAfter: 2
	// 99 1 2 98
	// 1 2 98 99
	// 1 2 98
}

// 遍历链表(元素);
func printList(l *list.List) {
	// 返回链表的后一个元素或者nil;
	// func (e *Element) Next() *Element;
	// 返回链表的前一个元素或者nil;
	// func (e *Element) Prev() *Element;
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}
