// GO语言中的伪随机数生成器(math/rand包);
package maths

import "fmt"
import "time"
import "math/rand"
import "github.com/kingreatwill/go_study/learning/common"

// rand01;
func rand01() {
	common.Println("math_rand.go")

	// 伪随机数(每次生成的rand值相同);
	// rand.Seed(42)
	// 同一时刻生成的rand值相同;
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(100))

	// 同一时刻rand值不重复;
	fmt.Println(time.Now().UnixNano())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		fmt.Println(r.Intn(100))
	}

	// 六位数验证码(同一时刻rand值不重复);
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	fmt.Println(vcode)
}
