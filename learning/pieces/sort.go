// GO语言中的切片和自定义数据集的排序(sort包);
package pieces

import "fmt"
import "sort"
import "github.com/kingreatwill/go_study/learning/common"

// sort;
func sort01() {
	common.Println("sort.go")

	// 升序;
	is21 := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(is21)
	fmt.Println("升序:", is21) // [1 2 3 4 5 6];

	// 是否已升序;
	fmt.Println("是否已升序:", sort.IntsAreSorted(is21)) // true;

	// 降序;
	sort.Sort(sort.Reverse(sort.IntSlice(is21)))
	fmt.Println("降序:", is21)                        // [6 5 4 3 2 1];
	fmt.Println("是否已升序:", sort.IntsAreSorted(is21)) // false;

	// 搜索并返回位置索引(已升序的切片);
	// func SearchInts(a []int, x int) int;

	// []float64;
	// func Float64s(a []float64);
	// func Float64sAreSorted(a []float64) bool;
	// func SearchFloat64s(a []float64, x float64) int;

	// []string;
	// func Strings(a []string);
	// func StringsAreSorted(a []string) bool;
	// func SearchStrings(a []string, x string) int;
}
