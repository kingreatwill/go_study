// GO语言中的文本模板数据驱动输出(text/template包);
// 用作模板的输入文本必须是utf-8编码的文本;
// Action————数据运算和控制单位, 由"{{"和"}}"界定, 在Action之外的所有文本都不做修改的拷贝到输出中;
// Action内部不能有换行, 但注释可以有换行;
package nets

import "fmt"
import "os"
import "strings"
import "text/template"
import "github.com/kingreatwill/go_study/learning/common"

// 文本模板数据驱动输出(简单例子);
func text_template01() {
	common.Println("text_template.go")

	type Inventory struct {
		Material string
		Count    uint
	}
	sweaters := Inventory{"wool", 17}

	tmpl, err := template.New("test").Parse("{{.Count}} of {{.Material}}.\r\n")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}

// ######################################Actions(动作)的定义######################################
// {{/* a comment */}}
// 注释，执行时会忽略。可以多行。注释不能嵌套，并且必须紧贴分界符始止，就像这里表示的一样。
// {{pipeline}}
// pipeline的值的默认文本表示会被拷贝到输出里。
// {{if pipeline}} T1 {{end}}
// 如果pipeline的值为empty，不产生输出，否则输出T1执行结果。不改变dot的值。Empty值包括false、0、任意nil指针或者nil接口，任意长度为0的数组、切片、字典。
// {{if pipeline}} T1 {{else}} T0 {{end}}
// 如果pipeline的值为empty，输出T0执行结果，否则输出T1执行结果。不改变dot的值。
// {{if pipeline}} T1 {{else if pipeline}} T0 {{end}}
// 用于简化if-else链条，else action可以直接包含另一个if；等价于：{{if pipeline}} T1 {{else}}{{if pipeline}} T0 {{end}}{{end}}
// {{range pipeline}} T1 {{end}}
// pipeline的值必须是数组、切片、字典或者通道。如果pipeline的值其长度为0，不会有任何输出；否则dot依次设为数组、切片、字典或者通道的每一个成员元素并执行T1；如果pipeline的值为字典，且键可排序的基本类型，元素也会按键的顺序排序。
// {{range pipeline}} T1 {{else}} T0 {{end}}
// pipeline的值必须是数组、切片、字典或者通道。如果pipeline的值其长度为0，不改变dot的值并执行T0；否则会修改dot并执行T1。
// {{template "name"}}
// 执行名为name的模板，提供给模板的参数为nil，如模板不存在输出为""
// {{template "name" pipeline}}
// 执行名为name的模板，提供给模板的参数为pipeline的值。
// {{with pipeline}} T1 {{end}}
// 如果pipeline为empty不产生输出，否则将dot设为pipeline的值并执行T1。不修改外面的dot。
// {{with pipeline}} T1 {{else}} T0 {{end}}
// 如果pipeline为empty，不改变dot并执行T0，否则dot设为pipeline的值并执行T1。

// ######################################Arguments(参数)的定义######################################
// - go语法的布尔值、字符串、字符、整数、浮点数、虚数、复数，视为无类型字面常数，字符串不能跨行
// - 关键字nil，代表一个go的无类型的nil值
// - 字符'.'（句点，用时不加单引号），代表dot的值
// - 变量名，以美元符号起始加上（可为空的）字母和数字构成的字符串，如：$piOver2和$；
//   执行结果为变量的值，变量参见下面的介绍
// - 结构体数据的字段名，以句点起始，如：.Field；
//   执行结果为字段的值，支持链式调用：.Field1.Field2；
//   字段也可以在变量上使用（包括链式调用）：$x.Field1.Field2；
// - 字典类型数据的键名；以句点起始，如：.Key；
//   执行结果是该键在字典中对应的成员元素的值；
//   键也可以和字段配合做链式调用，深度不限：.Field1.Key1.Field2.Key2；
//   虽然键也必须是字母和数字构成的标识字符串，但不需要以大写字母起始；
//   键也可以用于变量（包括链式调用）：$x.key1.key2；
// - 数据的无参数方法名，以句点为起始，如：.Method；
//   执行结果为dot调用该方法的返回值，dot.Method()；
//   该方法必须有1到2个返回值，如果有2个则后一个必须是error接口类型；
//   如果有2个返回值的方法返回的error非nil，模板执行会中断并返回给调用模板执行者该错误；
//   方法可和字段、键配合做链式调用，深度不限：.Field1.Key1.Method1.Field2.Key2.Method2；
//   方法也可以在变量上使用（包括链式调用）：$x.Method1.Field；
// - 无参数的函数名，如：fun；
//   执行结果是调用该函数的返回值fun()；对返回值的要求和方法一样；函数和函数名细节参见后面。
// - 上面某一条的实例加上括弧（用于分组）
//   执行结果可以访问其字段或者键对应的值：
//   print (.F1 arg1) (.F2 arg2)
//   (.StructValuedMethod "arg").Field

// ######################################Pipelines(管道)的定义######################################
// pipeline通常是将一个command序列分割开，再使用管道符'|'连接起来(但不使用管道符的command序列也可以视为一个管道);
// 在一个链式的pipeline里, 每个command的结果都作为下一个command的最后一个参数, pipeline最后一个command的输出作为整个管道执行的结果;
// Pipeline是一个(可能是链状的)command序列, Command可以是一个简单值(Argument)/对函数/方法的(可以有多个参数的)调用：
// Argument
// 执行结果是argument的执行结果
// .Method [Argument...]
// 方法可以独立调用或者位于链式调用的末端，不同于链式调用中间的方法，可以使用参数;
// 执行结果是使用给出的参数调用该方法的返回值：dot.Method(Argument1, etc.);
// functionName [Argument...]
// 执行结果是使用给定的参数调用函数名指定的函数的返回值：function(Argument1, etc.);

// ######################################单行模板示例######################################
// 单行模板(展示pipeline和变量), 所有都生成加引号的单词"output";
// {{"\"output\""}}
// 字符串常量
// {{`"output"`}}
// 原始字符串常量
// {{printf "%q" "output"}}
// 函数调用
// {{"output" | printf "%q"}}
// 函数调用，最后一个参数来自前一个command的返回值
// {{printf "%q" (print "out" "put")}}
// 加括号的参数
// {{"put" | printf "%s%s" "out" | printf "%q"}}
// 玩出花的管道的链式调用
// {{"output" | printf "%s" | printf "%q"}}
// 管道的链式调用
// {{with "output"}}{{printf "%q" .}}{{end}}
// 使用dot的with action
// {{with $x := "output" | printf "%q"}}{{$x}}{{end}}
// 创建并使用变量的with action
// {{with $x := "output"}}{{printf "%q" $x}}{{end}}
// 将变量使用在另一个action的with action
// {{with $x := "output"}}{{$x | printf "%q"}}{{end}}
// 以管道形式将变量使用在另一个action的with action

// ######################################Functions(预定义的全局函数)的定义######################################
// and
//     函数返回它的第一个empty参数或者最后一个参数；
//     就是说"and x y"等价于"if x then y else x"；所有参数都会执行；
// or
//     返回第一个非empty参数或者最后一个参数；
//     亦即"or x y"等价于"if x then x else y"；所有参数都会执行；
// not
//     返回它的单个参数的布尔值的否定
// len
//     返回它的参数的整数类型长度
// index
//     执行结果为第一个参数以剩下的参数为索引/键指向的值；
//     如"index x 1 2 3"返回x[1][2][3]的值；每个被索引的主体必须是数组、切片或者字典。
// print
//     即fmt.Sprint
// printf
//     即fmt.Sprintf
// println
//     即fmt.Sprintln
// html
//     返回其参数文本表示的HTML逸码等价表示。
// urlquery
//     返回其参数文本表示的可嵌入URL查询的逸码等价表示。
// js
//     返回其参数文本表示的JavaScript逸码等价表示。
// call
//     执行结果是调用第一个参数的返回值，该参数必须是函数类型，其余参数作为调用该函数的参数；
//     如"call .X.Y 1 2"等价于go语言里的dot.X.Y(1, 2)；
//     其中Y是函数类型的字段或者字典的值，或者其他类似情况；
//     call的第一个参数的执行结果必须是函数类型的值（和预定义函数如print明显不同）；
//     该函数类型值必须有1到2个返回值，如果有2个则后一个必须是error接口类型；
//     如果有2个返回值的方法返回的error非nil，模板执行会中断并返回给调用模板执行者该错误；
// 定义为函数的二元比较运算的集合
//     eq      如果arg1 == arg2则返回真
//     ne      如果arg1 != arg2则返回真
//     lt      如果arg1 < arg2则返回真
//     le      如果arg1 <= arg2则返回真
//     gt      如果arg1 > arg2则返回真
//     ge      如果arg1 >= arg2则返回真

// ######################################Nested template definitions(嵌套模板)的定义######################################
// `{{define "T1"}}ONE{{end}}
// {{define "T2"}}TWO{{end}}
// {{define "T3"}}{{template "T1"}} {{template "T2"}}{{end}}
// {{template "T3"}}`

// 其他函数;
func text_template02() {
	// 向w中写入b的HTML转义等价表示;
	// func HTMLEscape(w io.Writer, b []byte);
	// 返回s的HTML转义等价表示字符串;
	// func HTMLEscapeString(s string) string;
	// 返回其所有参数文本表示的HTML转义等价表示字符串;
	// func HTMLEscaper(args ...interface{}) string;

	// 向w中写入b的JavaScript转义等价表示;
	// func JSEscape(w io.Writer, b []byte);
	// 返回s的JavaScript转义等价表示字符串;
	// func JSEscapeString(s string) string;
	// 返回其所有参数文本表示的JavaScript转义等价表示字符串;
	// func JSEscaper(args ...interface{}) string;

	// 返回其所有参数文本表示的可以嵌入URL查询的转义等价表示字符串;
	// func URLQueryEscaper(args ...interface{}) string;
}

// 文本模板数据驱动输出;
func text_template03() {
	// 用于包装返回(*Template, error)的函数/方法调用, 它会在err非nil时panic, 一般用于变量初始化;
	// func Must(t *Template, err error) *Template;
	// 创建一个名为name的模板;
	// func New(name string) *Template;
	var tmpl = template.Must(template.New("name").Parse("text...\r\n"))

	// 创建一个模板并解析filenames指定的文件里的模板定义;
	// func ParseFiles(filenames ...string) (*Template, error);
	// 创建一个模板并解析匹配pattern的文件里的模板定义;
	// func ParseGlob(pattern string) (*Template, error);

	// 返回模板t的名字;
	// func (t *Template) Name() string;
	// 用于设置action的分界字符串, 应用于之后的Parse、ParseFiles、ParseGlob方法(默认{{、}});
	// func (t *Template) Delims(left, right string) *Template;
	// 向模板t的函数字典里加入参数funcMap内的键值对;
	// func (t *Template) Funcs(funcMap FuncMap) *Template;

	// 返回模板的一个副本, 包括所有相关联的模板;
	// func (t *Template) Clone() (*Template, error);
	// 返回与t关联的名为name的模板;
	// func (t *Template) Lookup(name string) *Template;
	// 返回与t相关联的模板的切片, 包括t自己;
	// func (t *Template) Templates() []*Template;
	// 创建一个和t关联的名字为name的模板并返回它;
	// func (t *Template) New(name string) *Template;
	// 使用name和tree创建一个模板并使它和t相关联;
	// func (t *Template) AddParseTree(name string, tree *parse.Tree) (*Template, error);

	// 将字符串text解析为模板;
	// func (t *Template) Parse(text string) (*Template, error);
	// 解析filenames指定的文件里的模板定义并将解析结果与t关联;
	// func (t *Template) ParseFiles(filenames ...string) (*Template, error);
	// 解析匹配pattern的文件里的模板定义并将解析结果与t关联;
	// func (t *Template) ParseGlob(pattern string) (*Template, error);

	// 将解析好的模板应用到data上, 并将输出写入wr;
	// func (t *Template) Execute(wr io.Writer, data interface{}) (err error);
	// 类似Execute方法, 但是使用名为name的t关联的模板产生输出;
	// func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error;
	tmpl.Execute(os.Stdout, nil)
}

// 文本模板数据驱动输出(常用);
func text_template04() {
	const letter = `
        Dear {{.Name}},
        {{if .Attended}}
        It was a pleasure to see you at the wedding.{{else}}
        It is a shame you couldn't make it to the wedding.{{end}}
        {{with .Gift}}Thank you for the lovely {{.}}.
        {{end}}
        Best wishes,
        Josie
    `
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rodney", "", false},
	}
	tmpl := template.Must(template.New("letter").Parse(letter))
	for _, recipient := range recipients {
		err := tmpl.Execute(os.Stdout, recipient)
		if err != nil {
			fmt.Println("executing template:", err)
		}
	}
}

// 文本模板数据驱动输出(FuncMap);
func text_template05() {
	funcMap := template.FuncMap{
		"title": strings.Title,
	}
	const templateText = `
        Input: {{printf "%q" .}}
        Output 0: {{title .}}
        Output 1: {{title . | printf "%q"}}
        Output 2: {{printf "%q" . | title}}
    `
	tmpl, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		fmt.Println("parsing:", err)
	}
	err = tmpl.Execute(os.Stdout, "the go programming language")
	if err != nil {
		fmt.Println("parsing:", err)
	}
	// Input: "the go programming language"
	// Output 0: The Go Programming Language
	// Output 1: "The Go Programming Language"
	// Output 2: "The Go Programming Language"
}
