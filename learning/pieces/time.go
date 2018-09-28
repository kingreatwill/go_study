// GO语言中的时间(time包);
package pieces

import "fmt"
import "time"
import "github.com/kingreatwill/go_study/learning/common"

// Time;
func time01() {
	common.Println("time.go")
	// 参考时间;
	layout := "2006-01-02 15:04:05"
	fmt.Println(time.Now().Format(layout))
	// 字符串转换时间;
	time1, err1 := time.Parse(layout, "2018-10-10 09:00:00")
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(time1)
	// 时间相等;
	if time1.Equal(time.Now()) {
		fmt.Println("time1 is equal now.")
	}
	// 时间小于;
	if time1.Before(time.Now()) {
		fmt.Println("time1 is before now.")
	}
	// 时间大于;
	if time1.After(time.Now()) {
		fmt.Println("time1 is after now.")
	}
	// 日期(年月日);
	fmt.Println(time1.Date())
	// 日期(时分秒);
	fmt.Println(time1.Clock())
	// 日期(年);
	fmt.Println(time1.Year())
	// 日期(月);
	fmt.Println(time1.Month())
	// 日期(日);
	fmt.Println(time1.Day())
	// 日期(时);
	fmt.Println(time1.Hour())
	// 日期(分);
	fmt.Println(time1.Minute())
	// 日期(秒);
	fmt.Println(time1.Second())
	// 加减时间;
	fmt.Println(time1.AddDate(-1, 1, 1))
	// fmt.Println(time1.AddDate(-1, 1, 1).String())
}

// Duration;
func time02() {
	// 时间公式;
	// const (
	//		Nanosecond  Duration = 1
	//		Microsecond          = 1000 * Nanosecond
	//		Millisecond          = 1000 * Microsecond
	//		Second               = 1000 * Millisecond
	//		Minute               = 60 * Second
	//		Hour                 = 60 * Minute
	// )
	// 参考时间;
	layout := "2006-01-02 15:04:05"
	time1, _ := time.Parse(layout, "2017-10-10 09:00:00")
	// 至今时间差, 等价于time.Now().Sub(t);
	d1 := time.Since(time1)
	fmt.Println(d1.Hours())
	// 1秒=1000000000纳秒;
	d2 := time.Duration(time.Second)
	fmt.Println(d2.Hours())
	fmt.Println(d2.Minutes())
	fmt.Println(d2.Seconds())
	fmt.Println(d2.Nanoseconds())
}

// Timer;
func time03() {

}

// Ticker;
func time04() {

}
