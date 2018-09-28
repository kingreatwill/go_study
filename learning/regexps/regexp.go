// GO语言中的正则表达式搜索(regexp包);
// 正则表达式采用RE2语法(除了\c、\C), 和Perl、Python等语言的正则基本一致;
package regexps

import "fmt"
import "regexp"
import "github.com/kingreatwill/go_study/learning/common"

// regexp01;
func regexp01() {
	common.Println("regexp.go")

	// (快速常用)func MatchString(pattern string, s string) (matched bool, err error);
	var rg = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	fmt.Println("MatchString:", rg.MatchString("adam[23]")) // true;
	fmt.Println(rg.MatchString("eve[7]"))                   // true;
	fmt.Println(rg.MatchString("Job[48]"))                  // false;
	fmt.Println(rg.MatchString("snakey"))                   // false;

	// func MatchString(pattern string, s string) (matched bool, err error);
	matched, err := regexp.MatchString("foo.*", "seafood") // true <nil>;
	fmt.Println("MatchString:", matched, err)
	matched, err = regexp.MatchString("bar.*", "seafood") // false <nil>;
	fmt.Println(matched, err)
	matched, err = regexp.MatchString("a(b", "seafood") // false error parsing regexp: missing closing ): `a(b`;
	fmt.Println(matched, err)

	// func (re *Regexp) FindAllString(s string, n int) []string;
	re := regexp.MustCompile("a.")
	fmt.Println("FindAllString:", re.FindAllString("paranormal", -1)) // [ar an al];
	fmt.Println(re.FindAllString("paranormal", 2))                    // [ar an];
	fmt.Println(re.FindAllString("graal", -1))                        // [aa];
	fmt.Println(re.FindAllString("none", -1))                         // [];

	// func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string;
	re = regexp.MustCompile("a(x*)b")
	fmt.Printf("FindAllStringSubmatch: %q\n", re.FindAllStringSubmatch("-ab-", -1)) // [["ab" ""]];
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-", -1))                      // [["axxb" "xx"]];
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-axb-", -1))                    // [["ab" ""] ["axb" "x"]];
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-ab-", -1))                   // [["axxb" "xx"] ["ab" ""]];

	// func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int;
	re = regexp.MustCompile("a(x*)b")
	fmt.Println("FindAllStringSubmatchIndex:", re.FindAllStringSubmatchIndex("-ab-", -1)) // [[1 3 2 2]];
	fmt.Println(re.FindAllStringSubmatchIndex("-axxb-", -1))                              // [[1 5 2 4]];
	fmt.Println(re.FindAllStringSubmatchIndex("-ab-axb-", -1))                            // [[1 3 2 2] [4 7 5 6]];
	fmt.Println(re.FindAllStringSubmatchIndex("-axxb-ab-", -1))                           // [[1 5 2 4] [6 8 7 7]];
	fmt.Println(re.FindAllStringSubmatchIndex("-foo-", -1))                               // [];

	// func (re *Regexp) FindString(s string) string;
	re = regexp.MustCompile("fo.?")
	fmt.Printf("FindString: %q\n", re.FindString("seafood")) // "foo";
	fmt.Printf("%q\n", re.FindString("meat"))                // "";

	// func (re *Regexp) FindStringIndex(s string) (loc []int);
	re = regexp.MustCompile("ab?")
	fmt.Println("FindStringIndex:", re.FindStringIndex("tablett")) // [1 3];
	fmt.Println(re.FindStringIndex("foo") == nil)                  // true;

	// func (re *Regexp) FindStringSubmatch(s string) []string;
	re = regexp.MustCompile("a(x*)b(y|z)c")
	fmt.Printf("FindStringSubmatch: %q\n", re.FindStringSubmatch("-axxxbyc-")) // ["axxxbyc" "xxx" "y"];
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))                        // ["abzc" "" "z"];

	// func (re *Regexp) ReplaceAllLiteralString(src, repl string) string;
	re = regexp.MustCompile("a(x*)b")
	fmt.Println("ReplaceAllLiteralString:", re.ReplaceAllLiteralString("-ab-axxb-", "T")) // -T-T-;
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "$1"))                            // -$1-$1-;
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "${1}"))                          // -${1}-${1}-;

	// func (re *Regexp) ReplaceAllString(src, repl string) string;
	re = regexp.MustCompile("a(x*)b")
	fmt.Println("ReplaceAllString:", re.ReplaceAllString("-ab-axxb-", "T")) // -T-T-;
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1"))                     // --xx-;
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1W"))                    // ---;
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "${1}W"))                  // -W-xxW-;

	// func (re *Regexp) SubexpNames() []string;
	re = regexp.MustCompile("(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)")
	fmt.Println("SubexpNames:", re.MatchString("Alan Turing")) // true;
	fmt.Printf("%q\n", re.SubexpNames())                       // ["" "first" "last"];
	reversed := fmt.Sprintf("${%s} ${%s}", re.SubexpNames()[2], re.SubexpNames()[1])
	fmt.Println(reversed)                                     // ${last} ${first};
	fmt.Println(re.ReplaceAllString("Alan Turing", reversed)) // Turing Alan;
}
