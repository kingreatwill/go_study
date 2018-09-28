// GO语言中的unicode(unicode包);
package encodings

import "fmt"
import "unicode"
import "github.com/kingreatwill/go_study/learning/common"

// unicode01;
func unicode01() {
	common.Println("unicode.go")

	// 报告r是否在rangeTab指定的字符范围内;
	// func Is(rangeTab *RangeTable, r rune) bool;
	for _, r := range "Hello世界." {
		// 判断字符是否为汉字;
		if unicode.Is(unicode.Scripts["Han"], r) {
			fmt.Printf("%c", r) // 世界;
		}
	}

	// 报告r是否是给出的某个字母集的成员;
	// func In(r rune, ranges ...*RangeTable) bool;
	// 报告r是否是ranges某个成员指定的字符范围内; 本函数的功能类似In, 应优先使用In函数;
	// func IsOneOf(ranges []*RangeTable, r rune) bool;

	// 报告r是否是空白字符;
	// func IsSpace(r rune) bool;
	for _, r := range "Hello \t世　界.\n" {
		if unicode.IsSpace(r) {
			fmt.Printf("%q, ", r) // ' ', '\t', '\u3000', '\n', ;
		}
	}

	// 报告r是否是十进制数字字符;
	// func IsDigit(r rune) bool;
	for _, r := range "Hello 123１２３一二三！" {
		if unicode.IsDigit(r) {
			fmt.Printf("%c", r) // 123１２３;
		}
	}

	// 报告r是否是数字字符;
	// func IsNumber(r rune) bool;
	for _, r := range "Hello 123１２３一二三！" {
		if unicode.IsNumber(r) {
			fmt.Printf("%c", r) // 123１２３;
		}
	}

	// 报告r是否是字母;
	// func IsLetter(r rune) bool;
	for _, r := range "Hello\n\t世界！" {
		if unicode.IsLetter(r) {
			fmt.Printf("%c", r) // Hello世界;
		}
	}

	// 报告r是否是unicode图形;
	// 包括字母、标记、数字、符号、标点、空白(参见L、M、N、P、S、Zs);
	// func IsGraphic(r rune) bool;
	for _, r := range "Hello　世界！\t" {
		if unicode.IsGraphic(r) {
			fmt.Printf("%c", r) // Hello　世界！;
		}
	}

	// 报告r是否是标记字符;
	// func IsMark(r rune) bool;
	for _, r := range "Hello ៉៊់៌៍！" {
		if unicode.IsMark(r) {
			fmt.Printf("%c", r) // ៉៊់៌៍;
		}
	}

	// IsPrint一个字符是否是可打印字符;
	// func IsPrint(r rune) bool;
	for _, r := range "Hello　世界！\t" {
		if unicode.IsPrint(r) {
			fmt.Printf("%c", r) // Hello世界！;
		}
	}

	// 报告r是否是控制字符;
	// func IsControl(r rune) bool;
	for _, r := range "Hello\n\t世界！" {
		if unicode.IsControl(r) {
			fmt.Printf("%#q, ", r) // '\n', '\t', ;
		}
	}

	// 报告r是否是unicode标点字符;\
	// func IsPunct(r rune) bool;
	for _, r := range "Hello 世界！" {
		if unicode.IsPunct(r) {
			fmt.Printf("%c", r) // ！;
		}
	}

	// 报告r是否是unicode符号字符;
	// func IsSymbol(r rune) bool;
	for _, r := range "Hello (<世=界>)" {
		if unicode.IsSymbol(r) {
			fmt.Printf("%c", r)
		}
	}

	fmt.Println()

	// 返回字符是否是大写字母;
	// func IsUpper(r rune) bool;
	for _, r := range "Hello ＡＢＣ！" {
		if unicode.IsUpper(r) {
			fmt.Printf("%c", r) // HＡＢＣ;
		}
	}

	// 返回字符是否是小写字母;
	// func IsLower(r rune) bool;
	for _, r := range "Hello ａｂｃ！" {
		if unicode.IsLower(r) {
			fmt.Printf("%c", r) // elloａｂｃ;
		}
	}

	// 返回字符是否是标题字母;
	// func IsTitle(r rune) bool;
	for _, r := range "Hello ᾏᾟᾯ！" {
		if unicode.IsTitle(r) {
			fmt.Printf("%c", r) // ᾏᾟᾯ;
		}
	}

	// 返回_case指定的对应类型的字母：UpperCase、LowerCase、TitleCase;
	// func To(_case int, r rune) rune;
	fmt.Println()

	// 返回对应的小写字母;
	// func ToLower(r rune) rune;
	for _, r := range "Hello 世界！" {
		// fmt.Printf("%c", unicode.ToLower(r))            // hello 世界！;
		fmt.Printf("%c", unicode.To(unicode.LowerCase, r)) // hello 世界！;
	}

	// 返回对应的大写字母;
	// func ToUpper(r rune) rune;
	for _, r := range "Hello 世界！" {
		// fmt.Printf("%c", unicode.ToUpper(r))            // HELLO 世界！;
		fmt.Printf("%c", unicode.To(unicode.UpperCase, r)) // HELLO 世界！;
	}

	// 返回对应的标题字母;
	// func ToTitle(r rune) rune;
	for _, r := range "Hello 世界！" {
		// fmt.Printf("%c", unicode.ToTitle(r))            // HELLO 世界！;
		fmt.Printf("%c", unicode.To(unicode.TitleCase, r)) // HELLO 世界！;
	}

	// 迭代在unicode标准字符映射中互相对应的unicode码值;
	// func SimpleFold(r rune) rune;
	fmt.Println()
}
