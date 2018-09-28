// GO语言中的环形链表(container/ring包);
package containers

import "fmt"
import "container/ring"
import "github.com/kingreatwill/go_study/learning/common"

// ring99;
func ring99() {
	common.Println("container_ring.go")

	// 创建一个具有n个元素的环形链表;
	// func New(n int) *Ring;
	r := ring.New(5)
	// 返回环形链表中的元素个数;
	// func (r *Ring) Len() int;
	fmt.Println("Len:", r.Len())

	// 返回后一个元素;
	// func (r *Ring) Next() *Ring;
	for i := 1; i <= 5; i++ {
		r.Value = i
		r = r.Next()
	}
	printRing(r)
	// 返回前一个元素;
	// func (r *Ring) Prev() *Ring;

	// 移动环形链表的当前下标;
	// func (r *Ring) Move(n int) *Ring;
	r.Move(3)
	printRing(r)

	// Link连接r和s, 并返回r原本的后继元素r.Next();
	// func (r *Ring) Link(s *Ring) *Ring;
	r2 := ring.New(5)
	for i := 6; i <= 10; i++ {
		r2.Value = i
		r2 = r2.Next()
	}
	r.Link(r2)
	printRing(r)
	// 删除链表中n % r.Len()个元素, 从r.Next()开始删除;
	// func (r *Ring) Unlink(n int) *Ring;
	r.Unlink(2)
	printRing(r)

	// 对链表的每一个元素都执行f(正向顺序), 注意如果f改变了*r, Do的行为是未定义的;
	// func (r *Ring) Do(f func(interface{}));

	// 编译输出:
	// Len: 5
	// 1 2 3 4 5
	// 1 2 3 4 5
	// 1 6 7 8 9 10 2 3 4 5
	// 1 8 9 10 2 3 4 5
}

// 遍历链表(元素);
func printRing(r *ring.Ring) {
	r.Do(func(v interface{}) {
		fmt.Print(v.(int), " ")
	})
	fmt.Println()
}
