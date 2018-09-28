// GO语言中的基本数据类型和字符串的相互转换(strconv包);
package pieces

import "fmt"
import "strconv"
import "github.com/kingreatwill/go_study/learning/common"

// strconv;
func strconv01() {
	common.Println("strconv.go")

	// 常用~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~;
	// FormatBool();
	s00 := strconv.FormatBool(true)
	fmt.Println(s00)
	// FormatInt();
	s10 := strconv.FormatInt(int64(-42), 10)
	fmt.Println(s10)
	// FormatUint();
	s40 := strconv.FormatUint(uint64(42), 10)
	fmt.Println(s40)
	// FormatFloat();
	s32 := strconv.FormatFloat(3.1415926535, 'E', -1, 32)
	fmt.Println(s32)

	// Quote();
	s := strconv.Quote(`"Fran & Freddie's Diner ☺"`)
	fmt.Println(s)
	// QuoteToASCII();
	s = strconv.QuoteToASCII(`"Fran & Freddie's Diner   ☺"`)
	fmt.Println(s)
	// QuoteRune();
	s = strconv.QuoteRune('☺')
	fmt.Println(s)
	// QuoteRuneToASCII();
	s = strconv.QuoteRuneToASCII('☺')
	fmt.Println(s)

	// 常用~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~;
	bytes := []byte("strconv:")
	// 等价于append(dst, FormatBool(b)...);
	bytes = strconv.AppendBool(bytes, true)
	fmt.Println(string(bytes))
	// 等价于append(dst, FormatInt(I, base)...);
	bytes = strconv.AppendInt(bytes, -42, 10)
	fmt.Println(string(bytes))
	// 等价于append(dst, FormatUint(I, base)...);
	bytes = strconv.AppendUint(bytes, 42, 10)
	fmt.Println(string(bytes))
	// 等价于append(dst, FormatFloat(f, fmt, prec, bitSize)...);
	bytes = strconv.AppendFloat(bytes, 3.1415926535, 'E', -1, 32)
	fmt.Println(string(bytes))

	// 等价于append(dst, Quote(s)...);
	bytes = strconv.AppendQuote(bytes, `"Fran & Freddie's Diner"`)
	fmt.Println(string(bytes))
	// 等价于append(dst, QuoteToASCII(s)...);
	bytes = strconv.AppendQuoteToASCII(bytes, `"Fran & Freddie's Diner"`)
	fmt.Println(string(bytes))
	// 等价于append(dst, QuoteRune(r)...);
	bytes = strconv.AppendQuoteRune(bytes, '中')
	fmt.Println(string(bytes))
	// 等价于append(dst, QuoteRuneToASCII(r)...);
	bytes = strconv.AppendQuoteRuneToASCII(bytes, '国')
	fmt.Println(string(bytes))

	// 常用~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~;
	//func ParseBool(str string) (value bool, err error)
	//func ParseInt(s string, base int, bitSize int) (i int64, err error)
	//func ParseUint(s string, base int, bitSize int) (n uint64, err error)
	//func ParseFloat(s string, bitSize int) (f float64, err error)
	// string到int64;
	a, _ := strconv.ParseInt("90", 10, 64)
	fmt.Println(a)
	// int64到string;
	b := strconv.FormatInt(int64(90), 10)
	fmt.Println(b)

	// 常用~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~;
	// Atoi() + Itoa();
	// string到int;
	c, _ := strconv.Atoi("90")
	fmt.Println(c)
	// int到string;
	d := strconv.Itoa(90)
	fmt.Println(d)
}
