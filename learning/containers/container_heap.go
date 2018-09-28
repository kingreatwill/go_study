// GO语言中的堆(container/heap包);

// 自定义一个堆需要实现5个接口:
// 1.继承自sort.Interface的接口: Len(), Less(), Swap();
// 2.堆自已的接口: Push(), Pop();

// type Interface interface {
//	 sort.Interface
//	 Push(x interface{})
//	 Pop() interface{}
// }

package containers

import "fmt"
import "container/heap"
import "github.com/kingreatwill/go_study/learning/common"

// heap99;
func heap99() {
	common.Println("container_heap.go")

	// 声明一个堆;
	h := MyHeap{4, 6, 35, 7, 5, 8, 3, 9, 0, 2, 1}

	// 一个堆在使用任何堆操作之前应先初始化;
	// func Init(h Interface);
	heap.Init(&h)
	fmt.Println("Init:", h)

	// 向堆h中插入元素x, 并保持堆的约束性;
	// func Push(h Interface, x interface{});
	heap.Push(&h, 34)
	fmt.Println("Push:", h)

	// 删除并返回堆h中的最小元素(不影响约束性);
	// func Pop(h Interface) interface{};
	fmt.Println("Pop1:", heap.Pop(&h))
	fmt.Println("Pop2:", heap.Pop(&h))
	fmt.Println("Pop:", h)

	// 删除堆中的第i个元素, 并保持堆的约束性;
	// func Remove(h Interface, i int) interface{};
	fmt.Println("Remove:", heap.Remove(&h, 1))
	fmt.Println("Remove:", h)

	// 在修改第i个元素后, 调用本函数修复堆, 比删除第i个元素后插入新元素更有效率;
	// func Fix(h Interface, i int);

	// 编译输出:
	// Init: [0 1 3 6 2 8 35 9 7 4 5]
	// Push: [0 1 3 6 2 8 35 9 7 4 5 34]
	// Pop1: 0
	// Pop2: 1
	// Pop: [2 4 3 6 5 8 35 9 7 34]
	// Remove: 4
	// Remove: [2 5 3 6 34 8 35 9 7]
}

// 堆的切片;
type MyHeap []int

// 返回集合中的元素个数;
func (h *MyHeap) Len() int {
	return len(*h)
}

// 报告索引i的元素是否比索引j的元素小;
func (h *MyHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

// 交换索引i和j的两个元素;
func (h *MyHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// 向堆h中插入元素;
func (h *MyHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// 返回堆h中的最小元素;
func (h *MyHeap) Pop() interface{} {
	o := *h
	n := len(o)
	x := o[n-1]
	*h = o[0 : n-1]
	return x
}
